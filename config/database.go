package config

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"starOJ-backend/model"
)

var db *gorm.DB

func GetDB() *gorm.DB { //提供接口，允许从连接池获取数据库连接
	return db
}

func init() {

	username := "root"  //账号
	password := "root"  //密码
	host := "127.0.0.1" //数据库地址，可以是ip或者域名
	port := 3306        //数据库端口
	dbname := "starOJ"  //数据库名
	timeout := "5s"     //超时时限

	lnk := fmt.Sprintf("%s:%s@tcp(%s:%d)/", username, password, host, port)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, dbname, timeout)

	con, err := sql.Open("mysql", lnk)
	if err != nil {
		panic("无法连接到数据库, error = " + err.Error())
	}
	_, err = con.Exec("USE staroj")
	if err != nil {
		_, err = con.Exec("CREATE DATABASE" + " " + dbname)
		if err != nil {
			panic("不存在初始数据库且无法创建!")
		}
		fmt.Println("检测到尚未拥有数据库，已创建初始数据库")
	}
	con.Close()

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 单数表名
		},
	})
	if err != nil {
		panic("无法连接到数据库, error = " + err.Error())
	}

	initMigrate() //自动同步数据库与当前模型

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10) //最大连接数
	sqlDB.SetMaxIdleConns(5)  //空闲时最大连接数
}

func initMigrate() {
	model.MigrateUser(db)
	model.MigrateProblem(db)
}
