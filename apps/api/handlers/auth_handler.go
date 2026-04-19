package handlers

import (
	"api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogInRequest struct {
	Email string `json:"email"`
	Password  string `json:"password"`
}

func Signup(ctx *gin.Context) {
	var signUpReq SignUpRequest

	// parse the request
	if err := ctx.ShouldBindJSON(&signUpReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Get the token
	token, err := service.Signup(signUpReq.Name, signUpReq.Email, signUpReq.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "signup failed"})
		return
	}

	// Return the response
	ctx.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}


func Login(ctx *gin.Context) {
	var logInReq LogInRequest;

	// parse the request
	if err := ctx.ShouldBindJSON(&logInReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// generate the token
	token, err := service.Login(logInReq.Email, logInReq.Password);
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "signup failed"})
		return
	}

	// Return the response
	ctx.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}