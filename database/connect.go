package database

import (
	"gRMS/modals"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(level logger.LogLevel) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  level,       // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(sqlite.Open("database/chatdata.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalln("failed to connected to database: ", err)
	}

	err = db.AutoMigrate(&modals.User{}, &modals.Chat{}, &modals.Message{}, &modals.Participant{},
		&modals.Admin{}, &modals.Photo{},
		&modals.Sticker{}, &modals.Video{}, &modals.Audio{}, &modals.Document{})
	if err != nil {
		log.Fatalln("failed to migrate database: ", err)
	}

	return db
}