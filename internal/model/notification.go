package model

import "time"

// Notification ...
type Notification struct {
	ID        int64            `json:"id"`
	Info      NotificationInfo `json:"info"`
	CreatedAt *time.Time       `json:"createdAt"`
	UpdatedAt *time.Time       `json:"updatedAt"`
}

// NotificationInfo ...
type NotificationInfo struct {
	OrderType  string `json:"orderType"`
	SessionID  string `json:"sessionID"`
	Card       string `json:"card"`
	WebsiteURL string `json:"websiteURL"`
	EventDate  string `json:"eventDate"`
}

// NotificationFilter ...
type NotificationFilter struct {
	SentToClient bool
	Limit        int
}
