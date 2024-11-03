package model

import "time"

// Notification data
type Notification struct {
	ID        int64
	Info      NotificationInfo
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// NotificationInfo data
type NotificationInfo struct {
	OrderType    string
	SessionID    string
	Card         string
	WebsiteURL   string
	EventDate    string
	SentToClient bool
}
