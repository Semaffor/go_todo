package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	todo_demo "github.com/semaffor/go-todo-app"
	"github.com/semaffor/go-todo-app/pkg/repostiroty"
	"github.com/semaffor/go-todo-app/pkg/util"
	"os"
	"time"
)

var (
	signKey = os.Getenv("SIGN_KEY")
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}

type AuthService struct {
	repo repostiroty.Authorization
}

func NewAuthService(repo repostiroty.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user todo_demo.User) (int, error) {
	user.Password = encryptPassword(user.Username, user.Password)
	return a.repo.CreateUser(user)
}

func (a *AuthService) GenerateToken(login, password string) (string, error) {
	user, err := a.repo.GetUser(login, encryptPassword(login, password))
	if err != nil {
		return "", err
	}
	return generateJwtAccessToken(user)
}

func (a *AuthService) ParseJwt(inputToken string) (int, error) {
	token, err := jwt.ParseWithClaims(inputToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		//Check on sign presence, except method HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		// key-sign
		return []byte(signKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims aren't of user type *tokenClaims")
	}

	return claims.UserId, nil
}

func encryptPassword(username, pass string) string {
	hash := sha256.New()
	hash.Write([]byte(pass))
	return fmt.Sprintf("%x", hash.Sum([]byte(username)))
}

func generateJwtAccessToken(user todo_demo.User) (string, error) {
	timeHours := util.ConvertToInt(os.Getenv("TOKEN_VALID_TIME_HOURS"))
	tokenValidTime := time.Duration(timeHours) * time.Hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenValidTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId:   user.Id,
		Username: user.Username,
	})
	return token.SignedString([]byte(signKey))
}
