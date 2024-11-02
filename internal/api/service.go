package api

import "github.com/knoxie-s/notification-service/internal/service"

type Implementation struct {
	notificationService service.NotificationService
}

func NewImplementation(notificationService service.NotificationService) *Implementation {
	return &Implementation{
		notificationService: notificationService,
	}
}
