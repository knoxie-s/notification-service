package model

type NotificationRequest struct {
	OrderType  string `json:"orderType"`
	SessionID  string `json:"sessionID"`
	Card       string `json:"card"`
	WebsiteURL string `json:"websiteURL"`
	EventDate  string `json:"eventDate"`
}
