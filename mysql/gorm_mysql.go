package mysql

import (
	"github.com/carefreex-io/config"
	"github.com/carefreex-io/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

type DB struct {
	Read  *gorm.DB
	Write *gorm.DB
}

type BaseOptions struct {
	ReadDns     string
	WriteDns    string
	MaxIdleConn int
	MaxOpenConn int
	MaxLifetime time.Duration
}

type CustomOptions struct {
	GormConfig *gorm.Config
}

var (
	Db                   *DB
	dbOnce               sync.Once
	baseOptions          *BaseOptions
	DefaultCustomOptions = &CustomOptions{
		GormConfig: &gorm.Config{
			Logger: New(),
		},
	}
	defaultSessionOptions = &gorm.Session{
		NewDB: true,
	}
)

func initBaseOptions() {
	baseOptions = &BaseOptions{
		ReadDns:     config.GetString("Mysql.Read"),
		WriteDns:    config.GetString("Mysql.Write"),
		MaxIdleConn: config.GetInt("MysqlConf.MaxIdleConn"),
		MaxOpenConn: config.GetInt("MysqlConf.MaxOpenConn"),
		MaxLifetime: config.GetDuration("MysqlConf.MaxLifetime") * time.Second,
	}
}

func InitDB() (err error) {
	if Db == nil {
		initBaseOptions()

		dbOnce.Do(func() {
			if Db.Read, err = gorm.Open(mysql.Open(baseOptions.ReadDns), DefaultCustomOptions.GormConfig); err != nil {
				logger.Errorf("gorm.Open failed: err=%v", err)
				return
			}
			if err = setConfig(Db.Read); err != nil {
				logger.Errorf("set read db config failed: err=%v", err)
				return
			}
			if Db.Write, err = gorm.Open(mysql.Open(baseOptions.WriteDns), DefaultCustomOptions.GormConfig); err != nil {
				logger.Errorf("gorm.Open failed: err=%v", err)
				return
			}
			if err = setConfig(Db.Write); err != nil {
				logger.Errorf("set write db config failed: err=%v", err)
				return
			}
		})
	}

	return err
}

func setConfig(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		logger.Errorf("get sqlDB failed: err=%v", err)
		return err
	}
	sqlDB.SetMaxIdleConns(baseOptions.MaxIdleConn)
	sqlDB.SetMaxOpenConns(baseOptions.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(baseOptions.MaxLifetime)

	return nil
}

func NewReadSession() (db *gorm.DB) {
	return Db.Read.Session(defaultSessionOptions)
}

func NewReadSessionWithOptions(options *gorm.Session) (db *gorm.DB) {
	return Db.Read.Session(options)
}

func NewWriteSession() (db *gorm.DB) {
	return Db.Write.Session(defaultSessionOptions)
}

func NewWriteSessionWithOptions(options *gorm.Session) (db *gorm.DB) {
	return Db.Write.Session(options)
}
