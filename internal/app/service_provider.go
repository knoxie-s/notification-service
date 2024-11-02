package app

import (
	"context"
	"log"

	"github.com/knoxie-s/notification-service/internal/api"
	"github.com/knoxie-s/notification-service/internal/config"
	"github.com/knoxie-s/notification-service/internal/repository"
	notificationRepository "github.com/knoxie-s/notification-service/internal/repository/notification"
	"github.com/knoxie-s/notification-service/internal/service"
	notificationService "github.com/knoxie-s/notification-service/internal/service/notification"
)

type serviceProvider struct {
	httpConfig config.HTTPConfig

	notificationRepository repository.NotificationRepository
	notificationService    service.NotificationService
	notificationImpl       *api.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) NotificationRepository(ctx context.Context) repository.NotificationRepository {
	if s.notificationRepository == nil {
		s.notificationRepository = notificationRepository.NewRepository()
	}

	return s.notificationRepository
}

func (s *serviceProvider) NotificationService(ctx context.Context) service.NotificationService {
	if s.notificationService == nil {
		s.notificationService = notificationService.NewService(s.NotificationRepository(ctx))
	}

	return s.notificationService
}

func (s *serviceProvider) NotificationAPI(ctx context.Context) *api.Implementation {
	if s.notificationImpl == nil {
		s.notificationImpl = api.NewImplementation(s.NotificationService(ctx))
	}

	return s.notificationImpl
}
