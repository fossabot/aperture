package controlplane

import (
	"fmt"

	"go.uber.org/fx"

	policylangv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/policy/language/v1"
	"github.com/fluxninja/aperture/pkg/policies/controlplane/components/actuators/concurrency"
	"github.com/fluxninja/aperture/pkg/policies/controlplane/iface"
	"github.com/fluxninja/aperture/pkg/policies/controlplane/runtime"
)

// componentStackFactoryModuleForPolicyApp for component factory run via the policy app. For singletons in the Policy scope.
func componentStackFactoryModuleForPolicyApp(circuitAPI runtime.CircuitAPI) fx.Option {
	return fx.Options()
}

// newComponentStackAndOptions creates components for component stack, sub components and their fx options.
func newComponentStackAndOptions(
	componentStackProto *policylangv1.Component,
	componentStackIndex int,
	policyReadAPI iface.Policy,
) (runtime.CompiledComponent, []runtime.CompiledComponent, fx.Option, error) {
	// Factory parser to determine what kind of component stack to create
	if concurrencyLimiterProto := componentStackProto.GetConcurrencyLimiter(); concurrencyLimiterProto != nil {
		var (
			compiledComponents []runtime.CompiledComponent
			options            []fx.Option
		)
		concurrencyLimiterOptions, agentGroupName, concurrencyLimiterErr := concurrency.NewConcurrencyLimiterOptions(concurrencyLimiterProto, componentStackIndex, policyReadAPI)
		if concurrencyLimiterErr != nil {
			return runtime.CompiledComponent{}, nil, nil, concurrencyLimiterErr
		}
		// Append concurrencyLimiter options
		options = append(options, concurrencyLimiterOptions)

		// Scheduler
		if schedulerProto := concurrencyLimiterProto.GetScheduler(); schedulerProto != nil {
			scheduler, schedulerOptions, schedulerErr := concurrency.NewSchedulerAndOptions(schedulerProto, componentStackIndex, policyReadAPI, agentGroupName)
			if schedulerErr != nil {
				return runtime.CompiledComponent{}, nil, nil, schedulerErr
			}
			schedulerMapStruct, err := encodeMapStruct(schedulerProto)
			if err != nil {
				return runtime.CompiledComponent{}, nil, nil, err
			}
			// Append scheduler as a runtime.CompiledComponent
			compiledComponents = append(compiledComponents, runtime.CompiledComponent{
				Component:     scheduler,
				MapStruct:     schedulerMapStruct,
				Name:          "Scheduler",
				ComponentType: runtime.ComponentTypeSource,
			})

			// Append scheduler options
			options = append(options, schedulerOptions)
		}

		// Actuation Strategy
		if loadShedActuatorProto := concurrencyLimiterProto.GetLoadShedActuator(); loadShedActuatorProto != nil {
			loadShedActuator, loadShedActuatorOptions, loadShedActuatorErr := concurrency.NewLoadShedActuatorAndOptions(loadShedActuatorProto, componentStackIndex, policyReadAPI, agentGroupName)
			if loadShedActuatorErr != nil {
				return runtime.CompiledComponent{}, nil, nil, loadShedActuatorErr
			}
			loadShedActuatorMapStruct, err := encodeMapStruct(loadShedActuatorProto)
			if err != nil {
				return runtime.CompiledComponent{}, nil, nil, err
			}
			// Append loadShedActuator as a runtime.CompiledComponent
			compiledComponents = append(compiledComponents, runtime.CompiledComponent{
				Component:     loadShedActuator,
				MapStruct:     loadShedActuatorMapStruct,
				Name:          "LoadShedActuator",
				ComponentType: runtime.ComponentTypeSink,
			})

			// Append loadShedActuator options
			options = append(options, loadShedActuatorOptions)
		}

		concurrencyLimiterMapStruct, err := encodeMapStruct(concurrencyLimiterProto)
		if err != nil {
			return runtime.CompiledComponent{}, nil, nil, err
		}
		return runtime.CompiledComponent{
			Component:     nil,
			MapStruct:     concurrencyLimiterMapStruct,
			Name:          "ConcurrencyLimiter",
			ComponentType: runtime.ComponentTypeStandAlone,
		}, compiledComponents, fx.Options(options...), nil
	}
	return runtime.CompiledComponent{}, nil, nil, fmt.Errorf("unsupported/missing component type")
}
