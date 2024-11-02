package notification

import (
	"context"

	"github.com/knoxie-s/notification-service/internal/model"
)

func (s *serv) Create(ctx context.Context, info *model.NotificationInfo) (*model.Notification, error) {
	notification, err := s.notificationRepository.Create(ctx, info)
	if err != nil {
		return nil, err
	}

	return notification, nil
}
