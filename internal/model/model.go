package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goproject/main/ginweb/global"
	"goproject/main/ginweb/pkg/setting"
	"log"
)

type Model struct {
	ID         uint   `gorm:"primary_key" json:"id"` //设置ID为默认主键 返回json数据格式为id
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint   `json:"created_on"`
	ModifiedOn uint   `json:"modified_on"`
	DeletedOn  uint   `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(dbConfig *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(dbConfig.DBType, fmt.Sprintf("%v:%v@/%v?charset=%v&parseTime=%v&loc=Local", dbConfig.UserName, dbConfig.Password, dbConfig.DBName, dbConfig.Charset, dbConfig.ParseTime))
	if err != nil {
		log.Fatal("数据库连接失败，err=", err)
	}
	defer db.Close()

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(dbConfig.MaxIdleConns)
	db.DB().SetMaxOpenConns(dbConfig.MaxOpenConns)
	return db, nil
}
