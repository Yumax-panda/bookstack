package model

import (
	"bookstack/utils"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"time"

	"github.com/gofrs/uuid"
)

var (
	// ErrUserWrongIDOrPassword : ユーザーエラー IDかパスワードが間違っています。
	ErrUserWrongIDOrPassword = errors.New("password or id is wrong")
)

type UserInfo interface {
	GetID() uuid.UUID
	GetName() string
	GetDisplayName() string
	GetIcon() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time

	GetBio() string
	Authenticate(password string) error
}

type User struct {
	ID          uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	Name        string    `gorm:"type:varchar(32);not null;unique"`
	DisplayName string    `gorm:"type:varchar(32);not null;default:''"`
	Password    string    `gorm:"type:char(128);not null;default:''"`
	Salt        string    `gorm:"type:char(128);not null;default:''"`
	Icon        string    `gorm:"type:TEXT COLLATE utf8mb4_bin NOT NULL"`
	CreatedAt   time.Time `gorm:"precision:6"`
	UpdatedAt   time.Time `gorm:"precision:6"`

	Profile *UserProfile `gorm:"constraint:user_profiles_user_id_users_id_foreign,OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (User) TableName() string {
	return "users"
}

type UserProfile struct {
	UserID    uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	Bio       string    `gorm:"type:TEXT COLLATE utf8mb4_bin NOT NULL"`
	UpdatedAt time.Time `gorm:"precision:6"`
}

func (UserProfile) TableName() string {
	return "user_profiles"
}

func (u *User) GetID() uuid.UUID {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetDisplayName() string {
	return u.DisplayName
}

func (u *User) GetIcon() string {
	return u.Icon
}

func (u *User) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (u *User) GetBio() string {
	if u.Profile == nil {
		panic("unexpected control flow: user profile is not loaded")
	}
	return u.Profile.Bio
}

func (u *User) Authenticate(password string) error {
	if len(u.Password) == 0 || len(u.Salt) == 0 {
		return ErrUserWrongIDOrPassword
	}

	storedPassword, err := hex.DecodeString(u.Password)

	if err != nil {
		return ErrUserWrongIDOrPassword
	}

	if subtle.ConstantTimeCompare(storedPassword, utils.HashPassword(password, []byte(u.Salt))) != 1 {
		return ErrUserWrongIDOrPassword
	}

	return nil
}
