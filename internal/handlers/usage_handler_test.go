package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/stackit/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestUsageHandler_PostUsage(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		payload    interface{}
		wantStatus int
		wantError  bool
		errorMsg   string
	}{
		{
			name: "Valid request",
			payload: dto.UsageDTO{
				InstanceID: "test-instance",
				SKUID:      "test-sku",
				Timestamp:  "2023-01-01T12:00:00Z",
			},
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name: "Missing instance_id",
			payload: dto.UsageDTO{
				SKUID:     "test-sku",
				Timestamp: "2023-01-01T12:00:00Z",
			},
			wantStatus: http.StatusBadRequest,
			wantError:  true,
			errorMsg:   "instance_id is required",
		},
		{
			name: "Missing skuid",
			payload: dto.UsageDTO{
				InstanceID: "test-instance",
				Timestamp:  "2023-01-01T12:00:00Z",
			},
			wantStatus: http.StatusBadRequest,
			wantError:  true,
			errorMsg:   "skuid is required",
		},
		{
			name: "Missing timestamp",
			payload: dto.UsageDTO{
				InstanceID: "test-instance",
				SKUID:      "test-sku",
			},
			wantStatus: http.StatusBadRequest,
			wantError:  true,
			errorMsg:   "timestamp is required",
		},
		{
			name:       "Invalid JSON",
			payload:    "invalid-json",
			wantStatus: http.StatusBadRequest,
			wantError:  true,
			errorMsg:   "Invalid request body",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a response recorder
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Create request body
			var reqBody []byte
			var err error

			if s, ok := tt.payload.(string); ok {
				// Handle the invalid JSON case
				reqBody = []byte(s)
			} else {
				reqBody, err = json.Marshal(tt.payload)
				assert.NoError(t, err)
			}

			// Set up the request
			c.Request = httptest.NewRequest(http.MethodPost, "/usage", bytes.NewBuffer(reqBody))
			c.Request.Header.Set("Content-Type", "application/json")

			// Initialize the handler and call the method
			handler := NewUsageHandler()
			handler.PostUsage(c)

			// Assert status code
			assert.Equal(t, tt.wantStatus, w.Code)

			// Parse response
			var response map[string]interface{}
			err = json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			// Check for expected error or success message
			if tt.wantError {
				assert.Contains(t, response, "error")
				assert.Equal(t, tt.errorMsg, response["error"])
			} else {
				assert.Contains(t, response, "message")
				assert.Equal(t, "Usage data received successfully", response["message"])
			}
		})
	}
}
