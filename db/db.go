package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

type Database struct {
	*gorm.DB
}

func New(config *Config) (*Database, error) {
	db, err := gorm.Open("postgres", config.DatabaseURI)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}
	return &Database{db}, nil
}
