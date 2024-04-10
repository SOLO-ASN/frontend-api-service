package retriever

import (
	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"
	"api-service/internal/types"
	"context"
	"fmt"

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

/*
filter:筛选条件
count: 每一次请求要求的数量
*/
func (s *spacesRetriever) Query(ctx context.Context, pageSize int, count int, filters []types.Filter) (*[]model.Space, error) {
	var spaces []model.Space
	for _, filter := range filters {
		s.db = s.db.Where(func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("%s = ?", filter.Filter), filter.Name)
		})
	}
	// 调用 FindInBatches 方法
	if err := s.db.FindInBatches(&spaces, 20, func(tx *gorm.DB, batch int) error {
		return nil
	}); err != nil {

	}
	return &spaces, nil
}
