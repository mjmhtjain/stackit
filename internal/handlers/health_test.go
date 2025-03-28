package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestGetHealthDirectMock tests the GetHealth handler directly without Gin test mode
func TestGetHealthDirectMock(t *testing.T) {
	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	handler := NewHealthHandler()

	// Call the handler function directly
	handler.GetHealth(c)

	// Assert status code
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")

	// Parse the response body
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Should be able to parse response")

	// Assert response content
	assert.Equal(t, "ok", response["status"], "Response status should be 'ok'")
}
