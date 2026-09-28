package services

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"slices"
	"strings"

	"github.com/google/uuid"

	"github.com/najayi/learning-go-shop/internal/interfaces"
)

type UploadService struct {
	provider interfaces.UploadProvider
}

func NewUploadService(provider interfaces.UploadProvider) *UploadService {
	return &UploadService{provider: provider}
}

func (s *UploadService) UploadProductImage(productID uint, file *multipart.FileHeader) (string, error) {

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !isValidImageExt(ext) {
		return "", fmt.Errorf("invalid file type: %s", ext)
	}

	newFileName := uuid.New().String()

	path := fmt.Sprintf("products/%d/%s%s", productID, newFileName, ext)

	return s.provider.UploadFile(file, path)
}

func isValidImageExt(ext string) bool {
	validExts := []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}
	return slices.Contains(validExts, ext)
}
