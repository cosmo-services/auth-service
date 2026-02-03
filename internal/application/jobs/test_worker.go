package jobs

import (
	"context"
	"main/pkg"
	"time"
)

type TestWorker struct {
	logger pkg.Logger
}

func NewTestWorker(logger pkg.Logger) *TestWorker {
	return &TestWorker{
		logger: logger,
	}
}

func (w *TestWorker) Run(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			w.logger.Info("TestWorker: Received stop signal")
			return
		case <-ticker.C:
			w.logger.Info("I'M WORKING!")
		}
	}
}
