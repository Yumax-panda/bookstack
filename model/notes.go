package model

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Note struct {
	ID        uuid.UUID      `gorm:"type:char(36);not null;primaryKey"`
	UserID    uuid.UUID      `gorm:"type:char(36);not null;"`
	Text      string         `gorm:"type:TEXT COLLATE utf8mb4_bin NOT NULL"`
	CreatedAt time.Time      `gorm:"precision:6;index;index:idx_messages_channel_id_deleted_at_created_at,priority:3;index:idx_messages_deleted_at_created_at,priority:2"`
	UpdatedAt time.Time      `gorm:"precision:6;index:idx_messages_deleted_at_updated_at,priority:2"`
	DeletedAt gorm.DeletedAt `gorm:"precision:6;index:idx_messages_channel_id_deleted_at_created_at,priority:2;index:idx_messages_deleted_at_created_at,priority:1;index:idx_messages_deleted_at_updated_at,priority:1"`

	User *User `gorm:"constraint:notes_user_id_users_id_foreign,OnUpdate:CASCADE,OnDelete:CASCADE"`
}
