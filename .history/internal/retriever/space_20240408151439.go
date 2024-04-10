package retriever

import (
	"context"

	"api-service/internal/dbEntity/cache"

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
	//TODO implement me
	panic("implement me")
}

func (s spaceRetriever) Query(c context.Context) ([]*interface{}, error) {
	//TODO implement me
	panic("implement me")
}
