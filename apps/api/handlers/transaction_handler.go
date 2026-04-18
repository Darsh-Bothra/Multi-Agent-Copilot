package handlers

import (
	"api/internal/validator"
	"api/pkg/errors"
	"api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetTransactions(ctx *gin.Context) {
	data, err := service.GetTransactions()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}


func CreateTransaction(ctx *gin.Context) {
	// Receive Input -> Bind the json with received data -> Validate the data -> Service -> Error/Accepted Response

	// Input format
	var input struct {
		Amount float64 `json:"amount"`
		Merchant string `json:"merchant"`
	}

	// Step 1: for parsing the struct body and Bind the json
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, 
			errors.NewBadRequest("invalid input"))
			
		return;
	}

	// Step 2: Validate
	if err := validator.Validate.Struct(input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {"error": err.Error()});
	}

	// Step 3 : Service layer
	err := service.CreateTransaction(input.Amount, input.Merchant);
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {"error": err.Error()});
		return;
	} 
	
	// 	No error means the transaction created
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Transaction Created Successfully"})
}