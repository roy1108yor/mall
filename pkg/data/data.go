package data

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
	"xorm.io/xorm"
)

type Data struct {
	DB *xorm.Engine
}

func NewData(conf *viper.Viper) (*Data, func(), error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&collation=utf8mb4_unicode_ci",
		conf.GetString("database.user"),
		conf.GetString("database.passwd"),
		conf.GetString("database.host"),
		conf.GetInt("database.port"),
		conf.GetString("database.db"),
	)

	db, err := xorm.NewEngine(conf.GetString("database.driver"), dsn)

	if err != nil {
		return nil, nil, err
	}

	return &Data{
			DB: db,
		}, func() {
			if err := db.Context(context.Background()).Close(); err != nil {
				log.Warnf("Failed to close database err: %s \n", err)
			}
		}, nil
}
