package service

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"review/internal/dto/req"
	"review/internal/models"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Login(req *req.LoginReq) (*models.User, error) {
	user := NewUserService().FindOneByAccount(req.Account)
	if user == nil {
		return nil, errors.New("账号不存在")
	}
	if user.Status != models.UserStatusNormal {
		return nil, errors.New("账号已被禁用")
	}
	sha1Passwd := s.PasswordHash(req.Password)
	if user.Password != sha1Passwd {
		return nil, errors.New("账号或密码有误")
	}
	return user, nil
}

func (s *AuthService) PasswordHash(pwd string) string {
	passwdBytes := sha1.Sum([]byte(pwd))
	return hex.EncodeToString(passwdBytes[:])
}
