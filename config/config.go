package config

import (
	"github.com/spf13/viper"
	"sync"
	"sync/atomic"
)

//Config global config
var gcf atomic.Value
var initOnce sync.Once

//AppConfig  config info
type AppConfig struct {
	Logger struct {
		Level    string
		FilePath string
	}
	MySQLConf struct {
		DSN             string
		MaxIdleConn     int
		MaxOpenConn     int
		ConnMaxLifeSecs int
		// 是否启动debug模式
		// 若开启则会打印具体的执行SQL
		Debug bool
	}
	Web struct {
		ListenAddress string
		Mode          string
	}
	// 短域名
	Domain string
}

func InitConfig() error {
	initOnce.Do(func() {
		var cf = AppConfig{}
		cf.initWebConfig()
		cf.initLogger()
		cf.initMySQLConf()
		gcf.Store(&cf)
	})
	return nil
}
func GetOpts() *AppConfig {
	return gcf.Load().(*AppConfig)
}

func (c *AppConfig) initLogger() {
	c.Logger.Level = viper.GetString("logger.level")
	c.Logger.FilePath = viper.GetString("logger.filepath")
}

func (c *AppConfig) initMySQLConf() {
	c.MySQLConf.DSN = viper.GetString("mysql.dsn")
	c.MySQLConf.MaxIdleConn = viper.GetInt("mysql.max_idle_conn")
	c.MySQLConf.MaxOpenConn = viper.GetInt("mysql.max_open_conn")
	c.MySQLConf.ConnMaxLifeSecs = viper.GetInt("mysql.conn_max_life_secs")
	c.MySQLConf.Debug = viper.GetBool("mysql.debug")

}

func (c *AppConfig) initWebConfig() {
	c.Web.Mode = viper.GetString("web.mode")
	c.Web.ListenAddress = viper.GetString("web.listen_address")
	c.Domain = viper.GetString("domain")
}
