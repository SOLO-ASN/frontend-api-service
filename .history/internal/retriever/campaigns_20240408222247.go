package retriever

import (
	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"
	"context"

	"gorm.io/gorm"
)

type CampaignsRetriever interface {
	//Create (c context.Context, space *model.Campaign) error
	//Query (c context.Context) ([]*model.Campaign, error)
}

type campaignsRetriever struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewCampaignsRetriever(db *gorm.DB, cache *cache.Cache) CampaignRetriever {
	return &campaignRetriever{
		db:    db,
		cache: cache,
	}
}
func (cams *campaignsRetriever) Create(c context.Context, tables *[]model.Campaign) error {

	err := cams.db.WithContext(c).Create(tables).Error
	return err

}
func (cams *campaignsRetriever) Query(c context.Context, alias string, limit int, pageNumber int) (*[]model.Campaign, error) {
	var campaign []model.Campaign
	var count int64
	offset := (pageNumber - 1) * limit
	cams.db.Count(&count)
	cams.db = cams.db.Where(func(db *gorm.DB) *gorm.DB {
		return db.Where("alias = ?", alias)
	})
	if int(count)-offset-limit < 0 {
		limit = int(count) - offset
	} else {
		offset += limit
	}
	db := s.db.WithContext(ctx).Offset(offset).Limit(limit)
	// 执行查询并获取结果
	if err := db.Find(&spaces).Error; err != nil {
		// 处理错误
	}
	return &campaign, nil
}
