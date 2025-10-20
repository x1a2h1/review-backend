package service

import (
	"review/internal/dto/res"
	"review/internal/models"
	"review/internal/pkg/database"
)

type LogService struct{}

func NewLogService() *LogService {
	return &LogService{}
}

func (s *LogService) AddLoginLog(info *models.LoginLog) error {
	if err := database.GetMysql().Save(info).Error; err != nil {
		return err
	}
	return nil
}

func (s *LogService) LoginLogPageList(keyword string, page, pageSize int) (*res.PageableData[*models.LoginLog], error) {
	logs := make([]*models.LoginLog, 0)
	pageableData := &res.PageableData[*models.LoginLog]{}
	db := database.GetMysql().Model(&models.LoginLog{})
	if keyword != "" {
		db.Where("nickname like ?", "%"+keyword+"%")
	}
	var count int64
	err := db.Count(&count).Order("create_time desc").Offset(pageSize * (page - 1)).Limit(pageSize).Find(&logs).Error
	if err != nil {
		return nil, err
	}
	pageableData.List = logs
	pageableData.Total = count
	return pageableData, nil
}
