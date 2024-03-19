package repository

import (
	"bookstack/model"

	"github.com/gofrs/uuid"
)

type CreateNoteArgs struct {
	UserID uuid.UUID
	Title  string
	Text   string
}

type NoteRepository interface {
	// CreateNote 新規ノートを作成
	//
	// 成功した場合、ノートとnilを返す。
	CreateNote(args CreateNoteArgs) (*model.Note, error)
}
