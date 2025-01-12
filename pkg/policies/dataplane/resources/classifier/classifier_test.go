package classifier_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/open-policy-agent/opa/ast"

	labelmatcherv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/common/labelmatcher/v1"
	selectorv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/common/selector/v1"
	classificationv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/policy/language/v1"
	"github.com/fluxninja/aperture/pkg/log"

	. "github.com/fluxninja/aperture/pkg/policies/dataplane/resources/classifier"
	"github.com/fluxninja/aperture/pkg/policies/dataplane/resources/classifier/compiler"
	"github.com/fluxninja/aperture/pkg/selectors"
	"github.com/fluxninja/aperture/pkg/services"
)

type object = map[string]interface{}

var _ = Describe("Classifier", func() {
	var classifier *Classifier

	BeforeEach(func() {
		log.SetGlobalLevel(log.WarnLevel)

		classifier = New()
	})

	It("returns empty slice, when no rules configured", func() {
		Expect(classifier.ActiveRules()).To(BeEmpty())
	})

	Context("configured with some classification rules", func() {
		// Classifier with a simple extractor-based rule
		rs1 := &classificationv1.Classifier{
			Selector: &selectorv1.Selector{
				Service: "my-service.default.svc.cluster.local",
				ControlPoint: &selectorv1.ControlPoint{
					Controlpoint: &selectorv1.ControlPoint_Traffic{
						Traffic: "ingress",
					},
				},
			},
			Rules: map[string]*classificationv1.Rule{
				"foo": {
					Source:    headerExtractor("foo"),
					Propagate: true,
				},
			},
		}

		// Classifier with Raw-rego rule, additionally gated for just "version one"
		rs2 := &classificationv1.Classifier{
			Selector: &selectorv1.Selector{
				Service: "my-service.default.svc.cluster.local",
				LabelMatcher: &labelmatcherv1.LabelMatcher{
					MatchLabels: map[string]string{"version": "one"},
				},
				ControlPoint: &selectorv1.ControlPoint{
					Controlpoint: &selectorv1.ControlPoint_Traffic{
						Traffic: "ingress",
					},
				},
			},
			Rules: map[string]*classificationv1.Rule{
				"bar-twice": {
					Source: &classificationv1.Rule_Rego_{
						Rego: &classificationv1.Rule_Rego{
							Source: `
								package my.pkg
								answer := input.attributes.request.http.headers.bar * 2
							`,
							Query: "data.my.pkg.answer",
						},
					},
					Propagate: true,
				},
			},
		}

		// Classifier with a no service populated
		rs3 := &classificationv1.Classifier{
			Selector: &selectorv1.Selector{
				ControlPoint: &selectorv1.ControlPoint{
					Controlpoint: &selectorv1.ControlPoint_Traffic{
						Traffic: "ingress",
					},
				},
			},
			Rules: map[string]*classificationv1.Rule{
				"fuu": {
					Source:    headerExtractor("fuu"),
					Propagate: true,
				},
			},
		}

		var ars1, ars2, ars3 ActiveRuleset
		BeforeEach(func() {
			var err error
			ars1, err = classifier.AddRules(context.TODO(), "one", rs1)
			Expect(err).NotTo(HaveOccurred())
			ars2, err = classifier.AddRules(context.TODO(), "two", rs2)
			Expect(err).NotTo(HaveOccurred())
			ars3, err = classifier.AddRules(context.TODO(), "three", rs3)
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns active rules", func() {
			Expect(classifier.ActiveRules()).To(ConsistOf(
				compiler.ReportedRule{
					RulesetName: "one",
					LabelName:   "foo",
					Rule:        rs1.Rules["foo"],
					Selector:    rs1.Selector,
				},
				compiler.ReportedRule{
					RulesetName: "two",
					LabelName:   "bar-twice",
					Rule:        rs2.Rules["bar-twice"],
					Selector:    rs2.Selector,
				},
				compiler.ReportedRule{
					RulesetName: "three",
					LabelName:   "fuu",
					Rule:        rs3.Rules["fuu"],
					Selector:    rs3.Selector,
				},
			))
		})

		It("classifies input by returning flow labels", func() {
			labels, err := classifier.Classify(
				context.TODO(),
				[]services.ServiceID{{
					Service: "my-service.default.svc.cluster.local",
				}},
				map[string]string{"version": "one", "other": "tag"},
				selectors.Ingress,
				attributesWithHeaders(object{
					"foo": "hello",
					"bar": 21,
				}),
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(labels).To(Equal(FlowLabels{
				"foo":       fl("hello"),
				"bar-twice": fl("42"),
			}))
		})

		It("doesn't classify if direction doesn't match", func() {
			labels, err := classifier.Classify(
				context.TODO(),
				[]services.ServiceID{{
					Service: "my-service.default.svc.cluster.local",
				}},
				map[string]string{"version": "one"},
				selectors.Egress,
				attributesWithHeaders(object{
					"foo": "hello",
					"bar": 21,
				}),
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(labels).To(BeEmpty())
		})

		It("skips rules with non-matching labels", func() {
			labels, err := classifier.Classify(
				context.TODO(),
				[]services.ServiceID{{
					Service: "my-service.default.svc.cluster.local",
				}},
				map[string]string{"version": "two"},
				selectors.Ingress,
				attributesWithHeaders(object{
					"foo": "hello",
					"bar": 21,
				}),
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(labels).To(Equal(FlowLabels{
				"foo": fl("hello"),
			}))
		})

		Context("when ruleset is dropped", func() {
			BeforeEach(func() { ars1.Drop() })

			It("removes removes subset of rules", func() {
				labels, err := classifier.Classify(
					context.TODO(),
					[]services.ServiceID{{
						Service: "my-service.default.svc.cluster.local",
					}},
					map[string]string{"version": "one"},
					selectors.Ingress,
					attributesWithHeaders(object{
						"foo": "hello",
						"bar": 21,
					}),
				)
				Expect(err).NotTo(HaveOccurred())
				Expect(labels).To(Equal(FlowLabels{
					"bar-twice": fl("42"),
				}))
			})
		})

		Context("when all rulesets dropped", func() {
			BeforeEach(func() {
				ars1.Drop()
				ars2.Drop()
				ars3.Drop()
			})

			It("removes all the rules", func() {
				Expect(classifier.ActiveRules()).To(BeEmpty())
			})
		})
	})

	// helper for setting rules with a "default" selector
	setRulesForMyService := func(labelRules map[string]*classificationv1.Rule) error {
		_, err := classifier.AddRules(context.TODO(), "test", &classificationv1.Classifier{
			Selector: &selectorv1.Selector{
				Service: "my-service.default.svc.cluster.local",
				ControlPoint: &selectorv1.ControlPoint{
					Controlpoint: &selectorv1.ControlPoint_Traffic{
						Traffic: "ingress",
					},
				},
			},
			Rules: labelRules,
		})
		return err
	}

	Context("configured classification rules with some label flags", func() {
		rules := map[string]*classificationv1.Rule{
			"foo": {
				Source:    headerExtractor("foo"),
				Propagate: false,
			},
			"bar": {
				Source: &classificationv1.Rule_Rego_{
					Rego: &classificationv1.Rule_Rego{
						Source: `
							package my.pkg
							answer := input.attributes.request.http.headers.bar
							`,
						Query: "data.my.pkg.answer",
					},
				},
				Propagate: true,
				Hidden:    true,
			},
		}

		BeforeEach(func() {
			Expect(setRulesForMyService(rules)).To(Succeed())
		})

		It("marks the returned flow labels with those flags", func() {
			labels, err := classifier.Classify(
				context.TODO(),
				[]services.ServiceID{{
					Service: "my-service.default.svc.cluster.local",
				}},
				nil,
				selectors.Ingress,
				attributesWithHeaders(object{
					"foo": "hello",
					"bar": 21,
				}),
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(labels).To(Equal(FlowLabels{
				"foo": FlowLabelValue{"hello", compiler.LabelFlags{Propagate: false}},
				"bar": FlowLabelValue{"21", compiler.LabelFlags{Hidden: true, Propagate: true}},
			}))
		})
	})

	Context("configured with same label for different rules in yaml", func() {
		// Note: we don't support multiple rules for the same label in a single
		// rulesets. But we might add support in the future, eg.:
		// "foo/1": ...
		// "foo/2": ...
		rules1 := map[string]*classificationv1.Rule{
			"foo": {
				Source:    headerExtractor("foo"),
				Propagate: true,
			},
		}
		rules2 := map[string]*classificationv1.Rule{
			"foo": {
				Source:    headerExtractor("xyz"),
				Propagate: true,
			},
		}

		BeforeEach(func() {
			Expect(setRulesForMyService(rules1)).To(Succeed())
			Expect(setRulesForMyService(rules2)).To(Succeed())
		})

		It("classifies and returns flow labels (overwrite order not specified)", func() {
			// Perhaps we can specify order by sorting rulesets? (eg. giving
			// them names from filenames)
			labels, err := classifier.Classify(
				context.TODO(),
				[]services.ServiceID{{
					Service: "my-service.default.svc.cluster.local",
				}},
				nil,
				selectors.Ingress,
				attributesWithHeaders(object{
					"foo": "hello",
					"xyz": "cos",
					"bar": 21,
				}),
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(labels).To(SatisfyAny(
				Equal(FlowLabels{"foo": fl("cos")}),
				Equal(FlowLabels{"foo": fl("hello")}),
			))
		})
	})

	Context("configured with same label for different rules in rego", func() {
		rules1 := map[string]*classificationv1.Rule{
			"bar": {
				Source: &classificationv1.Rule_Rego_{
					Rego: &classificationv1.Rule_Rego{
						Source: `
							package my.pkg
							answer := input.attributes.request.http.headers.bar * 3
						`,
						Query: "data.my.pkg.answer",
					},
				},
				Propagate: true,
			},
		}
		rules2 := map[string]*classificationv1.Rule{
			"bar": {
				Source: &classificationv1.Rule_Rego_{
					Rego: &classificationv1.Rule_Rego{
						Source: `
							package my.pkg
							answer2 := input.attributes.request.http.headers.bar * 2
						`,
						Query: "data.my.pkg.answer2",
					},
				},
				Propagate: true,
			},
		}

		BeforeEach(func() {
			Expect(setRulesForMyService(rules1)).To(Succeed())
			Expect(setRulesForMyService(rules2)).To(Succeed())
		})

		It("classifies and returns flow labels (overwrite order not specified)", func() {
			// Perhaps we can specify order by sorting rulesets? (eg. giving
			// them names from filenames)
			labels, err := classifier.Classify(
				context.TODO(),
				[]services.ServiceID{{
					Service: "my-service.default.svc.cluster.local",
				}},
				nil,
				selectors.Ingress,
				attributesWithHeaders(object{
					"foo": "hello",
					"bar": 21,
				}),
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(labels).To(SatisfyAny(
				Equal(FlowLabels{"bar": fl("63")}),
				Equal(FlowLabels{"bar": fl("42")}),
			))
		})
	})

	Context("incorrect rego passed", func() {
		rules := map[string]*classificationv1.Rule{
			"bar-twice": {
				Source: &classificationv1.Rule_Rego_{
					Rego: &classificationv1.Rule_Rego{
						Source: `
							Package my.pkg
							bar := input.attributes.request.http.headers.bar * 2
							bar := input.attributes.request.http.headers.foo
						`,
						Query: "data.my.pkg.bar",
					},
				},
				Propagate: true,
			},
		}

		It("fails to compile rego", func() {
			err := setRulesForMyService(rules)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(compiler.BadRego))
		})
	})

	Context("configured with ambiguous rules in rego", func() {
		rules := map[string]*classificationv1.Rule{
			"bar": {
				Source: &classificationv1.Rule_Rego_{
					Rego: &classificationv1.Rule_Rego{
						Source: `
							package my.pkg
							answer = input.attributes.request.http.headers.bar * 3
							answer = input.attributes.request.http.headers.foo
						`,
						Query: "data.my.pkg.answer",
					},
				},
				Propagate: true,
			},
		}

		BeforeEach(func() {
			Expect(setRulesForMyService(rules)).To(Succeed())
		})

		It("classifies and returns empty flow labels - could not decide which rego to use", func() {
			labels, err := classifier.Classify(
				context.TODO(),
				[]services.ServiceID{{
					Service: "my-service.default.svc.cluster.local",
				}},
				nil,
				selectors.Ingress,
				attributesWithHeaders(object{
					"foo": "hello",
					"bar": 21,
				}),
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(labels).To(Equal(FlowLabels{}))
		})
	})

	Context("configured with invalid label name", func() {
		// Classifier with a simple extractor-based rule
		rs := &classificationv1.Classifier{
			Selector: &selectorv1.Selector{
				Service: "my-service.default.svc.cluster.local",
				ControlPoint: &selectorv1.ControlPoint{
					Controlpoint: &selectorv1.ControlPoint_Traffic{
						Traffic: "ingress",
					},
				},
			},
			Rules: map[string]*classificationv1.Rule{
				"user-agent": {
					Source:    headerExtractor("foo"),
					Propagate: true,
				},
			},
		}

		It("should reject the ruleset", func() {
			_, err := classifier.AddRules(context.TODO(), "one", rs)
			Expect(err).To(HaveOccurred())
		})
	})
})

func fl(s string) FlowLabelValue {
	return FlowLabelValue{
		Value: s,
		Flags: compiler.LabelFlags{Propagate: true},
	}
}

func attributesWithHeaders(headers object) ast.Value {
	return ast.MustInterfaceToValue(
		object{
			"attributes": object{
				"request": object{
					"http": object{
						"headers": headers,
					},
				},
			},
		},
	)
}

func headerExtractor(headerName string) *classificationv1.Rule_Extractor {
	return &classificationv1.Rule_Extractor{
		Extractor: &classificationv1.Extractor{
			Variant: &classificationv1.Extractor_From{
				From: "request.http.headers." + headerName,
			},
		},
	}
}
