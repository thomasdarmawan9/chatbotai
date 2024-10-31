package routes

import (
	"geminiaibot/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers the API routes with the Gin router
func RegisterRoutes(r *gin.Engine) {
	// Define the route for the chatbot controller
	r.POST("/generate", controllers.GenerateChatResponse)
}
