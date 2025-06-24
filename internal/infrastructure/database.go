package infrastructure

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"todo-list/internal/domain"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() (*Database, error) {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./data/todo.db"
		log.Println("DB_PATH not set, using default:", dbPath)
	} else {
		log.Println("Using database path from environment:", dbPath)
	}

	dataDir := "./data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, err
	}

	gormLogger := logger.Default
	mode := os.Getenv("MODE")
	if mode == "debug" {
		gormLogger = logger.Default.LogMode(logger.Info)
		log.Println("Database logging set to INFO mode")
	} else {
		log.Println("Database logging set to default mode")
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&domain.User{}, &domain.Todo{}); err != nil {
		return nil, err
	}

	log.Println("Database connected successfully")
	return &Database{DB: db}, nil
}

func (d *Database) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
