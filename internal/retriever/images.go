package retriever

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	"api-service/internal/dbEntity/cache"
	"api-service/internal/types"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ImagesRetriever interface {
	//Create(c context.Context) error
	Upload(c context.Context, request types.ImageUploadRequest, uploadPath string) (types.ImageUploadResponse, error)
}

type imagesRetriever struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewImagesRetriever(db *gorm.DB, cache *cache.Cache) ImagesRetriever {
	return &imagesRetriever{
		db:    db,
		cache: cache,
	}
}

// func (s imagesRetriever) Create(c context.Context, table *model.Images) error {

// 	err := s.db.WithContext(c).Create(table).Error
// 	return err

// }

/*

 */
func (s imagesRetriever) Upload(c context.Context, request types.ImageUploadRequest, uploadPath string) (types.ImageUploadResponse, error) {
	var c1 *gin.Context
	file, err := c1.FormFile("file")
	if err != nil {
		c1.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return types.ImageUploadResponse{}, nil
	}
	filename := filepath.Base(file.Filename)
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.Mkdir(uploadPath, os.ModePerm)
	}
	filepath := uploadPath + filename
	if err := c1.SaveUploadedFile(file, filepath); err != nil {
		c1.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return types.ImageUploadResponse{}, nil
	}
	c1.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"message":    "Image uploaded successfully",
		"image_path": filepath,
	})
	return types.ImageUploadResponse{Url: "http://yourdomain.com/" + filepath}, nil
}
