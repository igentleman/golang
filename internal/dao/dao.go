package dao

import (
	"goproject/main/ginweb/global"

	"github.com/jinzhu/gorm"
)

type Dao struct {
	Engine *gorm.DB
}

func New() *Dao {
	return &Dao{
		Engine: global.DbEngine,
	}
}
