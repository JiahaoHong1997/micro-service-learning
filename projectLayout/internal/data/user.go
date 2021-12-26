package data

import (
	"github.com/gotomicro/ego-component/egorm"
	"github.com/gotomicro/ego-component/eredis"
)

type UserRepo struct {
	db    *egorm.Component
	cache *eredis.Component
}

type DB struct {
	dsn      string
	username string
	password string
	cfg      DBConfig
}

type DBConfig struct {
	dsn      string
	username string
	password string
}

func NewUserRepo(db *egorm.Component, cache *eredis.Component) *UserRepo {
	return &UserRepo{
		db:    db,
		cache: cache,
	}
}

func (u *UserRepo) GetUser(id uint64) (*User, error) {
	var user User
	err := u.db.Where("id = ?", id).Find(&user).Error
	return &user, err
}

type User struct {
	Nickname string `gorm:"not null" json:"name"`
}
