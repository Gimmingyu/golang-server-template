package config

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"os/user"
	"time"
)

func NewGormClient() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	database, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Panicf("Panic occured while opening database : %v\n", err)
	}
	client, err := gorm.Open(postgres.New(postgres.Config{
		Conn:                 database,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Panicf("Panic occured while opening postgresql client : %v\n", err)
	}

	if err := client.Migrator().AutoMigrate(&user.User{}); err != nil {
		log.Println(err)
	}

	return client
}
