package service

import (
	"crypto/sha256"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	todo_demo "github.com/semaffor/go-todo-app"
	"github.com/semaffor/go-todo-app/pkg/repostiroty"
	"github.com/semaffor/go-todo-app/pkg/util"
	"os"
	"time"
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
	signKey := os.Getenv("SIGN_KEY")
	return token.SignedString([]byte(signKey))
}
