package db

import (
	"accounts/internal/config"
	"accounts/internal/models"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*sql.DB, *gorm.DB, error) {

    dsn := config.GetConfig().DbConfig.Dsn
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        logrus.Fatalf("open db error: %v", err)
        return nil, nil, err
    }

    // Get the underlying sql.DB object to close the connection later
    sqlDB, err := db.DB()
    if err != nil {
        logrus.Errorf("Error getting database: %v", err)
        return nil, nil, err
    }

    // Ping the database to check if the connection is successful
    err = sqlDB.Ping()
    if err != nil {
        logrus.Errorf("Error pinging database: %v", err)
        return nil, nil, errors.New("ping db error")
    }

    logrus.Info("Database connection successful")

    // Perform auto migration
    err = db.AutoMigrate(models.Example{})
    if err != nil {
        logrus.Errorf("Error auto migrating database : %v", err)
        return nil, nil, err
    }
    logrus.Info("Database auto migration completed")

	return sqlDB, db, nil
}