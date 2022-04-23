package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goproject/main/ginweb/global"
	"goproject/main/ginweb/pkg/setting"
	"log"
	"time"
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
	//defer db.Close()

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)

	//自定义model回调
	db.Callback().Create().Replace("gorm:update_time_stamp", createTimeStampFormCreateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampFormCreateCallback)

	db.DB().SetMaxIdleConns(dbConfig.MaxIdleConns)
	db.DB().SetMaxOpenConns(dbConfig.MaxOpenConns)
	return db, nil
}

func createTimeStampFormCreateCallback(scope *gorm.Scope) { //创建记录时更新CreateOn字段
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		field, ok := scope.FieldByName("CreateOn") //判断是否存在createOn字段
		if ok {
			if field.IsBlank { //判断该字段是否为空
				_ = field.Set(nowTime) //设置值
			}
		}
	}
}

func updateTimeStampFormCreateCallback(scope *gorm.Scope) { //更新记录时更新ModifiedOn字段
	_, ok := scope.Get("gorm:update_column") //判断是否设置gorm:update_column属性
	if !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) { //在删除前判断是否存在DeleteOn与isDel
	if !scope.HasError() {
		var extraOption string
		str, ok := scope.Get("gorm:delete_option") //判断是否设置gorm:update_column属性
		if ok {
			extraOption = fmt.Sprint(str)
		}
		deleteField, deleteOk := scope.FieldByName("DeleteOn") //判断是否存在DeleteOn字段
		isDelField, isDelOk := scope.FieldByName("IsDel")      //判断是否存在DeleteOn字段
		if !scope.Search.Unscoped && deleteOk && isDelOk {
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf("update %v set %v=%v, %v=%v%v%v ",
				scope.QuotedTableName(),
				scope.Quote(deleteField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpace(scope.CombinedConditionSql()),
				addExtraSpace(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf("delete from %v%v%v",
				scope.QuotedTableName(),
				addExtraSpace(scope.CombinedConditionSql()),
				addExtraSpace(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpace(str string) string {
	if str == "" {
		return " " + str
	}
	return ""
}
