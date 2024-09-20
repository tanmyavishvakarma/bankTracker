package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define a general POST route for the webhook
	r.POST("/webhook", func(c *gin.Context) {
		var json map[string]interface{} // Using a generic map to handle any JSON structure

		// Bind the JSON body to a generic map
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Log the received JSON payload
		fmt.Printf("Received payload: %v\n", json)

		// Respond back to the webhook sender
		c.JSON(http.StatusOK, gin.H{"status": "received"})
	})

	// Start the server on port 8080
	r.Run(":8080")
}
