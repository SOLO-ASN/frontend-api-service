package retriever

import (
	"context"
	"errors"
	"fmt"

	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"
	"api-service/internal/types"

	"gorm.io/gorm"
)

type SpaceRetriever interface {
	Create(c context.Context, request *types.SpaceCreateRequest) (string, error)
	Query(c context.Context, request types.SpaceQueryRequest) (*model.Space, error)
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

func (s spaceRetriever) Create(c context.Context, request *types.SpaceCreateRequest) (string, error) {
	var space model.Space
	var user model.Jpa_web_authn_user
	space.Name = request.Name
	space.Alias = request.Alias
	space.Categories = request.Categories
	space.Links = request.Links
	space.Info = request.Info
	space.Thumbnail = request.Thumbnail
	deSession := s.db.Session(&gorm.Session{})
	deSession = deSession.Model(user).Where("username = ?", request.Username)
	deSession.First(&user)
	space.Owner = string(user.Id)
	deSession = s.db.Session(&gorm.Session{})
	result := deSession.Create(&space) // 通过数据的指针来创建
	if result.Error != nil {

		return "false", nil
	}
	return "success", nil

}

/*

 */
func (s spaceRetriever) Query(c context.Context, request types.SpaceQueryRequest) (*model.Space, error) {
	var space model.Space
	var token model.Token
	var SpaceFollower model.SpaceFollower
	var user model.Jpa_web_authn_user

	deSession := s.db.Session(&gorm.Session{})
	if err := deSession.First(&space, "id = ?", request.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 未找到记录
			return nil, err
		}
		// 发生了其他错误
		return nil, err
	}

	deSession.First(&token, "id = ?", space.TokenID)
	// if err := deSession.First(&token, "id = ?", space.TokenID).Error; err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		// 未找到记录
	// 		return nil, err
	// 	}
	// 	// 发生了其他错误
	// 	return nil, err
	// }
	// res := s.db.First(&space, "alias = ?", alias)
	// if res.Error != nil {

	// } else {
	// 	fmt.Printf("Found user: %+v\n", space)
	// }

	space.Token = token
	deSession = s.db.Session(&gorm.Session{})
	deSession = deSession.Model(user).Where("username = ?", request.Username)
	deSession.First(&user)
	deSession = s.db.Session(&gorm.Session{})
	deSession = deSession.Model(SpaceFollower).Where("participantId = ? AND spaceId = ? ", user.Id, request.Id)
	deSession.First(&SpaceFollower)
	space.IsFollowing = SpaceFollower.IsFollowing
	fmt.Println(space)
	return &space, nil
}
