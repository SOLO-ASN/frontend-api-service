package retriever

import (
	"context"
	"errors"
	"fmt"

	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"

	"gorm.io/gorm"
)

type SpaceRetriever interface {
	Create(c context.Context, table *model.Space) error
	Query(c context.Context, alias string) (*model.Space, error)
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

func (s spaceRetriever) Create(c context.Context, table *model.Space) error {

	err := s.db.WithContext(c).Create(table).Error
	return err

}

/*

 */
func (s spaceRetriever) Query(c context.Context, alias string) (*model.Space, error) {
	var space model.Space
	var token model.Token
	if err := s.db.First(&space, "alias = ?", alias).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 未找到记录
			return nil, err
		}
		// 发生了其他错误
		return nil, err
	}
	fmt.Println(space.TokenID)
	if err := s.db.First(&token, "id = ?", space.TokenID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 未找到记录
			return nil, err
		}
		// 发生了其他错误
		return nil, err
	}
	// res := s.db.First(&space, "alias = ?", alias)
	// if res.Error != nil {

	// } else {
	// 	fmt.Printf("Found user: %+v\n", space)
	// }

	space.Token = token

	return &space, nil
}
