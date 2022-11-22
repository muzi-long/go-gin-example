package model

import (
	"context"

	"gorm.io/gorm"
)

type (
	UserModel interface {
		FindOne(ctx context.Context, id int64) (*User, error)
		Insert(ctx context.Context, data *User) (*User, error)
		Update(ctx context.Context, data *User) (*User, error)
	}
	DefaultUserModel struct {
		conn  *gorm.DB
		table string
	}
	User struct {
		ID   uint   `gorm:"column:id"`
		Name string `gorm:"column:name"`
	}
)

func NewUserModel(conn *gorm.DB) UserModel {
	return &DefaultUserModel{
		conn:  conn,
		table: "user",
	}
}

func (m *DefaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	user := new(User)
	err := m.conn.Table(m.table).First(user, id).Error
	return user, err
}

func (m *DefaultUserModel) Insert(ctx context.Context, data *User) (*User, error) {
	err := m.conn.Table(m.table).Create(data).Error
	return data, err
}

func (m *DefaultUserModel) Update(ctx context.Context, data *User) (*User, error) {
	err := m.conn.Table(m.table).Save(data).Error
	return data, err
}
