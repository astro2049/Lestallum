package startup

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"osaka/global"
	"osaka/model"
)

func InitMySQL() error {
	var err error
	dsn := "root:XRuul1203ds9090-iif2021@tcp(127.0.0.1:3306)/osaka?charset=utf8mb4&parseTime=True&loc=Local"
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return errors.New("failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")

	// see: [https://gorm.io/docs/migration.html]
	err = global.DB.AutoMigrate(
		// resource file
		&model.ResourceFile{},
	)

	return err
}
