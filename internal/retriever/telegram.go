package retriever

import (
	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"
	"api-service/internal/types"
	"context"

	"gorm.io/gorm"
)

type Telegram interface {
	isFollow(ctx context.Context, request types.FollowRequest) (string, error)
}

type telegram struct {
	db    *gorm.DB
	cache *cache.Cache
}

func (s *spacesRetriever) isFollow(ctx context.Context, request types.FollowRequest) (string, error) {
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
