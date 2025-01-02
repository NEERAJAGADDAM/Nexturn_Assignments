package controller

import (
	"blog-management/model"
	"blog-management/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *service.AuthService
}

// NewAuthController creates a new AuthController instance
func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

// Register handles user registration
func (controller *AuthController) Register(c *gin.Context) {
	var user model.User
	// Bind the JSON body to the user object
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the Register method of AuthService to create a user
	err := controller.AuthService.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Login handles user login and token generation
func (controller *AuthController) Login(c *gin.Context) {
	var user model.User
	// Bind the JSON body to the user object
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the Login method of AuthService to authenticate the user and generate a token
	token, err := controller.AuthService.Login(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Respond with the generated token
	c.JSON(http.StatusOK, gin.H{"token": token})
}
