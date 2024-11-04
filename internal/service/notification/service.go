package notification

import (
	"github.com/knoxie-s/notification-service/internal/config"
	"github.com/knoxie-s/notification-service/internal/repository"
	"github.com/knoxie-s/notification-service/internal/service"
)

type serv struct {
	notificationRepository repository.NotificationRepository
	workerConfig           config.WorkerConfig
}

// NewService construct NotificationService
func NewService(
	notificationRepository repository.NotificationRepository,
	workerConfig config.WorkerConfig,
) service.NotificationService {
	return &serv{
		notificationRepository: notificationRepository,
		workerConfig:           workerConfig,
	}
}
