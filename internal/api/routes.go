package api

import (
	"encoding/json"
	"net/http"

	"github.com/knoxie-s/notification-service/internal/api/model"
	"github.com/knoxie-s/notification-service/internal/converter"
)

// RegisterRoutes http
func RegisterRoutes(mux *http.ServeMux, impl *Implementation) {
	mux.HandleFunc("/notification", impl.CreateNotification)
}

// CreateNotification handle notification
func (impl *Implementation) CreateNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var notificationRequest model.NotificationRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&notificationRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	notification, err := impl.notificationService.Create(r.Context(), converter.ToNotificationInfoFromAPI(&notificationRequest))
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := model.Response{
		Status:  "success",
		Message: "notification received",
		Data:    notification,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode notification data", http.StatusInternalServerError)
		return
	}
}
