package service

type UploadService struct{}

func New() *UploadService {
	return &UploadService{}
}
func (s *UploadService) Upload() {}

// Config 上传配置
// 应用对象存储平台
// 可选对象存储服务商
func (s *UploadService) Config() {}
