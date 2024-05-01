package database

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	blog_mysql      *gorm.DB
	blog_mysql_once sync.Once
)

func init() {
	//
}

func createMysqlDB(dbname, host, user, pass string, port int) *gorm.DB {
	// data source name 是 tester:123456@tcp(localhost:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname) //mb4兼容emoji表情符号

	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
	if err != nil {
		panic("Cannot connect DB ")
	}
	//设置数据库连接池参数，提高并发性能
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
	fmt.Println("Connected to mysql database")
	return db
}

func GetTimeToFitDBConnection() *gorm.DB {
	blog_mysql_once.Do(func() {
		if blog_mysql == nil {
			dbName := "timetofit"
			host := "localhost"
			port := 3306
			user := "root"
			pass := "zhouli1118"
			blog_mysql = createMysqlDB(dbName, host, user, pass, port)
		}
	})
	return blog_mysql
}
