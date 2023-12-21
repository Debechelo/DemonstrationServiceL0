package rest

import (
	"DemonstrationServiceL0/internal/caching"
	"DemonstrationServiceL0/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func StartServer() {
	router := gin.Default()

	// Обработчик для получения данных заказа по ID
	router.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		var OrderId int

		// Преобразование параметра id в число
		if _, err := fmt.Sscanf(id, "%d", &OrderId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
			return
		}

		// Поиск заказа по ID в кеше
		item, found := caching.GetCache(OrderId)
		order, ok := item.(service.Order)
		if !ok {
			log.Println("Error to order!")
		}
		if !found {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}

		// Отправка данных заказа в формате JSON
		c.JSON(http.StatusOK, order)
	})

	log.Println("Connected to Server!")
	// Запускаем сервер
	router.Run(":8080")
}
