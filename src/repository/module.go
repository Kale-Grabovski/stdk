package repository

import (
	"github.com/Kale-Grabovski/stdk/src/domain"
)

type ModuleRepository struct {
	db domain.Storage
}

type ModuleRepositoryInterface interface {
	GetById(id int) (*domain.Module, error)
	Get() (modules []domain.Module, err error)
	Create(name string) error
}

func NewModuleRepository(db domain.Storage) ModuleRepositoryInterface {
	return &ModuleRepository{db}
}

func (r *ModuleRepository) GetById(id int) (*domain.Module, error) {
	module := &domain.Module{}
	query := `SELECT id, name FROM module WHERE id = ?`
	err := r.db.QueryRow(query, id).Scan(&module.Id, &module.Name)
	if err != nil {
		return nil, err
	}

	return module, nil
}

func (r *ModuleRepository) Get() (modules []domain.Module, err error) {
	query := `SELECT id, name FROM module ORDER BY id`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		m := domain.Module{}
		if err := rows.Scan(&m.Id, &m.Name); err != nil {
			return nil, err
		}

		modules = append(modules, m)
	}

	err = rows.Close()
	if err != nil {
		return nil, err
	}

	return modules, nil
}

func (r *ModuleRepository) Create(name string) error {
	query := `INSERT INTO module (name) VALUES (?)`
	_, err := r.db.Exec(query, name)
	return err
}
