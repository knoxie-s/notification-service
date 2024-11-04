package api

import "github.com/knoxie-s/notification-service/internal/service"

// Implementation API
type Implementation struct {
	notificationService service.NotificationService
}

// NewImplementation construct Implementation
func NewImplementation(notificationService service.NotificationService) *Implementation {
	return &Implementation{
		notificationService: notificationService,
	}
}
