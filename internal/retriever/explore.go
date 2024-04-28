package retriever

import (
	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"
	"api-service/internal/types"
	"context"

	"gorm.io/gorm"
)

// 定义 ExploreRetriever 接口
type ExploreRetriever interface {
	Query(c context.Context, queryRequest types.ExploreQueryReqest, limit int, after int) (*[]types.Exploredata, int, bool, error)
}

// 定义 ExploreRetriever 结构体
type exploreRetriever struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewExploreRetriever(db *gorm.DB, cache *cache.Cache) ExploreRetriever {
	return &exploreRetriever{
		db:    db,
		cache: cache,
	}
}
func (exp *exploreRetriever) Create(c context.Context, tables *[]model.Campaign) error {

	err := exp.db.WithContext(c).Create(tables).Error
	return err

}

func (exp *exploreRetriever) Query(c context.Context, queryRequest types.ExploreQueryReqest, limit int, after int) (*[]types.Exploredata, int, bool, error) {
	var Campaigns []model.Campaign
	var Explores []types.Exploredata
	var campaign model.Campaign
	var Space model.Space
	var count int64
	var HasNextPage bool

	//var campaign model.Campaign
	//var conditions []string

	HasNextPage = true

	deSession := exp.db.Session(&gorm.Session{})
	deSession = deSession.Model(campaign)

	if queryRequest.Statuses[0] != "all" {
		deSession = deSession.Model(campaign).Where("(status = ?", queryRequest.Statuses[0])
		for i, status := range queryRequest.Statuses {
			if i == len(queryRequest.Statuses)-1 {
				deSession.Or("status = ?)", status)
				continue
			}
			deSession.Or("status = ?", status)
		}
	}
	if queryRequest.Chains[0] != "all" {
		deSession = deSession.Model(campaign).Where("(chain = ?", queryRequest.Chains[0])
		for i, Chain := range queryRequest.Chains {
			if i == len(queryRequest.Chains)-1 {
				deSession.Or("chain = ?)", Chain)
				continue
			}
			deSession.Or("chain = ?", Chain)
		}
	}
	if queryRequest.RewardTypes[0] != "all" {
		deSession = deSession.Model(campaign).Where("(rewardTypes = ?", queryRequest.RewardTypes[0])
		for i, RewardType := range queryRequest.RewardTypes {
			if i == len(queryRequest.RewardTypes)-1 {
				deSession.Or("rewardTypes = ?)", RewardType)
				continue
			}
			deSession.Or("rewardTypes = ?", RewardType)
		}
	}
	if queryRequest.CredSources[0] != "all" {
		deSession = deSession.Model(campaign).Where("(credSources = ?", queryRequest.CredSources[0])
		for i, credSource := range queryRequest.CredSources {
			if i == len(queryRequest.CredSources)-1 {
				deSession.Or("credSources = ?)", credSource)
				continue
			}
			deSession.Or("credSources = ?", credSource)
		}
	}
	if queryRequest.SearchString != "" {
		deSession = deSession.Where("(name like ? OR name like ? OR name like ?)", queryRequest.SearchString+"%", "%"+queryRequest.SearchString, "%"+queryRequest.SearchString+"%")
	}

	//deSession.Model(campaign).Where("alias = ? AND status = ? AND chain =? AND gasType AND rewardTypes = ? AND credSources = ?  AND (name like ? OR name like ? OR name like ?)", queryRequest.Alias, queryRequest.Statuses, queryRequest.Chains, queryRequest.RewardTypes, queryRequest.CredSources, queryRequest.SearchString+"%", "%"+queryRequest.SearchString, "%"+queryRequest.SearchString+"%")

	// Find the intersection of ids from CredSources, RewardTypes, and Chains
	if queryRequest.ListType != "" {
		deSession = deSession.Order(queryRequest.ListType + " desc")
	}
	deSession.Count(&count)
	if int(count)-after-limit < 0 {
		limit = int(count) - after
		HasNextPage = false
	}
	db := deSession.Offset(after).Limit(limit)
	if err := db.Find(&Campaigns).Error; err != nil {
		// 处理错误
	}

	after = after + limit
	// // 使用 map 来存储别名，自动去除重复
	// aliasMap := make(map[string]struct{})
	// for _, camp := range Campaigns {
	// 	aliasMap[camp.Alias] = struct{}{}
	// }

	// // 将 map 的键转换为字符串切片
	// var uniqueAliases []string
	// for alias := range aliasMap {
	// 	uniqueAliases = append(uniqueAliases, alias)
	// }
	for _, camp := range Campaigns {
		deSessionSpace := exp.db.Session(&gorm.Session{})
		var spac model.Space
		deSessionSpace = deSessionSpace.Model(Space).Where("id=?", camp.SpaceID)
		deSessionSpace.Find(&spac)
		exp := types.Exploredata{
			Campaign: camp,
			Space:    spac,
		}
		Explores = append(Explores, exp)
	}

	return &Explores, after, HasNextPage, nil
}
