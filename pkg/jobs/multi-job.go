package jobs

import (
	"context"
	"errors"
	"sync"

	"go.uber.org/fx"
	"google.golang.org/protobuf/proto"

	"github.com/fluxninja/aperture/pkg/config"
	"github.com/fluxninja/aperture/pkg/log"
	"github.com/fluxninja/aperture/pkg/panichandler"
	"github.com/fluxninja/aperture/pkg/status"
)

// MultiJobConfig holds configuration for MultiJob.
// swagger:model
type MultiJobConfig struct {
	JobConfig
	// Sets whether the job is always passing
	AlwaysHealthy bool `json:"always_healthy" default:"false"`
}

// MultiJobConstructor holds fields to create annotated instance of MultiJob.
type MultiJobConstructor struct {
	Name          string
	JobGroupName  string
	JWS           JobWatchers
	GWS           GroupWatchers
	DefaultConfig MultiJobConfig
}

// Annotate provides annotated instance of MultiJob.
func (mjc MultiJobConstructor) Annotate() fx.Option {
	name := config.NameTag(mjc.JobGroupName + "." + mjc.Name)
	group := config.GroupTag(mjc.Name)
	jgName := config.NameTag(mjc.JobGroupName)
	return fx.Provide(
		fx.Annotate(
			mjc.provideMultiJob,
			fx.ParamTags(group, group, jgName),
			fx.ResultTags(name),
		),
	)
}

func (mjc MultiJobConstructor) provideMultiJob(
	gws GroupWatchers,
	jws JobWatchers,
	jg *JobGroup,
	unmarshaller config.Unmarshaller,
	lifecycle fx.Lifecycle,
) (*MultiJob, error) {
	config := mjc.DefaultConfig

	if err := unmarshaller.UnmarshalKey(mjc.Name, &config); err != nil {
		log.Error().Err(err).Msg("Unable to deserialize")
		return nil, err
	}

	gwAll := GroupWatchers{}

	if len(mjc.GWS) > 0 || len(gws) > 0 {
		gwAll = append(gwAll, mjc.GWS...)
		gwAll = append(gwAll, gws...)
	}

	jwAll := JobWatchers{}
	if len(mjc.JWS) > 0 || len(jws) > 0 {
		jwAll = append(jwAll, mjc.JWS...)
		jwAll = append(jwAll, jws...)
	}

	// Create a new MultiJob instance
	mj := NewMultiJob(mjc.Name, config.AlwaysHealthy, jwAll, gwAll)

	lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			// Register multijob
			err := jg.RegisterJob(mj, config.JobConfig)
			if err != nil {
				return err
			}
			return nil
		},
		OnStop: func(context.Context) error {
			// Deregister all jobs
			mj.gt.reset()
			_ = jg.DeregisterJob(mjc.Name)
			return nil
		},
	})

	return mj, nil
}

// MultiJob runs multiple jobs in asynchronous manner.
type MultiJob struct {
	gt *groupTracker
	JobBase
	alwaysHealthy bool // Always marks the jobs as healthy even when Jobs fail.
}

// Make sure MultiJob complies with Job interface.
var _ Job = (*MultiJob)(nil)

// NewMultiJob creates a new instance of MultiJob.
func NewMultiJob(name string, alwaysHealthy bool, jws JobWatchers, gws GroupWatchers) *MultiJob {
	return &MultiJob{
		JobBase: JobBase{
			JobName: name,
			JWS:     jws,
		},
		alwaysHealthy: alwaysHealthy,
		gt:            newGroupTracker(gws, status.NewRegistry()),
	}
}

// Name returns the name of the job.
func (mj *MultiJob) Name() string {
	return mj.JobBase.Name()
}

// JobWatchers returns the list of job watchers.
func (mj *MultiJob) JobWatchers() JobWatchers {
	return mj.JobBase.JobWatchers()
}

// Execute executes all jobs, collects that results, and returns the aggregated status.
func (mj *MultiJob) Execute(ctx context.Context) (proto.Message, error) {
	jobs := mj.gt.getJobs()

	var wg sync.WaitGroup
	for _, job := range jobs {
		wg.Add(1)

		execFunc := func(job Job) func() {
			return func() {
				defer wg.Done()
				_, _ = mj.gt.execute(ctx, job)
			}
		}
		panichandler.Go(execFunc(job))
	}
	// wait for results
	wg.Wait()

	results, healthy := mj.gt.results()

	var err error
	if healthy || mj.alwaysHealthy {
		err = nil
	} else {
		err = errors.New("one or more jobs are unhealthy")
	}

	return results, err
}

// RegisterJob registers a job with the MultiJob.
func (mj *MultiJob) RegisterJob(job Job) error {
	return mj.gt.registerJob(job)
}

// DeregisterJob deregisters a job with the MultiJob.
func (mj *MultiJob) DeregisterJob(name string) error {
	_, err := mj.gt.deregisterJob(name)
	return err
}

// DeregisterAll removes all jobs from the MultiJob.
func (mj *MultiJob) DeregisterAll() {
	mj.gt.reset()
}
