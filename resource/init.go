package resource

import (
	"github.com/gogf/gf/os/gcache"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/vearne/tinyurl/config"
	zlog "github.com/vearne/tinyurl/log"
	"go.uber.org/zap"
	"os"
	"time"
)

const (
	MaxCacheCapacity = 10000
)

var (
	MySQLClient *gorm.DB
)

var (
	FixedCache *gcache.Cache
)

func InitResource() {
	zlog.Info("InitCache")
	initCache()
	zlog.Info("InitMySQL")
	initMySQL()
}

func initCache() {
	FixedCache = gcache.New(MaxCacheCapacity)
}

func initMySQL() {
	mysqlConf := config.GetOpts().MySQLConf
	mysqldb, err := gorm.Open("mysql", mysqlConf.DSN)
	if err != nil {
		zlog.Error("initialize_db error", zap.Error(err))
		os.Exit(4)
	}
	if mysqlConf.Debug {
		mysqldb = mysqldb.Debug()
	}
	mysqldb.DB().SetMaxIdleConns(mysqlConf.MaxIdleConn)
	mysqldb.DB().SetMaxOpenConns(mysqlConf.MaxOpenConn)
	mysqldb.DB().SetConnMaxLifetime(time.Duration(mysqlConf.ConnMaxLifeSecs) * time.Second)
	// 赋值给全局变量
	MySQLClient = mysqldb
}
