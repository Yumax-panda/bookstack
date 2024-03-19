package repository

// Repository データリポジトリ
type Repository interface {
	UserRepository
	NoteRepository
}
