package infrastructure

import (
	domain "main/internal/domain/user"
	"main/pkg"
)

type MockPublisher struct {
	logger pkg.Logger
}

func NewMockPublisher(logger pkg.Logger) domain.Publisher {
	return &MockPublisher{logger: logger}
}

func (p *MockPublisher) UserActivated(userId string) error {
	p.logger.Info("user activated: " + userId)
	return nil
}

func (p *MockPublisher) UserDeleted(userId string) error {
	p.logger.Info("user deleted: " + userId)
	return nil
}
