package url

import "github.com/jmoiron/sqlx"

type UrlRepoSQL struct {
	db *sqlx.DB
}

func NewRepoSQL(db *sqlx.DB) *UrlRepoSQL {
	return &UrlRepoSQL{db: db}
}
