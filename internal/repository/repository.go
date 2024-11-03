package repository

import (
	"context"

	"github.com/knoxie-s/notification-service/internal/model"
)

type NotificationRepository interface {
	Create(ctx context.Context, notificationInfo *model.NotificationInfo) (*model.Notification, error)
	Get(ctx context.Context, filter model.NotificationFilter) ([]*model.Notification, error)
	UpdateSendToClientStatus(_ context.Context, id int64) error
}
