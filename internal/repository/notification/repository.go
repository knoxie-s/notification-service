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
	m     sync.Mutex
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
