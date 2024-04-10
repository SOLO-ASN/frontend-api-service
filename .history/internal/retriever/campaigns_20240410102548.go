package retriever

import (
	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"
	"api-service/internal/types"
	"context"

	"gorm.io/gorm"
)

type CampaignsRetriever interface {
	//Create (c context.Context, space *model.Campaign) error
	Query(c context.Context, queryRequest types.CampaignsQueryReqest, limit int, after int) (*[]model.Campaign, int, bool, error)
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
func (cams *campaignsRetriever) Query(c context.Context, queryRequest types.CampaignsQueryReqest, limit int, after int) (*[]model.Campaign, int, bool, error) {
	var campaigns []model.Campaign
	var count int64
	var HasNextPage bool
	HasNextPage = true
	cams.db.Count(&count)

	// for _, filter := range filters {
	// 	s.db = s.db.Model(space).Where(fmt.Sprintf("%s = ?", filter), true)
	// 	//s.db = s.db.Model(space).Where(fmt.Sprintf("%s = ?", filter), 1)
	// 	// (func(db *gorm.DB) *gorm.DB {
	// 	// 	return db.Where(fmt.Sprintf("%s = ?", filter), 1)
	// 	// })
	// }
	cams.db = cams.db.Where(func(db *gorm.DB) *gorm.DB {
		return db.Where("alias = ?", queryRequest.Alias)
	})
	for _, CredSource := range queryRequest.CredSources {
		cams.db = cams.db.Where(func(db *gorm.DB) *gorm.DB {
			return db.Where("credSources = ?", CredSource)
		})
	}
	for _, RewardTypes := range queryRequest.CredSources {
		cams.db = cams.db.Where(func(db *gorm.DB) *gorm.DB {
			return db.Where("rewardTypes = ?", RewardTypes)
		})
	}
	for _, Chain := range queryRequest.Chains {
		cams.db = cams.db.Where(func(db *gorm.DB) *gorm.DB {
			return db.Where("chain = ?", Chain)
		})
	}
	if int(count)-after-limit < 0 {
		limit = int(count) - after
		HasNextPage = false
	} else {
		after += limit
	}
	db := cams.db.WithContext(c).Offset(after).Limit(limit)
	after = after + limit
	// 执行查询并获取结果
	if err := db.Find(&campaigns).Error; err != nil {
		// 处理错误
	}
	return &campaigns, after, HasNextPage, nil
}
