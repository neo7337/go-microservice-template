package postgres

import (
	"database/sql"
	"errors"
)

type PostgresRepo struct {
	Database *sql.DB
}

func NewPostgresRepo() *PostgresRepo {
	return &PostgresRepo{}
}

func (p *PostgresRepo) Connect(dsn string) (*PostgresRepo, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.InfoF("failed to connect to the database: %v", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		err = errors.New("Failed to ping Postgres: " + err.Error())
		return nil, err
	}
	p.Database = db
	return p, nil
}

func (p *PostgresRepo) Close() error {
	if p.Database != nil {
		err := p.Database.Close()
		if err != nil {
			return errors.New("Failed to close database connection: " + err.Error())
		}
		p.Database = nil
	}
	return nil
}
