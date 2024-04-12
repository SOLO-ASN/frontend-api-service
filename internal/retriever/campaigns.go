package retriever

import (
	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"
	"api-service/internal/types"
	"context"
	"fmt"

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

func NewCampaignsRetriever(db *gorm.DB, cache *cache.Cache) CampaignsRetriever {
	return &campaignsRetriever{
		db:    db,
		cache: cache,
	}
}
func (cams *campaignsRetriever) Create(c context.Context, tables *[]model.Campaign) error {

	err := cams.db.WithContext(c).Create(tables).Error
	return err

}

type UseCredIds struct {
	Ids []string
}
type ids struct {
	Id string
}

func (cams *campaignsRetriever) Query(c context.Context, queryRequest types.CampaignsQueryReqest, limit int, after int) (*[]model.Campaign, int, bool, error) {
	var campaigns []model.Campaign
	var campaign model.Campaign
	var count int64
	var HasNextPage bool
	var idsFromCredSource []ids
	var idsFromRewardTypes []ids
	var idsFromChains []ids
	var filter string
	//var campaign model.Campaign
	//var conditions []string

	HasNextPage = true

	deSession := cams.db.Session(&gorm.Session{})

	//campaignRequirementTwitterEngagement
	//deSession = deSession.Where("alias = ?", queryRequest.Alias)
	if queryRequest.CredSources[0] != "all" {
		for _, CredSource := range queryRequest.CredSources {
			var ids []ids
			deSession.Table("campaignRequirement"+CredSource).Select("id").Where("alias = ?", queryRequest.Alias).Scan(&ids)
			idsFromCredSource = append(idsFromCredSource, ids...)

			//cams.db.Where("id IN (?)", cams.db.Table("campaignRequirement"+CredSource).Select("id")).Find(&campaigns)
		}
	}

	fmt.Println(idsFromCredSource)
	if queryRequest.RewardTypes[0] != "all" {
		for _, RewardTypes := range queryRequest.RewardTypes {
			var ids []ids
			deSession.Table("campaignRewardType"+RewardTypes).Select("id").Where("alias = ?", queryRequest.Alias).Scan(&ids)
			idsFromRewardTypes = append(idsFromRewardTypes, ids...)

			//cams.db.Where("id IN (?)", cams.db.Table("campaignRequirement"+CredSource).Select("id")).Find(&campaigns)
		}
	}
	fmt.Println(idsFromRewardTypes)
	if queryRequest.Chains[0] != "all" {
		for _, Chains := range queryRequest.Chains {
			var ids []ids
			deSession.Table("campaignRewardDistributedOn"+Chains).Select("id").Where("alias = ?", queryRequest.Alias).Scan(&ids)
			idsFromChains = append(idsFromChains, ids...)

			//cams.db.Where("id IN (?)", cams.db.Table("campaignRequirement"+CredSource).Select("id")).Find(&campaigns)
		}
	}
	fmt.Println(idsFromChains)
	// Find the intersection of ids from CredSources, RewardTypes, and Chains
	var intersection []string
	intersectionMap := make(map[string]bool)
	if queryRequest.CredSources[0] != "all" {

		for _, credID := range idsFromCredSource {
			if intersectionMap[credID.Id] {
				continue
			}

			foundInRewardTypes := false
			foundInChains := false
			if queryRequest.RewardTypes[0] == "all" {
				foundInRewardTypes = true
			}
			if queryRequest.Chains[0] == "all" {
				foundInChains = true
			}
			for _, rewardID := range idsFromRewardTypes {
				if credID.Id == rewardID.Id {
					foundInRewardTypes = true
					break
				}
			}

			if !foundInRewardTypes {
				continue
			}

			for _, chainID := range idsFromChains {
				if credID.Id == chainID.Id {
					foundInChains = true
					break
				}
			}

			if foundInChains {
				intersection = append(intersection, credID.Id)
				intersectionMap[credID.Id] = true
			}
		}
	} else if queryRequest.RewardTypes[0] != "all" {
		for _, rewardID := range idsFromRewardTypes {
			if intersectionMap[rewardID.Id] {
				continue
			}

			foundInChains := false

			if queryRequest.Chains[0] == "all" {
				foundInChains = true
			}
			for _, chainID := range idsFromChains {
				if rewardID.Id == chainID.Id {
					foundInChains = true
					break
				}
			}

			if foundInChains {
				intersection = append(intersection, rewardID.Id)
				intersectionMap[rewardID.Id] = true
			}
		}

	} else if queryRequest.Chains[0] == "all" {
		for _, chainID := range idsFromChains {
			intersection = append(intersection, chainID.Id)

		}

	} else {
		filter = "all"
	}

	fmt.Println(intersection)
	deSessionForCampaign := cams.db.Session(&gorm.Session{})
	if filter != "all" {

		deSessionForCampaign = deSessionForCampaign.Model(campaign).Where("id IN (?) AND (name like ? OR name like ? OR name like ?)", intersection, queryRequest.SearchString+"%", "%"+queryRequest.SearchString, "%"+queryRequest.SearchString+"%")

	} else {
		deSessionForCampaign = deSessionForCampaign.Model(campaign).Where("name like ? OR name like ? OR name like ?", queryRequest.SearchString+"%", "%"+queryRequest.SearchString, "%"+queryRequest.SearchString+"%")
	}
	deSessionForCampaign = deSessionForCampaign.Order(queryRequest.ListType)
	deSessionForCampaign.Count(&count)
	fmt.Println(count)
	if int(count)-after-limit < 0 {
		limit = int(count) - after
		HasNextPage = false
	}
	db := deSessionForCampaign.Offset(after).Limit(limit)
	if err := db.Find(&campaigns).Error; err != nil {
		// 处理错误
	}

	after = after + limit
	return &campaigns, after, HasNextPage, nil
}
