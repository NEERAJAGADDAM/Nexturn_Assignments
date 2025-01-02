package service

import (
	"blog-management/model"
	"blog-management/repository"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

var jwtKey = []byte("your_secret_key") // Replace with a strong secret key

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		UserRepo: userRepo,
	}
}

func (service *AuthService) Register(user *model.User) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return service.UserRepo.CreateUser(user)
}

func (service *AuthService) Login(user *model.User) (string, error) {
	dbUser, err := service.UserRepo.GetUserByUsername(user.Username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// Generate JWT token
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   user.Username,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
