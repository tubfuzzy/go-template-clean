package db

import (
	"go-template-clean/config"
)

type DB struct {
}

func NewDB(conf *config.Configuration) (*DB, error) {

	return &DB{}, nil
}
