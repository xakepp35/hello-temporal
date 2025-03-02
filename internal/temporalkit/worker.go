package temporalkit

import (
	"github.com/rs/zerolog/log"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

type Options struct {
	TaskQueue  string
	ClientOpts client.Options
	WorkerOpts worker.Options
	Workflows  []any
	Activities []any
}

func RunWorker(opts *Options) {
	// Подключение к серверу Temporal
	cli, err := client.Dial(opts.ClientOpts)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("client.Dial")
		return
	}
	defer cli.Close()
	// Настройка воркера
	wrk := worker.New(cli, opts.TaskQueue, opts.WorkerOpts)
	for _, workflow := range opts.Workflows {
		wrk.RegisterWorkflow(workflow)
	}
	for _, activity := range opts.Activities {
		wrk.RegisterActivity(activity)
	}
	// Запуск воркера
	err = wrk.Run(worker.InterruptCh())
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("wrk.Run")
		return
	}
	log.Info().
		Msg("wrk.Run")
}
