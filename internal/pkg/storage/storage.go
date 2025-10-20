package storage

import (
	"context"
	"review/internal/pkg/config"
)

type Storage struct{
	uploader Uploader
}
type Input struct {
	FilePath string
	FileName string
	Size     int64
	Hash     string
}

type Output struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Url        string `json:"url"`
	BucketHash string `json:"bucketHash"`
}

func NewStorage() (*Storage, error) {
	var uploader Uploader
	
	switch config.Cfg.Storage.Platform {
	case "kodo":
		uploader = NewKODOUploader(config.Cfg.Storage.KODO)
	default:
		uploader = NewKODOUploader(config.Cfg.Storage.KODO) // 保持向后兼容
	}
	
	return &Storage{uploader: uploader}, nil
}

func (s *Storage) Upload(ctx context.Context, input *Input) (*Output, error) {
	return s.uploader.Upload(ctx, input)
}


