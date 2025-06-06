package routes

import (
	"bluebell/controller"
	"bluebell/middlewares"
	"bluebell/pkg/jwt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Mock middleware for JWT authentication
func mockJWTAuthMiddleware(success bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if success {
			// Simulate a valid token by setting a mock user ID
			c.Set("userID", int64(12345))
			c.Next()
		} else {
			// Simulate an invalid token
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Invalid token"})
		}
	}
}

// Mock controller for UserInfoHandler
func mockUserInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"userID": uid})
	}
}

// go test -v -run ^TestUserInfoRoute$ your_package/routes
func TestUserInfoRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Test cases
	tests := []struct {
		name           string
		mockMiddleware gin.HandlerFunc
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Success",
			mockMiddleware: mockJWTAuthMiddleware(true),
			expectedStatus: http.StatusOK,
			expectedBody:   `{"userID":12345}`,
		},
		{
			name:           "Failure",
			mockMiddleware: mockJWTAuthMiddleware(false),
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"msg":"Invalid token"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new Gin engine
			r := gin.New()

			// Register the route with mock middleware and handler
			r.GET("/userinfo", tt.mockMiddleware, mockUserInfoHandler())

			// Create a test request
			req, _ := http.NewRequest(http.MethodGet, "/userinfo", nil)
			w := httptest.NewRecorder()

			// Perform the request
			r.ServeHTTP(w, req)

			// Assert the response
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

// go test -v -run ^TestUserInfoRoute_WithRealToken$ your_package/routes
func TestUserInfoRoute_WithRealToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin engine with the actual middleware and handler
	r := gin.New()
	r.Use(middlewares.JWTAuthMiddleware())           // Use the real JWT middleware
	r.GET("/userinfo", controller.UserInfoHandler()) // Use the real handler

	// Generate a valid token (use the same logic as your project)
	token, err := jwt.GenToken(12345) // Replace with your actual token generation logic
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Test cases
	tests := []struct {
		name           string
		authHeader     string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Success with valid token",
			authHeader:     "Bearer " + token, // Include the valid token
			expectedStatus: http.StatusOK,
			expectedBody:   `{"code":1000,"data":12345,"msg":"成功"}`,
		},
		{
			name:           "Failure with invalid token",
			authHeader:     "Bearer invalid_token", // Simulate an invalid token
			expectedStatus: http.StatusOK,
			expectedBody:   `{"code":1007,"data":null,"msg":"无效的Token"}`,
		},
		{
			name:           "Failure with missing token",
			authHeader:     "", // No Authorization header
			expectedStatus: http.StatusOK,
			expectedBody:   `{"code":1006,"data":null,"msg":"请求头中auth为空"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test request
			req, _ := http.NewRequest(http.MethodGet, "/userinfo", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader) // Set the Authorization header
			}
			w := httptest.NewRecorder()

			// Perform the request
			r.ServeHTTP(w, req)

			// Assert the response
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}
