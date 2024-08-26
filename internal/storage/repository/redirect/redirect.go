package redirect

import (
	"database/sql"
	"fmt"

	"github.com/tonitaga/url-shortener/internal/model/redirect"
)

type RedirectRepository struct {
	database *sql.DB
}

func New(database *sql.DB) *RedirectRepository {
	return &RedirectRepository{database: database}
}

func (r *RedirectRepository) Create(model *redirect.RedirectModel) error {
	const query = "INSERT INTO redirects (alias, target) VALUES ($1, $2)"
	const op = "redirect.RedirectRepository.Create"

	result, err := r.database.Exec(query, model.Alias, model.Target)
	if err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}

	if rows != 1 {
		return fmt.Errorf("%s: Alias exists in database, try again with another alias", op)
	}

	return nil
}

func (r *RedirectRepository) FindByAlias(alias string) (*redirect.RedirectModel, error) {
	const query = "SELECT target FROM redirects WHERE alias = $1"
	const op = "redirect.RedirectRepository.FindByAlias"

	model := &redirect.RedirectModel{
		Alias: alias,
	}

	if err := r.database.QueryRow(query, alias).Scan(&model.Target); err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	return model, nil
}

func (r *RedirectRepository) FindByTarget(target string) (*redirect.RedirectModel, error) {
	const query = "SELECT alias FROM redirects WHERE target = $1"
	const op = "redirect.RedirectRepository.FindByTarget"

	model := &redirect.RedirectModel{
		Target: target,
	}

	if err := r.database.QueryRow(query, target).Scan(&model.Alias); err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	return model, nil
}
