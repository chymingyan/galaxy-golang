package initialize


import (
	"galaxy-golang/server/global"
	"galaxy-golang/server/initialize/include"
	"gorm.io/gorm/logger"
	"os"

	//"gin-vue-admin/initialize/internal"
	//"gin-vue-admin/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"os"
)

//@author: chy
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "postgresql":
		return GormPostgresql()
	case "sqllite3":
		return GormSqlite3()
	case "oracle":
		return GormOracle()
	case "dameng":
		return GormDameng()
	default:
		return GormMysql()
	}
}

//
//@author: chy
//@function: GormMysql
//@description: 初始化Mysql数据库
//@return: *gorm.DB

func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormMysqlConfig(m.LogMode)); err != nil {
		global.GVA_LOG.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

//
//@author: chy
//@function: GormPostgresql
//@description: 初始化postgresql数据库
//@return: *gorm.DB

func GormPostgresql() *gorm.DB {
	m := global.GVA_CONFIG.Postgresql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormPostgresqlConfig(m.LogMode)); err != nil {
		global.GVA_LOG.Error("Postgresql启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}


//
//@author: chy
//@function: GormSqlite3
//@description: 初始化sqlite3数据库
//@return: *gorm.DB

func GormSqlite3() *gorm.DB {
	m := global.GVA_CONFIG.Sqlite3
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormSqlite3Config(m.LogMode)); err != nil {
		global.GVA_LOG.Error("Sqlite3启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}


//
//@author: chy
//@function: GormOracle
//@description: 初始化oracle数据库
//@return: *gorm.DB

func GormOracle() *gorm.DB {
	m := global.GVA_CONFIG.Oracle
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormOracleConfig(m.LogMode)); err != nil {
		global.GVA_LOG.Error("Oracle启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

//
//@author: chy
//@function: GormDameng
//@description: 初始化dameng数据库
//@return: *gorm.DB

func GormDameng() *gorm.DB {
	m := global.GVA_CONFIG.Dameng
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormDamengConfig(m.LogMode)); err != nil {
		global.GVA_LOG.Error("Dameng启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}


//@author: chy
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config

func gormMysqlConfig(mod bool) *gorm.Config {
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.GVA_CONFIG.Mysql.LogZap {
	case "silent", "Silent":
		config.Logger = include.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = include.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = include.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = include.Default.LogMode(logger.Info)
	case "zap", "Zap":
		config.Logger = include.Default.LogMode(logger.Info)
	default:
		if mod {
			config.Logger = include.Default.LogMode(logger.Info)
			break
		}
		config.Logger = include.Default.LogMode(logger.Silent)
	}
	return config
}

//@author: chy
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config

func gormPostgresqlConfig(mod bool) *gorm.Config {
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.GVA_CONFIG.Postgresql.LogZap {
	case "silent", "Silent":
		config.Logger = include.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = include.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = include.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = include.Default.LogMode(logger.Info)
	case "zap", "Zap":
		config.Logger = include.Default.LogMode(logger.Info)
	default:
		if mod {
			config.Logger = include.Default.LogMode(logger.Info)
			break
		}
		config.Logger = include.Default.LogMode(logger.Silent)
	}
	return config
}

//@author: chy
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config

func gormSqlite3Config(mod bool) *gorm.Config {
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.GVA_CONFIG.Sqlite3.LogZap {
	case "silent", "Silent":
		config.Logger = include.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = include.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = include.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = include.Default.LogMode(logger.Info)
	case "zap", "Zap":
		config.Logger = include.Default.LogMode(logger.Info)
	default:
		if mod {
			config.Logger = include.Default.LogMode(logger.Info)
			break
		}
		config.Logger = include.Default.LogMode(logger.Silent)
	}
	return config
}

//@author: chy
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config

func gormOracleConfig(mod bool) *gorm.Config {
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.GVA_CONFIG.Oracle.LogZap {
	case "silent", "Silent":
		config.Logger = include.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = include.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = include.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = include.Default.LogMode(logger.Info)
	case "zap", "Zap":
		config.Logger = include.Default.LogMode(logger.Info)
	default:
		if mod {
			config.Logger = include.Default.LogMode(logger.Info)
			break
		}
		config.Logger = include.Default.LogMode(logger.Silent)
	}
	return config
}

//@author: chy
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config

func gormDamengConfig(mod bool) *gorm.Config {
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.GVA_CONFIG.Dameng.LogZap {
	case "silent", "Silent":
		config.Logger = include.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = include.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = include.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = include.Default.LogMode(logger.Info)
	case "zap", "Zap":
		config.Logger = include.Default.LogMode(logger.Info)
	default:
		if mod {
			config.Logger = include.Default.LogMode(logger.Info)
			break
		}
		config.Logger = include.Default.LogMode(logger.Silent)
	}
	return config
}