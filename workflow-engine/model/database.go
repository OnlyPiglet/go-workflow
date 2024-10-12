package model

import (
	"fmt"
	config "github.com/OnlyPiglet/go-workflow/workflow-config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

// 配置
var conf = *config.Config

// Setup 初始化一个db连接
func Setup() {
	newLogger := gormLogger.New(log.New(os.Stdout, "\r\n", 0), gormLogger.Config{
		LogLevel: gormLogger.Info,
		Colorful: true,
	})
	//log.New(flog.DefaultLogger.GetWriter(), "\r\n", log.LstdFlags), // io writer
	//gormLogger.Config{
	//	SlowThreshold:             3 * time.Second, // Slow SQL threshold
	//	LogLevel:                  logLevel,        // Log level
	//	IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
	//	Colorful:                  true,            // Disable color
	//},
	//)
	var err error
	log.Println("数据库初始化！！")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName)
	idb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	//db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName))
	if err != nil {
		log.Fatalf("数据库连接失败 err: %v", err)
	}
	DB, err := idb.DB()

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(30 * time.Second)
	// 启用Logger，显示详细日志
	SetUpWithDb(idb)
}

func SetUpWithDb(db *gorm.DB) {
	db = db
	err := db.AutoMigrate(&Procdef{},
		&Execution{},
		&Task{},
		&Identitylink{},
		&ProcInst{},
		&ExecutionHistory{},
		&IdentitylinkHistory{},
		&ProcInstHistory{},
		&TaskHistory{},
		&ProcdefHistory{})
	if err != nil {
		panic(err)
		return
	}
	DB = db
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	//defer db.Close()
}

// GetDB getdb
func GetDB() *gorm.DB {
	return DB
}

// GetTx GetTx
func GetTx() *gorm.DB {
	return DB.Begin()
}
