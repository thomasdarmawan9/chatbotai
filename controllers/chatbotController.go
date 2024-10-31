package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"geminiaibot/models"

	"github.com/gin-gonic/gin"
)

// API key is stored as an environment variable for security
var apiKey = "AIzaSyBddFG13fDIutMOLj_3L5Ap3odr31eYmaw"

// GenerateChatResponse handles the chat response from the Gemini API
func GenerateChatResponse(c *gin.Context) {
	// Extract user input from the JSON body
	var requestData struct {
		Message string `json:"message"`
	}
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// API URL
	apiUrl := "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash-latest:generateContent?key=" + apiKey

	// Prepare payload for the Gemini API request
	payload := fmt.Sprintf(`{
		"contents": [{
			"parts": [{"text": "%s"}]
		}]
	}`, requestData.Message)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send the request to the Gemini API
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
		return
	}
	defer resp.Body.Close()

	// Decode the JSON response into the Response struct
	var response models.Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
		return
	}

	// Check and extract the text from the response
	if len(response.Candidates) > 0 && len(response.Candidates[0].Content.Parts) > 0 {
		c.JSON(http.StatusOK, gin.H{"response_text": response.Candidates[0].Content.Parts[0].Text})
	} else {
		c.JSON(http.StatusOK, gin.H{"response_text": "No text content found in response."})
	}
}
