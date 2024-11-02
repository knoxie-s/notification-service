package notification

import (
	"github.com/knoxie-s/notification-service/internal/repository"
	"github.com/knoxie-s/notification-service/internal/service"
)

type serv struct {
	notificationRepository repository.NotificationRepository
}

func NewService(
	notificationRepository repository.NotificationRepository,
) service.NotificationService {
	return &serv{
		notificationRepository: notificationRepository,
	}
}
