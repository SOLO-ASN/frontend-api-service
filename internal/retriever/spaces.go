package retriever

import (
	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"
	"api-service/internal/types"
	"context"

	"gorm.io/gorm"
)

type SpacesRetriever interface {
	//Create (c context.Context, space *model.Space) error
	Query(ctx context.Context, request types.SpacesQueryRequest, limit int, after int) (*[]model.Space, int, bool, error)
	Follow(ctx context.Context, request types.FollowRequest) (string, error)
	UnFollow(ctx context.Context, request types.FollowRequest) (string, error)
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
	var SpaceFollower model.SpaceFollower
	var SpaceFollowers []model.SpaceFollower
	var user model.Jpa_web_authn_user
	HasNextPage = true
	if request.Filter == "follow" {
		if &request.Username == nil {
			return &spaces, after, HasNextPage, nil
		}

		deSession := s.db.Session(&gorm.Session{})
		deSession = deSession.Model(user).Where("username = ?", request.Username)
		deSession.First(&user)
		deSession = s.db.Session(&gorm.Session{})
		deSession = deSession.Model(SpaceFollower).Where("participantId = ? ", user.Id)
		res := deSession.Find(&SpaceFollowers)
		if res.Error != nil {
			return &spaces, after, HasNextPage, nil
		}
		var spaceids []string
		for _, spacefollower := range SpaceFollowers {
			if spacefollower.IsFollowing == true {
				spaceids = append(spaceids, spacefollower.SpaceId)
			}
		}
		deSession = s.db.Session(&gorm.Session{})
		deSession = deSession.Model(spaces).Where("(id = ?", spaceids[0])
		for i, spaceId := range spaceids {
			if i == len(spaceids)-1 {
				deSession.Or("id = ?)", spaceId)
				continue
			}
			deSession.Or("id = ?", spaceId)
		}
		if request.VerifiedOnly != false {
			deSession = deSession.Where("(name like ? OR name like ? OR name like ?) AND isVerified = ?", request.SearchString+"%", "%"+request.SearchString, "%"+request.SearchString+"%", request.VerifiedOnly)
		} else {
			deSession = deSession.Where("(name like ? OR name like ? OR name like ?)", request.SearchString+"%", "%"+request.SearchString, "%"+request.SearchString+"%")
		}
		deSession.Count(&count)
		if int(count)-after-limit <= 0 {
			limit = int(count) - after
			HasNextPage = false
		}

		db := deSession.Offset(after).Limit(limit)
		if err := db.Order(request.SpaceListType + " desc").Find(&spaces).Error; err != nil {
			// 处理错误
		}
		after = after + limit
		for i, _ := range spaces {
			spaces[i].IsFollowing = true
		}
		return &spaces, after, HasNextPage, nil
	}

	deSession := s.db.Session(&gorm.Session{})
	deSession = deSession.Model(space)
	if request.SearchString != "" {

		if request.VerifiedOnly != false {
			deSession = deSession.Model(space).Where("(name like ? OR name like ? OR name like ?) AND isVerified = ?", request.SearchString+"%", "%"+request.SearchString, "%"+request.SearchString+"%", request.VerifiedOnly)
		} else {
			deSession = deSession.Model(space).Where("(name like ? OR name like ? OR name like ?)", request.SearchString+"%", "%"+request.SearchString, "%"+request.SearchString+"%")
		}
	} else {
		if request.VerifiedOnly != false {
			deSession = deSession.Model(space).Where(" isVerified = ?", request.VerifiedOnly)
		} else {

		}
	}
	deSession.Count(&count)
	if int(count)-after-limit <= 0 {
		limit = int(count) - after
		HasNextPage = false
	}

	db := deSession.Offset(after).Limit(limit)
	// 执行查询并获取结果
	if err := db.Order(request.SpaceListType + " desc").Find(&spaces).Error; err != nil {
		// 处理错误
	}

	after = after + limit

	if request.Username != "" {
		deSession = s.db.Session(&gorm.Session{})
		deSession = deSession.Model(user).Where("username = ?", request.Username)
		deSession.First(&user)

		for i, space := range spaces {
			var SpaceFollower1 model.SpaceFollower
			deSession1 := s.db.Session(&gorm.Session{})
			deSession1 = deSession1.Model(SpaceFollower1).Where("participantId = ? AND spaceId = ? ", user.Id, space.ID)
			res := deSession1.First(&SpaceFollower1)
			if res.Error != nil {
				spaces[i].IsFollowing = false
				continue
			}

			//fmt.Println(SpaceFollower)
			spaces[i].IsFollowing = SpaceFollower1.IsFollowing
		}
	}
	return &spaces, after, HasNextPage, nil
}
func (s *spacesRetriever) Follow(ctx context.Context, request types.FollowRequest) (string, error) {
	var SpaceFollower model.SpaceFollower
	var user model.Jpa_web_authn_user
	var space model.Space
	success := "Follow Success"
	deSession := s.db.Session(&gorm.Session{})
	deSession = deSession.Model(user).Where("username = ?", request.Username)
	res := deSession.First(&user)

	if res.Error != nil {
		return "false", nil
	}
	deSession = s.db.Session(&gorm.Session{})
	deSession = deSession.Model(SpaceFollower).Where("participantId = ? AND spaceId = ? ", user.Id, request.SpaceId)

	res1 := deSession.First(&SpaceFollower)

	if res1.Error != nil {
		Follower := model.SpaceFollower{SpaceId: request.SpaceId, ParticipantId: user.Id, IsFollowing: true}
		deSession = s.db.Session(&gorm.Session{})
		result := deSession.Create(&Follower) // 通过数据的指针来创建
		if result.Error != nil {

			return "false", nil
		}
		deSession = s.db.Session(&gorm.Session{})
		deSession = deSession.Model(space).Where("id=?", request.SpaceId)
		deSession = deSession.First(&space)

		space.Followers += 1
		deSession.Save(&space)
		return success, nil
	}
	SpaceFollower.IsFollowing = true
	deSession.Save(&SpaceFollower)

	return success, nil
}
func (s *spacesRetriever) UnFollow(ctx context.Context, request types.FollowRequest) (string, error) {
	var SpaceFollower model.SpaceFollower
	var user model.Jpa_web_authn_user
	var space model.Space
	deSession := s.db.Session(&gorm.Session{})
	deSession = deSession.Model(user).Where("username = ?", request.Username)
	res := deSession.First(&user)
	if res.Error != nil {
		return "false", nil
	}
	deSession = s.db.Session(&gorm.Session{})
	deSession = deSession.Model(SpaceFollower).Where("participantId = ? AND spaceId = ? ", user.Id, request.SpaceId)
	res = deSession.First(&SpaceFollower)
	if res.Error != nil {
		Follower := model.SpaceFollower{SpaceId: request.SpaceId, ParticipantId: user.Id, IsFollowing: false}
		result := deSession.Create(&Follower) // 通过数据的指针来创建
		if result.Error != nil {
			return "false", nil
		}
	}
	SpaceFollower.IsFollowing = false
	deSession.Save(&SpaceFollower)
	deSession = s.db.Session(&gorm.Session{})
	deSession = deSession.Model(space).Where("id=?", request.SpaceId)
	deSession = deSession.First(&space)

	space.Followers -= 1
	deSession.Save(&space)
	success := "UnFollow Success"
	return success, nil
}
