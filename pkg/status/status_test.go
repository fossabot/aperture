package status_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	statusv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/common/status/v1"
	"github.com/fluxninja/aperture/pkg/status"
)

var (
	rootRegistry status.Registry
)

var _ = Describe("Status Registry", func() {
	BeforeEach(func() {
		rootRegistry = status.NewRegistry()
	})

	Context("there is single level root registry", func() {
		It("returns the root registry itself", func() {
			Expect(rootRegistry.Parent()).To(BeNil())
			Expect(rootRegistry.Root()).To(Equal(rootRegistry))
			Expect(rootRegistry.ChildIfExists(rootRegistry.Key())).To(BeNil())
		})
		It("returns empty status information", func() {
			Expect(rootRegistry.GetStatus()).To(Equal(&statusv1.Status{}))
			Expect(rootRegistry.GetGroupStatus()).To(Equal(&statusv1.GroupStatus{
				Status: &statusv1.Status{},
				Groups: make(map[string]*statusv1.GroupStatus),
			}))
		})
		It("returns updated status information", func() {
			test_status := status.NewStatus(nil, errors.New("test status"))
			rootRegistry.SetStatus(test_status)
			Expect(rootRegistry.GetStatus()).To(Equal(test_status))
			Expect(rootRegistry.HasError()).To(BeTrue())

			test_groupstatus := &statusv1.GroupStatus{
				Status: rootRegistry.GetStatus(),
				Groups: make(map[string]*statusv1.GroupStatus),
			}
			rootRegistry.SetGroupStatus(test_groupstatus)
			Expect(rootRegistry.GetGroupStatus()).To(Equal(test_groupstatus))
		})
		It("creates a new child registry and then detaches it", func() {
			child_registry := rootRegistry.Child("child")
			Expect(rootRegistry.ChildIfExists("child")).To(Equal(child_registry))
			rootRegistry.Child("child").Detach()
			Expect(rootRegistry.ChildIfExists("child")).To(BeNil())
		})
		It("returns the key of the root registry", func() {
			Expect(rootRegistry.Key()).To(Equal(""))
		})
	})

	var (
		child1      status.Registry
		child2      status.Registry
		grandChild1 status.Registry
		grandChild2 status.Registry
	)
	Context("there multiple registries in hierarchy", func() {
		BeforeEach(func() {
			child1 = rootRegistry.Child("child1")
			child2 = rootRegistry.Child("child2")
			grandChild1 = child1.Child("grandChild1")
			grandChild2 = child2.Child("grandChild2")
		})

		It("returns root, parent, child registry", func() {
			Expect(child1.Parent()).To(Equal(rootRegistry))
			Expect(child2.Parent()).To(Equal(rootRegistry))
			Expect(grandChild1.Parent()).To(Equal(child1))
			Expect(grandChild2.Parent()).To(Equal(child2))
			Expect(child1.Root()).To(Equal(rootRegistry))
			Expect(child2.Root()).To(Equal(rootRegistry))
			Expect(grandChild1.Root()).To(Equal(rootRegistry))
			Expect(grandChild2.Root()).To(Equal(rootRegistry))
			Expect(child1.Child("grandChild1")).To(Equal(grandChild1))
			Expect(child2.Child("grandChild2")).To(Equal(grandChild2))
			Expect(child1.ChildIfExists("grandChild1")).To(Equal(grandChild1))
			Expect(child2.ChildIfExists("grandChild2")).To(Equal(grandChild2))
			Expect(rootRegistry.Child("child1")).To(Equal(child1))
			Expect(rootRegistry.Child("child2")).To(Equal(child2))
			Expect(rootRegistry.ChildIfExists("child1")).To(Equal(child1))
			Expect(rootRegistry.ChildIfExists("child2")).To(Equal(child2))
		})
		It("returns status information", func() {
			test_status1 := status.NewStatus(nil, errors.New("test status1"))
			rootRegistry.SetStatus(test_status1)
			Expect(rootRegistry.GetStatus()).To(Equal(test_status1))
			Expect(rootRegistry.HasError()).To(BeTrue())

			// multiple_groupstatus := &statusv1.GroupStatus{
			// 	Status: rootRegistry.GetStatus(),
			// 	Groups: map[string]*statusv1.GroupStatus{
			// 		"child1": {
			// 			Status: child1.GetStatus(),
			// 			Groups: map[string]*statusv1.GroupStatus{
			// 				"grandChild1": {
			// 					Status: grandChild1.GetStatus(),
			// 					Groups: make(map[string]*statusv1.GroupStatus),
			// 				},
			// 			},
			// 		},
			// 		"child2": {
			// 			Status: child2.GetStatus(),
			// 			Groups: map[string]*statusv1.GroupStatus{
			// 				"grandChild2": {
			// 					Status: grandChild2.GetStatus(),
			// 					Groups: make(map[string]*statusv1.GroupStatus),
			// 				},
			// 			},
			// 		},
			// 	},
			// }
			// rootRegistry.SetGroupStatus(multiple_groupstatus)
			// Expect(rootRegistry.GetGroupStatus()).To(Equal(multiple_groupstatus))
			// Expect(rootRegistry.GetGroupStatus()).To(Equal(nil))
		})
		It("creates multiple child registries then detaches them", func() {
			grandChild3 := child1.Child("grandChild3")
			grandChild4 := child1.Child("grandChild4")
			ggrandChild1 := grandChild3.Child("ggrandChild1")
			ggrandChild2 := grandChild4.Child("ggrandChild2")
			Expect(child1.ChildIfExists("grandChild3")).To(Equal(grandChild3))
			Expect(child1.ChildIfExists("grandChild4")).To(Equal(grandChild4))
			Expect(grandChild3.ChildIfExists("ggrandChild1")).To(Equal(ggrandChild1))
			Expect(grandChild4.ChildIfExists("ggrandChild2")).To(Equal(ggrandChild2))

			child1.Child("grandChild3").Detach()
			Expect(child1.ChildIfExists("grandChild3")).To(BeNil())
			grandChild4.Child("ggrandChild2").Detach()
			Expect(grandChild4.ChildIfExists("ggrandChild2")).To(BeNil())
			child1.Child("grandChild4").Detach()
			Expect(child1.ChildIfExists("grandChild4")).To(BeNil())
		})
		It("returns the key of the registry that is registered with the parent", func() {
			Expect(child1.Key()).To(Equal("child1"))
			Expect(child2.Key()).To(Equal("child2"))
			Expect(grandChild1.Key()).To(Equal("grandChild1"))
			Expect(grandChild2.Key()).To(Equal("grandChild2"))
		})
	})
})