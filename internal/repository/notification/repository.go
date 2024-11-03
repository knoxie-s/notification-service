package notification

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/knoxie-s/notification-service/internal/model"
	"github.com/knoxie-s/notification-service/internal/repository"
	"github.com/knoxie-s/notification-service/internal/repository/notification/converter"
	modelRepo "github.com/knoxie-s/notification-service/internal/repository/notification/model"
)

type repo struct {
}

func NewRepository() repository.NotificationRepository {
	return &repo{}
}

// Cache to store Notification data
type Cache struct {
	elems map[int64]*modelRepo.Notification
	m     sync.RWMutex
}

var notificationCache = &Cache{
	elems: make(map[int64]*modelRepo.Notification),
}

func (r *repo) Create(ctx context.Context, info *model.NotificationInfo) (*model.Notification, error) {
	if info == nil {
		return nil, fmt.Errorf("please provide notification info")
	}

	n, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		return nil, fmt.Errorf("failed to genererate id: %v", err.Error())
	}

	id := n.Int64()
	now := time.Now()

	notification := &modelRepo.Notification{
		ID:        id,
		Info:      converter.ToNotificationInfoFromService(*info),
		CreatedAt: &now,
	}

	notificationCache.m.Lock()
	notificationCache.elems[notification.ID] = notification
	notificationCache.m.Unlock()

	return converter.ToNotificationFromRepo(notification), nil
}

// Get simple notification filter with no search optimization
func (r *repo) Get(_ context.Context, filter model.NotificationFilter) ([]*model.Notification, error) {
	notifications := make([]*model.Notification, 0, filter.Limit)

	notificationCache.m.RLock()
	defer notificationCache.m.RUnlock()

	for _, ntf := range notificationCache.elems {
		switch {
		case filter.Limit == 0:
			break
		case ntf.Info.SentToClient == false:
			notifications = append(notifications, converter.ToNotificationFromRepo(ntf))
			filter.Limit--
		default:
		}
	}

	return notifications, nil
}
