package repository

import "bookstack/model"

type CreateUserArgs struct {
	Name        string
	DisplayName string
	Icon        string
	Password    string
}

type UserRepository interface {
	FindAll() ([]model.User, error)
	// CreateUser 新規ユーザーを作成
	//
	// 成功した場合、ユーザーとnilを返す。
	// Nameが重複している場合、ErrAlreadyExistsを返す。
	CreateUser(args CreateUserArgs) (*model.User, error)
}
