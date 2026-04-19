package service

import (
	"api/internal/auth"
	"api/repository"
	"fmt"
)


func Signup(name, email, password string) (string, error){

	// Hash the password
	hashedPassword, err := auth.HashPassword(password)
	
	if err != nil {
		return "", err;
	}

	// save the user
	id, err := repository.CreateUser(name, email, hashedPassword);

	if err != nil {
		return "", err;
	}

	// Generate token
	token, err := auth.GenerateToken(id);

	if err != nil {
		return "", err;
	}
	return token, nil;
}

func Login(email, password string) (string, error) {
	// get the user
	userId, hashedPassword, err := repository.GetUserByEmail(email);
	if err != nil {
		return "", fmt.Errorf("Invalid Credentials");
	}

	// Compare the passwords
	err = auth.CheckPassword(password, hashedPassword);
	if err != nil {
		return "", fmt.Errorf("Invalid Credentials");
	}

	// Generate a token
	token, err := auth.GenerateToken(userId);

	if err != nil {
		return "", err;
	}

	return token, nil;
}

