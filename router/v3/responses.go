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
	DisplayName string    `json:"display_name"`
	Icon        string    `json:"icon"`
	UpdatedAt   time.Time `json:"update_at"`
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
