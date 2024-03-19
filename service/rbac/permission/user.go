package permission

const (
	// GetUser ユーザー情報取得権限
	GetUser = Permission("get_user")
	// RegisterUser ユーザー登録権限
	RegisterUser = Permission("register_user")
	// GetMe 自分のユーザー情報取得権限
	GetMe = Permission("get_me")
	// EditMe 自分のユーザー情報編集権限
	EditMe = Permission("edit_me")
)
