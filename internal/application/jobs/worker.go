package jobs

import (
	"context"

	"go.uber.org/fx"
)

type Worker interface {
	Run(ctx context.Context)
}

type Workers []Worker

func NewWorkers(
	tokenClearWorker *TokenClearWorker,
	inactiveUserWorker *InactiveUserWorker,
) Workers {
	return Workers{
		tokenClearWorker,
		inactiveUserWorker,
	}
}

func (w Workers) Run(ctx context.Context) {
	for _, worker := range w {
		go worker.Run(ctx)
	}
}

var Module = fx.Options(
	fx.Provide(NewWorkers),
	fx.Provide(NewTokenClearWorker),
	fx.Provide(NewInactiveUserWorker),
)
