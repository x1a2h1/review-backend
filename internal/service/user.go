package service

import (
	"errors"
	"review/internal/dto/req"
	"review/internal/models"
	"review/internal/pkg/database"
	"time"

	"gorm.io/gorm"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}
func (s *UserService) FindOneByAccount(account string) *models.User {
	db := database.GetMysql()
	user := &models.User{}
	if err := db.Where("account = ?", account).First(user).Error; err != nil {
		return nil
	}
	return user
}

func (s *UserService) Create(info *req.UserUpsertReq) error {
	existUser := s.FindOneByAccount(info.Account)
	if existUser != nil {
		return errors.New("账号已存在")
	}
	user := s.toModel(info)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Status = models.UserStatusNormal

	user.Password = NewAuthService().PasswordHash("123456")
	return database.GetMysql().Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("Root").Save(user).Error; err != nil {
			return err
		}
		// 将密码发送至邮箱
		// if info.ShouldSendPassword {
		// 	s.sendUserAddedEmailAsync(user.Email, generatedPasswd)
		// }
		return nil
	})
}
func (s *UserService) toModel(info *req.UserUpsertReq) *models.User {
	return &models.User{
		Account:   info.Account,
		Nickname:  info.Nickname,
		Avatar:    info.Avatar,
		Email:     info.Email,
		Cellphone: info.Cellphone,
	}
}

func (s *UserService) FindOneById(id uint, preloads ...string) *models.User {
	user := &models.User{}
	db := database.GetMysql()
	if len(preloads) > 0 {
		for _, preload := range preloads {
			db = db.Preload(preload)
		}
	}
	if err := db.First(user, id).Error; err != nil {
		return nil
	}
	return user
}
