package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/43nvy/todo-app"
	"github.com/43nvy/todo-app/internal/repository"
)

const salt = "dashdi123dja*Ds21y3480_DS"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePassHash(user.Password)

	return s.repo.CreateUser(user)
}

func (a *AuthService) generatePassHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
