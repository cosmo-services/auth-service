package jobs

import (
	"context"
	"main/internal/config"
	"main/internal/domain/tokens"
	"main/pkg"
	"time"
)

type TokenClearWorker struct {
	tokenService *tokens.TokenService
	logger       pkg.Logger
	clearTTL     time.Duration
}

func NewTokenClearWorker(
	logger pkg.Logger,
	env config.Env,
	tokenService *tokens.TokenService,
) *TokenClearWorker {
	return &TokenClearWorker{
		logger:       logger,
		clearTTL:     env.ClearExpiredTokensTTL,
		tokenService: tokenService,
	}
}

func (w *TokenClearWorker) Run(ctx context.Context) {
	ticker := time.NewTicker(w.clearTTL)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			w.logger.Info("Token clear worker stoped.")
			return
		case <-ticker.C:
			if err := w.tokenService.ClearExpiredTokens(); err != nil {
				w.logger.Error(err.Error())
			} else {
				w.logger.Info("Expired tokens are successfully deleted.")
			}
		}
	}
}
