package model

import (
	"EK-Server/config"
	"fmt"
	"math"
	"time"

	// _ 数据库驱动
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB DB
var DB *gorm.DB

// InitDB InitDB
func InitDB(DataBaseConfig string) *gorm.DB {
	t := time.Now().Format("2006年01-02 15:04:05")
	fmt.Printf("数据库链接>>>>>>>> %s \n", t)
	db, err := gorm.Open("mysql", DataBaseConfig)
	if err != nil {
		panic(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.Global.TablePrefix + "_" + defaultTableName
	}

	db.LogMode(true)
	db.AutoMigrate(&User{}, &WechatOauth{}, &Article{}, &API{}, &APICate{}, &Goods{}, &GoodsCate{})
	return db
}

// DataBaselimit  获取所有用户列表
func DataBaselimit(limit int, page int, model interface{}, list interface{}, table string, Order string) map[string]interface{} {
	db := DB
	// 用户列表
	// users := []model{}
	// 初始化数据库对象
	userModal := db.Table(config.Global.TablePrefix + "_" + table).Where(model)
	// 总数
	var count int
	// 绑定总数
	userModal.Count(&count)
	// 查询绑定用户列表
	userModal.Offset(limit * (page - 1)).Limit(limit).Order(Order).Find(list)
	pageCount := float64(count) / float64(limit)
	return map[string]interface{}{
		"count":     count,
		"list":      list,
		"pageSize":  limit,
		"pageNow":   page,
		"pageCount": math.Ceil(pageCount),
	}
}

//TableName 拼接默认表名
func TableName(str string) (result string) {
	result = config.Global.TablePrefix + "_" + str
	return result
}
