package workflows

import (
	"github.com/xakepp35/hello-temporal/internal/activities"
	"go.temporal.io/sdk/workflow"
)

// Workflow: выполняет единственную активность и завершает работу
func HelloWorld(ctx workflow.Context) error {
	opts := workflow.ActivityOptions{
		ScheduleToStartTimeout: DefaultScheduleToStartTimeout,
		StartToCloseTimeout:    DefaultStartToCloseTimeout,
	}
	ctx = workflow.WithActivityOptions(ctx, opts)

	promise := workflow.ExecuteActivity(ctx, activities.HelloWorld)
	return promise.Get(ctx, nil)
}
