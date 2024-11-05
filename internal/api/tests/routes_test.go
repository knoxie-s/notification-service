package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/knoxie-s/notification-service/internal/api"
	modelAPI "github.com/knoxie-s/notification-service/internal/api/model"
	"github.com/knoxie-s/notification-service/internal/model"
	"github.com/knoxie-s/notification-service/internal/service"
	serviceMocks "github.com/knoxie-s/notification-service/internal/service/mocks"
	"github.com/stretchr/testify/require"
)

func TestCreateNotification(t *testing.T) {
	type notificationServiceMockFunc func(mc *minimock.Controller) service.NotificationService

	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}

	var (
		mc = minimock.NewController(t)

		id  = gofakeit.Int64()
		now = time.Now()
	)

	defer mc.Finish()

	notificationReq := modelAPI.NotificationRequest{
		OrderType:  "Purchase",
		SessionID:  "1234",
		Card:       "4433**1409",
		EventDate:  "2023-01-04T13:44:52Z",
		WebsiteURL: "https://amazon.com",
	}

	notificationInfo := &model.NotificationInfo{
		OrderType:  "Purchase",
		SessionID:  "1234",
		Card:       "4433**1409",
		EventDate:  "2023-01-04T13:44:52Z",
		WebsiteURL: "https://amazon.com",
	}

	notificationResp := &model.Notification{
		ID: id,
		Info: model.NotificationInfo{
			OrderType:  "Purchase",
			SessionID:  "1234",
			Card:       "4433**1409",
			EventDate:  "2023-01-04T13:44:52Z",
			WebsiteURL: "https://amazon.com",
		},
		CreatedAt: &now,
	}

	tests := []struct {
		name                        string
		args                        args
		expectedStatus              int
		expectedStatusText          string
		expectedMessage             string
		notificationServiceMockFunc notificationServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "/notification", toJSONBody(notificationReq)),
			},
			expectedStatus:     http.StatusCreated,
			expectedStatusText: "success",
			expectedMessage:    "notification received",
			notificationServiceMockFunc: func(mc *minimock.Controller) service.NotificationService {
				mock := serviceMocks.NewNotificationServiceMock(mc)
				mock.CreateMock.Expect(context.Background(), notificationInfo).Return(notificationResp, nil)
				return mock
			},
		},
		{
			name: "error case",
			args: args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "/notification", toJSONBody(notificationReq)),
			},
			expectedStatus:     http.StatusInternalServerError,
			expectedStatusText: "error",
			expectedMessage:    "internal server error",
			notificationServiceMockFunc: func(mc *minimock.Controller) service.NotificationService {
				mock := serviceMocks.NewNotificationServiceMock(mc)
				mock.CreateMock.Expect(context.Background(), notificationInfo).Return(nil, errors.New("service error"))
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			notificationServiceMock := tt.notificationServiceMockFunc(mc)

			API := api.NewImplementation(notificationServiceMock)

			API.CreateNotification(tt.args.w, tt.args.req)

			resp := tt.args.w.(*httptest.ResponseRecorder).Result()
			defer resp.Body.Close()

			require.Equal(t, tt.expectedStatus, resp.StatusCode)

			if resp.StatusCode == http.StatusInternalServerError {
				return
			}

			var response modelAPI.Response

			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				t.Fatalf("Expected valid JSON response, got error: %v", err)
			}

			require.Equal(t, tt.expectedStatusText, response.Status)
			require.Equal(t, tt.expectedMessage, response.Message)
		})
	}
}

func toJSONBody(data interface{}) *bytes.Reader {
	body, _ := json.Marshal(data)
	return bytes.NewReader(body)
}
