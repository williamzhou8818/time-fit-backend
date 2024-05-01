package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/williamzhou8818/time-fit-backend/database"
)

func CategoryLists(ctx *gin.Context) {
	categoryList := database.GetAllCategory()
	if categoryList == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again late"})

	}

	ctx.JSON(200, categoryList)
}
