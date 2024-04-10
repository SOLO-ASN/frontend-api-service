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
	s.db = s.db.Session(&gorm.Session{})
	//s.db = s.db.Model(space).Where("name like ?", request.SearchString+"%").Or("name like ?", "%"+request.SearchString).Or("name like ?", "%"+request.SearchString+"%")
	s.db = s.db.Model(space).Where("(name like ? OR name like ? OR name like ?) AND isVerified = ?", request.SearchString+"%", "%"+request.SearchString, "%"+request.SearchString+"%", request.VerifiedOnly)
	fmt.Println(request.VerifiedOnly)
	//s.db = s.db.Model(space).Where("isVerified = ?", request.VerifiedOnly)
	// for _, filter := range filters {
	// 	s.db = s.db.Model(space).Where(fmt.Sprintf("%s = ?", filter), true)
	// 	//s.db = s.db.Model(space).Where(fmt.Sprintf("%s = ?", filter), 1)
	// 	// (func(db *gorm.DB) *gorm.DB {
	// 	// 	return db.Where(fmt.Sprintf("%s = ?", filter), 1)
	// 	// })
	// }

	s.db.Count(&count)
	fmt.Println(count)
	if int(count)-after-limit <= 0 {
		limit = int(count) - after
		HasNextPage = false
	}
	fmt.Println(limit)
	db := s.db.Offset(after).Limit(limit)
	// 执行查询并获取结果
	if err := db.Order(request.SpaceListType).Find(&spaces).Error; err != nil {
		// 处理错误
	}
	fmt.Println(spaces)
	after = after + limit
	// 调用 FindInBatches 方法
	// if err := s.db.FindInBatches(&spaces, count, func(tx *gorm.DB, batch int) error {
	// 	return nil
	// }); err != nil {

	// }
	return &spaces, after, HasNextPage, nil
}
