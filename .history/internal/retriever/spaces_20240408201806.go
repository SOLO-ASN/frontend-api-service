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

/*
filter:筛选条件
pageSize:第几次请求
count: 每一次请求要求的数量
*/
func (s *spacesRetriever) Query(ctx context.Context, pageSize int, count int, filter []string) (*[]model.Space, error) {
	var spaces []model.Space
	// 计算偏移量
	offset := (pageNumber - 1) * count

	// 设置查询结果的数量限制和偏移量
	db := s.db.WithContext(ctx).Limit(pageSize).Offset(offset)
	for _, filterValue := range filter {
		s.db = s.db.Where(func(db *gorm.DB) *gorm.DB {
			return db.Where("column = ?", filterValue)
		})
	}
	if err := s.db.Find(&spaces).Error; err != nil {
		// 处理错误
	}
	return &spaces, nil
}
