package nats

import (
	"main/internal/domain"

	"go.uber.org/fx"
)

type Nats struct {
	eventBus         *domain.EventBus
	userEventHandler *UserEventHandler
}

func NewNats(
	eventBus *domain.EventBus,
	userEventHandler *UserEventHandler,
) *Nats {
	return &Nats{
		eventBus:         eventBus,
		userEventHandler: userEventHandler,
	}
}

func (n *Nats) SetupPublishers() {
	n.eventBus.On("user.registered", n.userEventHandler.UserRegistered)
	n.eventBus.On("user.deleted", n.userEventHandler.UserDeleted)
	n.eventBus.On("user.updated", n.userEventHandler.UserUpdated)
}

var Module = fx.Options(
	fx.Provide(NewNats),
	fx.Provide(NewUserEventHandler),
)
