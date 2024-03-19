package permission

const (
	// GetNote ノート取得権限
	GetNote = Permission("get_note")
	// CreateNote ノート作成権限
	CreateNote = Permission("create_note")
	// EditNote ノート編集権限
	EditNote = Permission("edit_note")
	// DeleteNote ノート削除権限
	DeleteNote = Permission("delete_note")
)
