package model

import (
	"time"

	"github.com/gofrs/uuid"
)

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

type UserProfile struct {
	UserID    uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	Bio       string    `gorm:"type:TEXT COLLATE utf8mb4_bin NOT NULL"`
	UpdatedAt time.Time `gorm:"precision:6"`
}
