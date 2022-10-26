package storage

import (
	"email-service/storage/postgres"
	"email-service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	email() repo.EmailStorageI
}

type storagePg struct {
	db        *sqlx.DB
	emailRepo repo.EmailStorageI
}

func NewStorage(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:        db,
		emailRepo: postgres.NewEmailRepo(db),
	}
}

func (s storagePg) Customer() repo.EmailStorageI {
	return s.emailRepo
}
