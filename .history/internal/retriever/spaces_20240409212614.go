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
	Query(ctx context.Context, filters []types.Filter, limit int, after int) (*[]model.Space, int, bool, error)
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
limit: 每一次请求要求的数量
pageNumber:第几次请求
*/
func (s *spacesRetriever) Query(ctx context.Context, filters []string, limit int, after int) (*[]model.Space, int, bool, error) {
	var spaces []model.Space
	var count int64
	var HasNextPage bool
	HasNextPage = true

	for _, filter := range filters {
		s.db = s.db.Where(func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("%s = ?", filter), true)
		})
	}
	s.db.Count(&count)
	if int(count)-after-limit < 0 {
		limit = int(count) - after
		HasNextPage = false
	} else {
		after += limit
	}
	db := s.db.Offset(after).Limit(limit)
	after = after + limit
	// 执行查询并获取结果
	if err := db.Find(&spaces).Error; err != nil {
		// 处理错误
	}
	// 调用 FindInBatches 方法
	// if err := s.db.FindInBatches(&spaces, count, func(tx *gorm.DB, batch int) error {
	// 	return nil
	// }); err != nil {

	// }
	return &spaces, after, HasNextPage, nil
}
