package main

import "net/http"

import (
	"async-bmstu/internal/api"
	"github.com/gin-gonic/gin"
)

type PK struct {
	PK       int `json:"id"`
	Quantity int `json:"quantity"`
}

func main() {
	r := gin.Default()
	r.POST("/calculate_losses", func(c *gin.Context) {
		var request PK
		if err := c.BindJSON(&request); err != nil {
			// DO SOMETHING WITH THE ERROR
		}
		// Отправка PUT-запроса к основному серверу
		url := "http://127.0.0.1:8000/applications/QuantityOfLosses" // Замените на ваш реальный URL
		go api.SendStatus(api.PK(request), url)
		c.JSON(http.StatusOK, gin.H{"message": "Status update initiated"})

	})

	r.Run()
}
