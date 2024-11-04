package notification

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/knoxie-s/notification-service/internal/model"
)

func (s *serv) SendNotificationsToClients(ctx context.Context) error {
	filter := model.NotificationFilter{
		SentToClient: false,
		Limit:        s.workerConfig.GetNotificationLimit(),
	}

	notifications, err := s.notificationRepository.Get(ctx, filter)
	if err != nil {
		return err
	}

	fmt.Println(notifications)

	wg := sync.WaitGroup{}
	sem := make(chan struct{}, s.workerConfig.GetCount())

	for _, ntfs := range notifications {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			wg.Add(1)
			sem <- struct{}{}

			go func(n *model.Notification) {
				defer wg.Done()
				defer func() { <-sem }()

				err = s.notificationRepository.UpdateSendToClientStatus(ctx, n.ID)
				if err != nil {
					log.Printf("failed to update notification send_to_client status: %v", err)
					return
				}

				log.Printf("Mock: notification %v has been sent to client", ntfs.Info)
			}(ntfs)
		}
	}

	wg.Wait()
	return nil
}
