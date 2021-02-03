package dbconnect

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// 数据库连接
var PgGormDb *gorm.DB

func init()  {
	initViper()
	InitSQLite3()
}

//sqlite连接
func InitSQLite3() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)
	gdb, gerr := gorm.Open(sqlite.Open("fund.db"), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
		},
	})
	if gerr != nil {
		panic(gerr)
	}
	sqlDB, err := gdb.DB()
	if err != nil {
		fmt.Println(err)
	}
	sqlDB.SetMaxIdleConns(3)
	sqlDB.SetMaxOpenConns(30)
	PgGormDb = gdb
	log.Println(">>>>>>>>>>>SQLite init success")
}

//初始化配置文件读取
func initViper() {
	viper.SetConfigName("config")                                // name of config file (without extension)
	viper.SetConfigType("yaml")                                  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/home/ubuntu/goProject/fund") // path to look for the config file in
	viper.AddConfigPath("$HOME/.fund")                 // call multiple times to add many search paths
	viper.AddConfigPath(".")                                     // optionally look for config in the working directory 	// optionally look for config in the working directory
	err := viper.ReadInConfig()                                  // Find and read the config file
	if err != nil {                                              // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}
