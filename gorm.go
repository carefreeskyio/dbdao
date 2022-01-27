// +build gorm

package dbdao

import (
	"github.com/carefreex-io/config"
	"github.com/carefreex-io/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"sync"
	"time"
)

var(
	Read  *gorm.DB
	Write *gorm.DB
)

type Log struct {
	ShowLog                   bool
	SlowThreshold             time.Duration
	IgnoreRecordNotFoundError bool
	LogLevel                  gormLog.LogLevel
}

type BaseOptions struct {
	ReadDns     string
	WriteDns    string
	Logger      Log
	MaxIdleConn int
	MaxOpenConn int
	MaxLifetime time.Duration
	TablePrefix string
}

type CustomOptions struct {
	GormConfig *gorm.Config
}

var (
	dbOnce               sync.Once
	baseOptions          *BaseOptions
	DefaultCustomOptions = &CustomOptions{
		GormConfig: &gorm.Config{},
	}
	defaultSessionOptions = &gorm.Session{
		NewDB: true,
	}
)

func initBaseOptions() {
	logLevel := gormLog.Info
	if !config.GetBool("MysqlConf.Log.ShowSql") {
		logLevel = gormLog.Silent
	}
	baseOptions = &BaseOptions{
		ReadDns:     config.GetString("Mysql.Read"),
		WriteDns:    config.GetString("Mysql.Write"),
		MaxIdleConn: config.GetInt("MysqlConf.MaxIdleConn"),
		MaxOpenConn: config.GetInt("MysqlConf.MaxOpenConn"),
		MaxLifetime: config.GetDuration("MysqlConf.MaxLifetime") * time.Second,
		Logger: Log{
			SlowThreshold:             config.GetDuration("MysqlConf.Log.SlowThreshold") * time.Millisecond,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: config.GetBool("MysqlConf.Log.IgnoreRecordNotFoundError"),
		},
		TablePrefix: config.GetString("MysqlConf.TablePrefix"),
	}
}

func InitDB() (err error) {
	dbOnce.Do(func() {
		initBaseOptions()
		if baseOptions.ReadDns == "" && baseOptions.WriteDns == "" {
			return
		}

		DefaultCustomOptions.GormConfig.Logger = New(gormLog.Config{
			SlowThreshold:             baseOptions.Logger.SlowThreshold,
			LogLevel:                  baseOptions.Logger.LogLevel,
			IgnoreRecordNotFoundError: baseOptions.Logger.IgnoreRecordNotFoundError,
			Colorful:                  false,
		})
		DefaultCustomOptions.GormConfig.NamingStrategy = schema.NamingStrategy{
			TablePrefix:   baseOptions.TablePrefix,
			SingularTable: true,
		}
		if Read, err = gorm.Open(mysql.Open(baseOptions.ReadDns), DefaultCustomOptions.GormConfig); err != nil {
			logger.Errorf("gorm.Open failed: err=%v", err)
			return
		}
		if err = setConfig(Read); err != nil {
			logger.Errorf("set read db config failed: err=%v", err)
			return
		}

		DefaultCustomOptions.GormConfig.Logger = New(gormLog.Config{
			SlowThreshold:             baseOptions.Logger.SlowThreshold,
			LogLevel:                  baseOptions.Logger.LogLevel,
			IgnoreRecordNotFoundError: baseOptions.Logger.IgnoreRecordNotFoundError,
			Colorful:                  false,
		})
		if Write, err = gorm.Open(mysql.Open(baseOptions.WriteDns), DefaultCustomOptions.GormConfig); err != nil {
			logger.Errorf("gorm.Open failed: err=%v", err)
			return
		}
		if err = setConfig(Write); err != nil {
			logger.Errorf("set write db config failed: err=%v", err)
			return
		}
	})

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
	return Read.Session(defaultSessionOptions)
}

func NewReadSessionWithOptions(options *gorm.Session) (db *gorm.DB) {
	return Read.Session(options)
}

func NewWriteSession() (db *gorm.DB) {
	return Write.Session(defaultSessionOptions)
}

func NewWriteSessionWithOptions(options *gorm.Session) (db *gorm.DB) {
	return Write.Session(options)
}
