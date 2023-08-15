package starshipcommsresolver

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"starshipCommsResolver/internal/pkg/entity"
)

func topsecret(c *gin.Context) {
	if c.Request.Method == "POST" {

		var satellite entity.Satellites
		if err := c.BindJSON(&satellite); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Input "})
			return
		}

	}
}
