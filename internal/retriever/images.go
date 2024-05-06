package retriever

import (
	"api-service/internal/dbEntity/cache"
	"api-service/internal/types"
	"context"

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
func (s imagesRetriever) Upload(c1 context.Context, request types.ImageUploadRequest, uploadPath string) (types.ImageUploadResponse, error) {

	// var c *gin.Context
	//r1()
	// file, err := c.FormFile("file")
	// if err != nil {

	// 	c.JSON(http.StatusOK, fmt.Sprintf("'%s' uploaded", file.Filename))
	// 	return types.ImageUploadResponse{}, nil
	// }

	// filepath := path.Join("./"+"images/", file.Filename)
	// err = c.SaveUploadedFile(file, filepath)
	// if err != nil {

	// 	c.JSON(http.StatusOK, err.Error())
	// 	return types.ImageUploadResponse{}, nil
	// }
	// c.JSON(http.StatusOK, gin.H{"uploading": "done", "message": "success", "url": "http://" + c.Request.Host + "images/" + file.Filename})

	return types.ImageUploadResponse{Url: "http://yourdomain.com/"}, nil
}
