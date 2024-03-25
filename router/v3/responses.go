package v3

import (
	"bookstack/model"
	"sort"
	"time"

	"github.com/gofrs/uuid"
)

type UserResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	DisplayName string    `json:"displayName"`
	Icon        string    `json:"icon"`
	UpdatedAt   time.Time `json:"updateAt"`
}

func formatUsers(users []model.UserInfo) []UserResponse {
	res := make([]UserResponse, len(users))
	for i, u := range users {
		res[i] = UserResponse{
			ID:          u.GetID(),
			Name:        u.GetName(),
			DisplayName: u.GetDisplayName(),
			Icon:        u.GetIcon(),
			UpdatedAt:   u.GetUpdatedAt(),
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].ID.String() < res[j].ID.String()
	})
	return res
}

type UserDetailResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	DisplayName string    `json:"displayName"`
	Icon        string    `json:"icon"`
	UpdatedAt   time.Time `json:"updateAt"`
	Bio         string    `json:"bio"`
}

func formatUserDetail(user model.UserInfo) *UserDetailResponse {
	return &UserDetailResponse{
		ID:          user.GetID(),
		Name:        user.GetName(),
		DisplayName: user.GetDisplayName(),
		Icon:        user.GetIcon(),
		UpdatedAt:   user.GetUpdatedAt(),
		Bio:         user.GetBio(),
	}
}
