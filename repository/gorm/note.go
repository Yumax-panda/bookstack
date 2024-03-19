package gorm

import (
	"bookstack/model"
	"bookstack/repository"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func (repo *Repository) CreateNote(args repository.CreateNoteArgs) (*model.Note, error) {
	if userId := args.UserID; userId == uuid.Nil {
		return nil, repository.ErrNilID
	}
	note := &model.Note{
		UserID: args.UserID,
		Title:  args.Title,
		Text:   args.Text,
	}

	err := repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(note).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return note, nil
}
