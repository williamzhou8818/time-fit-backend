package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/williamzhou8818/time-fit-backend/database"
)

func TimsLists(ctx *gin.Context) {
	timsLists := database.GetTimsListByCateID()
	if timsLists == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again late"})

	}

	ctx.JSON(200, timsLists)
}
