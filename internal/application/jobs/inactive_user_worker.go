package jobs

import (
	"context"
	"main/internal/config"
	"main/internal/domain/user"
	"main/pkg"
	"time"
)

type InactiveUserWorker struct {
	userService *user.UserService
	logger      pkg.Logger
	clearTTL    time.Duration
}

func NewInactiveUserWorker(
	userService *user.UserService,
	logger pkg.Logger,
	env config.Env,
) *InactiveUserWorker {
	return &InactiveUserWorker{
		userService: userService,
		logger:      logger,
		clearTTL:    env.DeleteInactiveUsersTTL,
	}
}

func (w *InactiveUserWorker) Run(ctx context.Context) {
	ticker := time.NewTicker(w.clearTTL)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			w.logger.Info("User clear worker stoped.")
			return
		case <-ticker.C:
			if err := w.userService.DeleteInactiveUsers(); err != nil {
				w.logger.Error(err.Error())
			} else {
				w.logger.Info("Inactive users are successfully deleted.")
			}
		}
	}
}
