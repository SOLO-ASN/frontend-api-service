package retriever

import (
	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type CampaignRetriever interface {
	//Create (c context.Context, space *model.Campaign) error
	//Query (c context.Context) ([]*model.Campaign, error)
}

type campaignRetriever struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewCampaignRetriever(db *gorm.DB, cache *cache.Cache) CampaignRetriever {
	return &campaignRetriever{
		db:    db,
		cache: cache,
	}
}
func (cam *campaignRetriever) Create(c context.Context, space *interface{}) error {

	err := cam.db.WithContext(c).Create(space).Error
	return err

}
func (cam *campaignRetriever) Query(c context.Context, alias string) (*model.Campaign, error) {
	var campaign model.Campaign
	if err := cam.db.WithContext(c).First(&campaign, "alias = ?", alias).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 未找到记录
			return nil, err
		}
		// 发生了其他错误
		return nil, err
	}
	return &campaign, nil
}
