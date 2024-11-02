package converter

import (
	"github.com/knoxie-s/notification-service/internal/model"
	modelRepo "github.com/knoxie-s/notification-service/internal/repository/notification/model"
)

func ToNotificationFromRepo(notification *modelRepo.Notification) *model.Notification {
	return &model.Notification{
		ID:        notification.ID,
		Info:      ToNotificationInfoFromRepo(notification.Info),
		CreatedAt: notification.CreatedAt,
		UpdatedAt: notification.UpdatedAt,
	}
}

func ToNotificationInfoFromRepo(info modelRepo.NotificationInfo) model.NotificationInfo {
	return model.NotificationInfo{
		OrderType:  info.OrderType,
		SessionID:  info.SessionID,
		Card:       info.Card,
		WebsiteURL: info.WebsiteURL,
		EventDate:  info.EventDate,
	}
}

func ToNotificationInfoFromService(info model.NotificationInfo) modelRepo.NotificationInfo {
	return modelRepo.NotificationInfo{
		OrderType:  info.OrderType,
		SessionID:  info.SessionID,
		Card:       info.Card,
		WebsiteURL: info.WebsiteURL,
		EventDate:  info.EventDate,
	}
}
