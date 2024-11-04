package config

import (
	"errors"
	"os"
	"strconv"
)

const (
	workerCountEnvName             = "WORKER_COUNT"
	workerIntervalSecondsEnvName   = "WORKER_INTERVAL_IN_SECONDS"
	workerNotificationLimitEnvName = "WORKER_NOTIFICATION_SEND_LIMIT_PER_THREAD"
)

// WorkerConfig for worker config
type WorkerConfig interface {
	GetCount() int
	GetInterval() int
	GetNotificationLimit() int
}

type workerConfig struct {
	count             string
	interval          string
	notificationLimit string
}

// NewWorkerConfig construct worker config
func NewWorkerConfig() (WorkerConfig, error) {
	count := os.Getenv(workerCountEnvName)
	if len(count) == 0 {
		return nil, errors.New("http host not found")
	}

	interval := os.Getenv(workerIntervalSecondsEnvName)
	if len(interval) == 0 {
		return nil, errors.New("http port not found")
	}

	workerNotificationLimit := os.Getenv(workerNotificationLimitEnvName)
	if len(interval) == 0 {
		return nil, errors.New("http port not found")
	}

	return &workerConfig{
		count:             count,
		interval:          interval,
		notificationLimit: workerNotificationLimit,
	}, nil
}

func (cfg *workerConfig) GetCount() int {
	n, _ := strconv.Atoi(cfg.count)
	return n
}

func (cfg *workerConfig) GetInterval() int {
	n, _ := strconv.Atoi(cfg.interval)
	return n
}

func (cfg *workerConfig) GetNotificationLimit() int {
	n, _ := strconv.Atoi(cfg.notificationLimit)
	return n
}
