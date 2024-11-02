package converter

import (
	modelAPI "github.com/knoxie-s/notification-service/internal/api/model"
	"github.com/knoxie-s/notification-service/internal/model"
)

func ToNotificationInfoFromAPI(notification *modelAPI.NotificationRequest) *model.NotificationInfo {
	return &model.NotificationInfo{
		OrderType:  notification.OrderType,
		SessionID:  notification.SessionID,
		Card:       notification.Card,
		WebsiteURL: notification.WebsiteURL,
		EventDate:  notification.EventDate,
	}
}
