package nats

import (
	"context"
	"main/internal/domain"
	"main/pkg"
	"time"
)

type UserEventHandler struct {
	natsClient *pkg.NatsClient
	logger     pkg.Logger
}

func NewUserEventHandler(
	natsClient *pkg.NatsClient,
	logger pkg.Logger,
) *UserEventHandler {
	return &UserEventHandler{
		natsClient: natsClient,
		logger:     logger,
	}
}

func (p *UserEventHandler) UserRegistered(event domain.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := p.natsClient.PublishJSON(ctx, "auth.user.registered", event)
	if err != nil {
		p.logger.Error(err)

		return err
	}

	p.logger.Infof("User event successfully delivered: %s", event)

	return nil
}

func (p *UserEventHandler) UserActivated(event domain.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := p.natsClient.PublishJSON(ctx, "auth.user.activated", event)
	if err != nil {
		p.logger.Error(err)

		return err
	}

	return nil
}

func (p *UserEventHandler) UserDeleted(event domain.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := p.natsClient.PublishJSON(ctx, "auth.user.deleted", event)
	if err != nil {
		p.logger.Error(err)

		return err
	}

	return nil
}

func (p *UserEventHandler) UserEmailChanged(event domain.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := p.natsClient.PublishJSON(ctx, "auth.user.email.changed", event)
	if err != nil {
		p.logger.Error(err)

		return err
	}

	return nil
}

func (p *UserEventHandler) UserUsernameChanged(event domain.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := p.natsClient.PublishJSON(ctx, "auth.user.username.changed", event)
	if err != nil {
		p.logger.Error(err)

		return err
	}

	return nil
}
