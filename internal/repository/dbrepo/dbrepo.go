package dbrepo

import (
	"database/sql"

	"github.com/lkk223303/bookings/internal/config"
	"github.com/lkk223303/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

func NewTestingRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: a,
	}
}

// For other DB just create a repo and and NewMariaDBRepo function
// type mariaDBRepo struct {
// 	App *config.AppConfig
// 	DB  *sql.DB
// }
