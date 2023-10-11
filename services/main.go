package services

import "paytabs-task/db"

type Services struct {
	db *db.Db
}

func NewService(db *db.Db) *Services {
	return &Services{
		db: db,
	}
}
