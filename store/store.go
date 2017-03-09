package store

import (
	"github.com/Dataman-Cloud/pressure-test-app/model"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Store struct {
	*gorm.DB
}

func New(driver, dsn string) *Store {
	db, err := gorm.Open(driver, dsn)
	if err != nil {
		log.Errorf("connect to driver %s and dsn %s failed", driver, dsn)
		panic("database connection failed")
	}

	db.AutoMigrate(&model.App{})
	db.LogMode(false)

	return &Store{db}
}

func (db *Store) SaveApp(a *model.App) error {
	return db.Save(a).Error
}

func (db *Store) GetApp(id int) (*model.App, error) {
	var app model.App
	err := db.First(&app, id).Error
	if err != nil {
		return nil, err
	}
	return &app, nil
}

func (db *Store) DeleteApp(a *model.App) error {
	return db.Delete(a).Error
}

func (db *Store) UpdateApp(a *model.App) error {
	return db.Save(a).Error
}
