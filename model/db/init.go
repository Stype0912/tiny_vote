package db

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB
var err error

const user = "root"
const passwd = "123456"
const scheme = "tiny_vote"

func Init() {
	var env, addr string
	env = os.Getenv("GIN_MODE")
	// localhost is for local test and release mode is for docker call.
	if env == gin.ReleaseMode {
		addr = "mysql-tiny-vote:3307"
	} else {
		addr = "localhost:3307"
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v", user, passwd, addr, scheme)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Fail to connect to the db, error is " + err.Error())
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(1000)
	sqlDB.SetMaxIdleConns(20)
}
