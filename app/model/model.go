package model

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type BaseModel[m any] interface {
	First(ctx context.Context, id uint) (m, error)
}

type Base struct {
	*gorm.DB
}

func NewBase(orm *gorm.DB) *Base {
	return &Base{orm}
}

func (this *Base) First(ctx context.Context, id uint) {

}
