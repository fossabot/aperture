package iface

import (
	"github.com/prometheus/client_golang/prometheus"

	flowcontrolv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/flowcontrol/v1"
)

// ResponseMetricsAPI is an interface for getting response metrics.
type ResponseMetricsAPI interface {
	GetFluxMeterHistogram(fluxmeterID, statusCode string, decisionType flowcontrolv1.DecisionType) (prometheus.Observer, error)
	GetTokenLatencyHistogram(labels map[string]string) (prometheus.Observer, error)

	DeleteFluxmeterHistogram(fluxmeterID string) bool
	DeleteTokenLatencyHistogram(policyName, policyHash, componentIndex string) bool
}
