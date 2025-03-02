package main

import (
	"github.com/rs/zerolog/log"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/xakepp35/hello-temporal/internal/activities"
	"github.com/xakepp35/hello-temporal/internal/temporalkit"
	"github.com/xakepp35/hello-temporal/internal/workflows"
	"github.com/xakepp35/pkg/env"
)

func main() {
	temporalkit.RunWorker(&temporalkit.Options{
		TaskQueue: env.Get("TEMPORAL_TASK_QUEUE", ""),
		ClientOpts: client.Options{
			HostPort: env.Get("TEMPORAL_ADDR", ""),
			Logger:   temporalkit.NewZerolog(&log.Logger),
		},
		WorkerOpts: worker.Options{},
		Workflows: []any{
			workflows.HelloWorld,
		},
		Activities: []any{
			activities.HelloWorld,
		},
	})
}
