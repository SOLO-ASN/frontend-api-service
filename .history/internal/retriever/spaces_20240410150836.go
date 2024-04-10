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
	Query(ctx context.Context, request types.SpacesQueryRequest, limit int, after int) (*[]model.Space, int, bool, error)
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
func (s *spacesRetriever) Query(ctx context.Context, request types.SpacesQueryRequest, limit int, after int) (*[]model.Space, int, bool, error) {
	var spaces []model.Space
	var space model.Space
	var count int64
	var HasNextPage bool

	HasNextPage = true
	deSession := s.db.Session(&gorm.Session{})
	deSession = deSession.Model(space).Where("(name like ? OR name like ? OR name like ?) AND isVerified = ?", request.SearchString+"%", "%"+request.SearchString, "%"+request.SearchString+"%", request.VerifiedOnly)
	fmt.Println(request.SearchString, request.VerifiedOnly)
	deSession.Count(&count)
	fmt.Println(count)
	if int(count)-after-limit <= 0 {
		limit = int(count) - after
		HasNextPage = false
	}
	fmt.Println(limit)
	db := deSession.Offset(after).Limit(limit)
	// 执行查询并获取结果
	if err := db.Order(request.SpaceListType).Find(&spaces).Error; err != nil {
		// 处理错误
	}
	fmt.Println(spaces)
	after = after + limit
	return &spaces, after, HasNextPage, nil
}
