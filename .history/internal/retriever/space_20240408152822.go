package retriever

import (
	"context"
	"errors"

	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"

	"gorm.io/gorm"
)

type SpaceRetriever interface {
	//Create(c context.Context, space *model.Space) error
	//Query(c context.Context) ([]*model.Space, error)
}

type spaceRetriever struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewSpaceRetriever(db *gorm.DB, cache *cache.Cache) SpaceRetriever {
	return &spaceRetriever{
		db:    db,
		cache: cache,
	}
}

func (s spaceRetriever) Create(c context.Context, space *interface{}) error {

	err := s.db.WithContext(c).Create(space).Error
	return err

}

func (s spaceRetriever) Query(c context.Context, alias string) ([]*interface{}, error) {
	var space model.Space
	if err := s.db.WithContext(ctx).First(&space, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 未找到记录
			return nil, err
		}
		// 发生了其他错误
		return nil, err
	}
	return &space, nil
}
