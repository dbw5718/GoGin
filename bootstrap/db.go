package bootstrap

import (
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"jassue-gin/app/models"
	"jassue-gin/global"
	"log"
	"os"
	"strconv"
	"time"
)

func InitializeDB() *gorm.DB {
	switch global.App.Config.Database.Driver {
	case "mysql":
		return initMySqlGorm()
	default:
		return initMySqlGorm()

	}
}
func initMySqlGorm() *gorm.DB {
	dbConfig := global.App.Config.Database
	if dbConfig.Database == "" {
		return nil
	}
	dsn := dbConfig.UserName + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" +
		dbConfig.Database + "?charset=" + dbConfig.Charset + "&parseTime=True&loc=Local"
	//fmt.Println("UserName:"+dbConfig.UserName)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         191,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   getGormLogger(),
	}); err != nil {
		global.App.Log.Error("mysql connect failed err:", zap.Any("err", err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		initMySqlTables(db)
		return db
	}
}
func initMySqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		models.User{},
	)
	if err != nil {
		global.App.Log.Error("migrate table failed", zap.Any("err", err))
		os.Exit(0)
	}
}
func getGromLogWriter() logger.Writer {
	var writer io.Writer

	if global.App.Config.Database.EnableFileLogWriter {
		writer = &lumberjack.Logger{
			Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.Database.LogFilename,
			MaxSize:    global.App.Config.Log.MaxSize,
			MaxAge:     global.App.Config.Log.MaxAge,
			MaxBackups: global.App.Config.Log.MaxBackups,
			Compress:   global.App.Config.Log.Compress,
		}
	} else {
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}
func getGormLogger() logger.Interface {
	var logMode logger.LogLevel
	switch global.App.Config.Database.LogMode {
	case "slient":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}
	return logger.New(getGromLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logMode,
		IgnoreRecordNotFoundError: false,
		Colorful:                  false,
	})
}
