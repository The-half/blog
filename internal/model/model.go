package model

import (
	"blog/global"
	"blog/pkg/setting"
	"fmt"
	otgorm "github.com/eddycjy/opentracing-gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type Model struct {
	ID uint32 `gorm:"primary_key" json:"id,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	ModifiedBy string `json:"modified_by,omitempty"`
	CreatedOn uint32 `json:"created_on,omitempty"`
	ModifiedOn uint32 `json:"modified_on,omitempty"`
	DeletedOn uint32 `json:"deleted_on,omitempty"`
	IsDel uint8 `json:"is_del,omitempty"`
}


type ArticleTag struct {
	*Model
	TagID uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (at ArticleTag) TableName() string {
	return "blog_article_tag"
}

//创建一个数据库引擎
func NewDBEngine (databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
		))

	if err != nil {
		return nil, err
	}

	//设置运行模式为debug
	if global.ServerSetting.RunMode == "debug" {
		//设置打印日志
		db.LogMode(true)
	}


	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp",
		updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp",
		updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete",deleteCallback)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	//增加gorm的链路追踪
	otgorm.AddGormCallbacks(db)

	return db, nil
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeFiled, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeFiled.IsBlank {
				_ = createTimeFiled.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}


func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}


func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm: delete_option"); !ok {
			extraOption = fmt.Sprint(str)
		}

		deleteOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
			nowTime := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v = %v, %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deleteOnField.DBName),
				scope.AddToVars(nowTime),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
				)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FORM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
				)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}