package api

import (
	"crypto/rand"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"sync"
	"time"
)

// Notification data
type Notification struct {
	ID        int64            `json:"id"`
	Info      NotificationInfo `json:"info"`
	CreatedAt *time.Time       `json:"createdAt"`
	UpdatedAt *time.Time       `json:"updatedAt"`
}

type NotificationInfo struct {
	OrderType  string `json:"orderType"`
	SessionID  string `json:"sessionId"`
	Card       string `json:"card"`
	WebsiteURL string `json:"websiteUrl"`
	EventDate  string `json:"eventDate"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Cache to store Notification data
type Cache struct {
	elems map[int64]*Notification
	m     sync.Mutex
}

var notificationCache = &Cache{
	elems: make(map[int64]*Notification),
}

func handleNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var notificationInfo NotificationInfo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&notificationInfo); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	n, err := rand.Int(rand.Reader, big.NewInt(27))
	if err != nil {
		log.Printf("Failed to generate user id: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	id := n.Int64()
	now := time.Now()

	notification := &Notification{
		ID:        id,
		Info:      notificationInfo,
		CreatedAt: &now,
	}

	notificationCache.m.Lock()
	notificationCache.elems[notification.ID] = notification
	notificationCache.m.Unlock()

	//response := fmt.Sprintf("Received Notification: %+v", notification)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := Response{
		Status:  "success",
		Message: "notification received",
		Data:    notification,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode note data", http.StatusInternalServerError)
		return
	}
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/notification", handleNotification)
}
