package comm

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"review/internal/dto/res"
	"review/internal/pkg/storage"

	"github.com/gin-gonic/gin"
)

// Upload 处理文件上传
func upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		res.Fail(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		res.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	defer src.Close()

	// 计算文件MD5
	fileHash, err := calculateFileMD5(src)
	if err != nil {
		res.Fail(ctx, http.StatusInternalServerError, "计算文件MD5失败: "+err.Error())
		return
	}

	// 重置文件读取位置
	src.Seek(0, 0)

	// 保存临时文件
	tempDir := "./temp"
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		res.Fail(ctx, http.StatusInternalServerError, "创建临时目录失败: "+err.Error())
		return
	}

	tempFile := filepath.Join(tempDir, file.Filename)
	dst, err := os.Create(tempFile)
	if err != nil {
		res.Fail(ctx, http.StatusInternalServerError, "创建临时文件失败: "+err.Error())
		return
	}
	defer dst.Close()

	// 复制文件内容
	if _, err := io.Copy(dst, src); err != nil {
		os.Remove(tempFile)
		res.Fail(ctx, http.StatusInternalServerError, "保存文件失败: "+err.Error())
		return
	}

	fmt.Printf("接收文件: %s, MD5: %s\n", file.Filename, fileHash)

	// 上传到对象存储
	c := context.Background()
	storageInstance, err := storage.NewStorage()
	if err != nil {
		os.Remove(tempFile)
		res.Fail(ctx, http.StatusInternalServerError, "存储初始化失败: "+err.Error())
		return
	}
	
	result, err := storageInstance.Upload(c, &storage.Input{
		FilePath: tempFile,
		FileName: file.Filename,
		Size:     file.Size,
		Hash:     fileHash,
	})
	if err != nil {
		os.Remove(tempFile)
		res.Fail(ctx, http.StatusInternalServerError, "上传失败: "+err.Error())
		return
	}

	// 清理临时文件
	os.Remove(tempFile)

	res.Success(ctx, result)
}

// calculateFileMD5 计算文件的MD5值
func calculateFileMD5(file io.Reader) (string, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
