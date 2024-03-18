package gorm

import (
	"bookstack/model"
	"bookstack/repository"
	"bookstack/utils"
	"bookstack/utils/gormutil"
	"bookstack/utils/random"
	"encoding/hex"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func makeUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

// FindAll implements UserRepository interface
func (r *userRepository) FindAll() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser implements UserRepository interface
func (r *userRepository) CreateUser(args repository.CreateUserArgs) (*model.User, error) {
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
