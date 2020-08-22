package libs

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var DB *gorm.DB

type DbConfig struct {
	Host         string
	Port         string
	Database     string
	User         string
	Password     string
	Charset      string
	MaxIdleConns int
	MaxOpenConns int
}

func (c *DbConfig) InitDB() *gorm.DB {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database, c.Charset)
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}
	db.SingularTable(true)                  //全局设置表名不可以为复数形式。
	db.DB().SetMaxIdleConns(c.MaxIdleConns) //空闲时最大的连接数
	db.DB().SetMaxOpenConns(c.MaxOpenConns) //最大的连接数
	return db
}
