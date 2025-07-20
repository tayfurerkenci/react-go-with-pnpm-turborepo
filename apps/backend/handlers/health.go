package handlers

import (
	"net/http"
	"time"

	"backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetHealth(db *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check database connection
		status := "healthy"
		dbStatus := "connected"

		if err := db.Client().Ping(c, nil); err != nil {
			status = "unhealthy"
			dbStatus = "disconnected"
		}

		response := models.HealthResponse{
			Status:    status,
			Timestamp: time.Now(),
			Database:  dbStatus,
			Version:   "1.0.0",
		}

		if status == "healthy" {
			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(http.StatusServiceUnavailable, response)
		}
	}
}
