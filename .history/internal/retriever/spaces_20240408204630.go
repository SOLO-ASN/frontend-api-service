package retriever

import (
	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"
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
pageSize:第几次请求
count: 每一次请求要求的数量
*/
func (s *spacesRetriever) Query(ctx context.Context, pageSize int, count int, filter []string) (*[]model.Space, error) {
	var spaces []model.Space
	for _, filterValue := range filter {
		s.db = s.db.Where(func(db *gorm.DB) *gorm.DB {
			return db.Where("column = ?", filterValue)
		})
	}
	// 调用 FindInBatches 方法
	if err := s.db.FindInBatches(&spaces, 20, func(tx *gorm.DB, batch int) error {
		// 在这里处理每个批次的记录
		for _, space := range spaces {
			// 例如，更新每个空间的某个字段
			space.Info = "processed"

			// 保存更改到数据库
			if err := tx.Save(&space).Error; err != nil {
				return err // 如果保存出错，返回错误并停止后续批次
			}
		}

		// 输出当前批次处理的记录数
		fmt.Printf("Batch %d processed %d records\n", batch, tx.RowsAffected)

		// 返回 nil 继续处理下一个批次
		return nil
	}); err != nil {

	}
	return &spaces, nil
}
