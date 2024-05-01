package main

import (
	"github.com/gin-gonic/gin"
	"github.com/williamzhou8818/time-fit-backend/handler"
)

func main() {
	//
	router := gin.Default()
	router.GET("/categorys", handler.CategoryLists)
	router.GET("/timslists", handler.TimsLists)

	router.Run("192.168.56.1:5300")
}
