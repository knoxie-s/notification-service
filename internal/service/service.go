package service

import (
	"context"

	"github.com/knoxie-s/notification-service/internal/model"
)

// NotificationService service
type NotificationService interface {
	Create(ctx context.Context, notificationInfo *model.NotificationInfo) (*model.Notification, error)
	SendNotificationsToClients(ctx context.Context) error
}
