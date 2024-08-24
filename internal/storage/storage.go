package storage

import (
	"database/sql"
	"fmt"

	"github.com/tonitaga/url-shortener/internal/config"
	"github.com/tonitaga/url-shortener/internal/storage/repository/redirect"

	_ "github.com/lib/pq"
)

type Storage struct {
	config              *config.DatabaseConfig
	database            *sql.DB
	redirect_repository *redirect.RedirectRepository
}

func New(config *config.DatabaseConfig) *Storage {
	return &Storage{
		config: config,
	}
}

func (s *Storage) Connect() error {
	const op = "storage.Storage.Connect"

	connection_string := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		s.config.Host,
		s.config.Port,
		s.config.Username,
		s.config.Password,
		s.config.Name,
	)

	db, err := sql.Open("postgres", connection_string)
	if err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}

	s.database = db

	return nil
}

func (s *Storage) Disconnect() {
	s.database.Close()
}

func (s *Storage) Redirect() *redirect.RedirectRepository {
	if s.redirect_repository == nil {
		s.redirect_repository = redirect.New(s.database)
	}

	return s.redirect_repository
}
