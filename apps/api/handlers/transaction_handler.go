package handlers

import (
	"api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetTransactions(ctx *gin.Context) {
	data := service.GetTransactions()
	ctx.JSON(http.StatusOK, data)
}