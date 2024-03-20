package repository

import (
	"bookstack/model"
	"bookstack/utils/optional"

	"github.com/gofrs/uuid"
)

type CreateUserArgs struct {
	Name        string
	DisplayName string
	Icon        string
	Password    string
}

type UsersQuery struct {
	Name                 optional.Of[string]
	EnableProfileLoading bool
}

func (q UsersQuery) NameOf(name string) UsersQuery {
	q.Name = optional.From(name)
	return q
}

type UserRepository interface {
	// CreateUser 新規ユーザーを作成
	//
	// 成功した場合、ユーザーとnilを返す。
	// Nameが重複している場合、ErrAlreadyExistsを返す。
	CreateUser(args CreateUserArgs) (model.UserInfo, error)
	// GetUser 指定したIDのユーザーを取得
	//
	// 成功した場合、ユーザーとnilを返す。
	// ユーザーが存在しない場合、ErrNotFoundを返す。
	GetUser(id uuid.UUID, withProfile bool) (model.UserInfo, error)
	// GetUsers
	//
	// 成功した場合、ユーザー一覧とnilを返す。
	GetUsers(query UsersQuery) ([]model.UserInfo, error)
}
