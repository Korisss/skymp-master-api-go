package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Korisss/skymp-master-api-go/internal/domain"
	"github.com/Korisss/skymp-master-api-go/internal/repository"
	"github.com/dgrijalva/jwt-go"
)

var salt string = os.Getenv("PASSWORD_SALT")
var jwtSecret string = os.Getenv("JWT_SECRET")

const tokenTTL = 8760 * time.Hour

type tokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user domain.User) (int64, error) {
	user.Password = generatePasswordHash(user.Password, salt)
	return s.repo.CreateUser(user)
}

func (s *AuthService) ParseToken(accessToken string) (int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return -1, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return -1, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *AuthService) GenerateToken(email, password string) (int64, string, error) {
	user, err := s.repo.GetUser(email, generatePasswordHash(password, salt))
	if err != nil {
		return -1, "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	signedString, err := token.SignedString([]byte(jwtSecret))

	return user.Id, signedString, err
}

func (s *AuthService) GetUserName(id int64) (string, error) {
	name, err := s.repo.GetUserName(id)
	if err != nil {
		return "", err
	}

	return name, err
}

func generatePasswordHash(password string, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
