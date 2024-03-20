package gorm

import (
	"bookstack/model"
	"bookstack/repository"
	"bookstack/utils"
	"bookstack/utils/gormutil"
	"bookstack/utils/random"
	"context"
	"encoding/hex"
	"time"

	"github.com/gofrs/uuid"
	"github.com/motoki317/sc"
	"gorm.io/gorm"
)

type getUserArg struct {
	id          uuid.UUID
	withProfile bool
}

var _ repository.UserRepository = (*userRepository)(nil)

type userRepository struct {
	db    *gorm.DB
	users *sc.Cache[getUserArg, model.UserInfo]
}

func makeUserRepository(db *gorm.DB) *userRepository {
	r := &userRepository{db: db}
	r.users = sc.NewMust(r.getUser, 1*time.Hour, 1*time.Hour)
	return r
}

func (r *userRepository) getUser(_ context.Context, arg getUserArg) (model.UserInfo, error) {
	tx := r.db
	if arg.withProfile {
		tx = tx.Preload("Profile")
	}
	var user model.User
	if err := tx.First(&user, &model.User{ID: arg.id}).Error; err != nil {
		return nil, convertError(err)
	}
	return &user, nil
}

func (r *userRepository) forgetCache(id uuid.UUID) {
	r.users.Forget(getUserArg{id: id, withProfile: false})
	r.users.Forget(getUserArg{id: id, withProfile: true})
}

// CreateUser implements UserRepository interface
func (r *userRepository) CreateUser(args repository.CreateUserArgs) (model.UserInfo, error) {
	uid := uuid.Must(uuid.NewV4())
	user := &model.User{
		ID:          uid,
		Name:        args.Name,
		DisplayName: args.DisplayName,
		Icon:        args.Icon,
		Profile: &model.UserProfile{
			UserID: uid,
		},
	}

	if len(args.Password) > 0 {
		salt := random.Salt()
		user.Password = hex.EncodeToString(utils.HashPassword(args.Password, salt))
		user.Salt = hex.EncodeToString(salt)
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if exists, err := gormutil.RecordExists(tx, &model.User{Name: user.Name}); err != nil {
			return err
		} else if exists {
			return repository.ErrAlreadyExists
		}

		// userとuser_profileを作成
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUser(id uuid.UUID, withProfile bool) (model.UserInfo, error) {
	if id == uuid.Nil {
		return nil, repository.ErrNotFound
	}
	return r.users.Get(context.Background(), getUserArg{id, withProfile})
}

func (r *userRepository) GetUsers(query repository.UsersQuery) ([]model.UserInfo, error) {
	arr := make([]model.User, 0)
	if err := r.makeGetUsersTx(query).Find(&arr).Error; err != nil {
		return nil, err
	}
	users := make([]model.UserInfo, len(arr))

	for i, u := range arr {
		users[i] = &u
	}
	return users, nil
}

func (r *userRepository) makeGetUsersTx(query repository.UsersQuery) *gorm.DB {
	tx := r.db.Table("users")

	if query.Name.Valid {
		tx = tx.Where("users.name = ?", query.Name.V)
	}

	if query.EnableProfileLoading {
		tx = tx.Preload("Profile")
	}

	return tx
}
