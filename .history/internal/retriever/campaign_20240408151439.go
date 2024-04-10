package retriever

import (
	"api-service/internal/dbEntity/cache"

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

func (c *campaignRetriever) Create() error {
	//TODO implement me
	panic("implement me")
}

func (c *campaignRetriever) Query() {
	//TODO implement me
	panic("implement me")
}
