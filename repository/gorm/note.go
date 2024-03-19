package gorm

import (
	"bookstack/model"
	"bookstack/repository"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type noteRepository struct {
	db *gorm.DB
}

func makeNoteRepository(db *gorm.DB) *noteRepository {
	return &noteRepository{db}
}

func (r *noteRepository) CreateNote(args repository.CreateNoteArgs) (*model.Note, error) {
	if userId := args.UserID; userId == uuid.Nil {
		return nil, repository.ErrNilID
	}
	note := &model.Note{
		UserID: args.UserID,
		Text:   args.Text,
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
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
