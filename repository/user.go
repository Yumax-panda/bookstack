package repository

import (
	"bookstack/model"

	"github.com/gofrs/uuid"
)

type CreateUserArgs struct {
	Name        string
	DisplayName string
	Icon        string
	Password    string
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
}
