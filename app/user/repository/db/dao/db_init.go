package dao

import (
	"context"
	"fmt"
	"gin-gorm-memo/v2/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 定义对User Model的CURD操作
var _db *gorm.DB

func InitDB() {
	host := config.DbHost
	post := config.DbPort
	user := config.DbUser
	password := config.DbPassWord
	database := config.DbName
	charset := config.Charset
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, password, host, post, database, charset)
	fmt.Println("tcp连接", dsn)
	err := Database(dsn)
	if err != nil {
		fmt.Println(err)
	}
}
func Database(conn string) error {
	var ormLogger logger.Interface = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conn,
		DefaultStringSize:         256,  //string类型字段默认长度
		DisableDatetimePrecision:  true, // datetime精度 兼容mysql5.6
		DontSupportRenameColumn:   true, // 不支持重命名索引
		DontSupportRenameIndex:    true, // 不知此重命名列
		SkipInitializeWithVersion: true, //根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	_db = db

	migration()

	return err
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
