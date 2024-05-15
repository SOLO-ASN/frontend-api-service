package retriever

import (
	"api-service/internal/dbEntity/cache"
	"api-service/internal/model"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UserRetriever interface {
	Create(ctx context.Context, table *model.User) error
	GetById(ctx context.Context, uuid string) (*model.User, error)
	GetByName(ctx context.Context, name string) (*model.User, error)
	CheckDuplicateName(ctx context.Context, table *model.User) bool
	UpdateSocialAccountById(ctx context.Context, username string, table *model.User) error
	UpdateEmailById(ctx context.Context, table *model.User) error
	UpdateById(ctx context.Context, table *model.User) error
	DeleteById(ctx context.Context, uuid string) error
}

type userRetriever struct {
	db    *gorm.DB
	cache cache.Cache
}

func NewUserRetriever(db *gorm.DB, cache cache.Cache) UserRetriever {
	return &userRetriever{
		db:    db,
		cache: cache,
	}
}

func (u userRetriever) Create(ctx context.Context, table *model.User) error {
	err := u.db.WithContext(ctx).Create(table).Error
	// todo delete cache
	return err
}

func (u userRetriever) GetById(ctx context.Context, uuid string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRetriever) GetByName(ctx context.Context, name string) (*model.User, error) {
	//TODO implement me
	user := &model.User{}

	err := u.db.Model(user).Where("name = ?", name).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	fmt.Println(user)
	return user, nil
}

func (u userRetriever) CheckDuplicateName(ctx context.Context, table *model.User) bool {
	if err := u.db.First(table, "name=?", table.Name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// record not found
			return false
		} else {
			// other errors
			return false
		}
	} else {
		return true
	}
}

func (u userRetriever) UpdateSocialAccountById(ctx context.Context, username string, table *model.User) error {
	updates := make(map[string]interface{})

	// check
	if table.SocialAccount.XAccountId != "" {
		updates["x_account_id"] = table.SocialAccount.XAccountId
		updates["x_account_name"] = table.SocialAccount.XAccountName
	}

	if table.SocialAccount.GithubAccountId != "" {
		updates["github_account_id"] = table.SocialAccount.GithubAccountId
		updates["github_account_name"] = table.SocialAccount.GithubAccountName
	}

	if table.SocialAccount.DiscordAccountId != "" {
		updates["discord_account_id"] = table.SocialAccount.DiscordAccountId
		updates["discord_account_name"] = table.SocialAccount.DiscordAccountName
	}

	if table.SocialAccount.TelegramAccountId != "" {
		updates["telegram_account_id"] = table.SocialAccount.TelegramAccountId
		updates["telegram_account_name"] = table.SocialAccount.TelegramAccountName
	}
	err := u.db.Model(table).Where("name = ?", username).First(&model.User{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user not found")
	}

	if err = u.db.Model(table).Where("name = ?", username).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func (u userRetriever) UpdateEmailById(ctx context.Context, table *model.User) error {
	updates := make(map[string]interface{})

	// check
	if table.Email != "" {
		updates["email"] = table.Email
	}

	if err := u.db.Model(table).Where("id =?", table.ID).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func (u userRetriever) UpdateById(ctx context.Context, table *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userRetriever) DeleteById(ctx context.Context, uuid string) error {
	//TODO implement me
	panic("implement me")
}
