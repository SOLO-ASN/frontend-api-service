package types

type CheckDuplicateRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateUserRequest struct {
	Name   string `json:"name" binding:""`
	Avatar string `json:"avatar" binding:""`
	Email  string `json:"email" binding:""`
}

type GetUserResponse struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	UUID   string `json:"uuid" binging:""`
	Email  string `json:"email"`
}

type UpdateSocialAccountRequest struct {
	XAccount        x `json:"xAccount" binding:""`
	GithubAccount   x `json:"githubAccount" binding:""`
	DiscordAccount  x `json:"discordAccount" binding:""`
	TelegramAccount x `json:"telegramAccount" binding:""`
}

type x struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
