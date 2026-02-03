package jobs

import "go.uber.org/fx"

type Worker interface {
	Run()
}

type Workers []Worker

func NewWorkers() Workers {
	return Workers{}
}

func (w Workers) Run() {
	for _, worker := range w {
		worker.Run()
	}
}

var Module = fx.Options(
	fx.Provide(NewWorkers),
)
