package models

import "time"

type Base struct {
	ID        uint      `json:"id" gorm:"primarykey;autoIncrement;comment:主键ID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;comment:更新时间"`
}
