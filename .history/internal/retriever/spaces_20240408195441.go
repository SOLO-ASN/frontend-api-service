package retriever

import (
	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"
	"context"

	"gorm.io/gorm"
)

type SpacesRetriever interface {
	//Create (c context.Context, space *model.Space) error
	//Query (c context.Context) ([]*model.Space, error)
}

type spacesRetriever struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewSpacesRetriever(db *gorm.DB, cache *cache.Cache) SpacesRetriever {
	return &spacesRetriever{
		db:    db,
		cache: cache,
	}
}

func (s *spacesRetriever) Create(c interface{}, space *interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (s *spacesRetriever) Query(ctx context.Context, filter []string) ([]*interface{}, error) {
	var spaces []model.Space
	for _, filterValue := range filter {
		s.db = s.db.Where(func(db *gorm.DB) *gorm.DB {
			return db.Where("column = ?", filterValue)
		})
	}
	if err := s.db.Find(&spaces).Error; err != nil {
		// 处理错误
	}

}
