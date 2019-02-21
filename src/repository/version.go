package repository

import (
	"github.com/Kale-Grabovski/stdk/src/domain"
)

type VersionRepository struct {
	db domain.Storage
}

type VersionRepositoryInterface interface {
	GetById(id int) (*domain.Version, error)
	SetActive(id int) error
	Get(moduleId int) (versions []domain.Version, err error)
	Create(*domain.Version) error
	GetActiveVersions() (versions map[string]domain.Version, err error)
	GetActiveVersionByModule(moduleName string) (*domain.Version, error)
}

func NewVersionRepository(db domain.Storage) VersionRepositoryInterface {
	return &VersionRepository{db}
}

func (r *VersionRepository) GetById(id int) (*domain.Version, error) {
	v := &domain.Version{}
	query := `SELECT id, is_active, module_id FROM module_version WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&v.Id, &v.IsActive, &v.ModuleId)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (r *VersionRepository) Get(moduleId int) (versions []domain.Version, err error) {
	query := `
		SELECT id, is_active, settings, created_at
		FROM module_version
		WHERE module_id = $1
		ORDER BY id DESC
	`
	rows, err := r.db.Query(query, moduleId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		v := domain.Version{}
		if err := rows.Scan(&v.Id, &v.IsActive, &v.Settings, &v.CreatedAt); err != nil {
			return nil, err
		}

		versions = append(versions, v)
	}

	err = rows.Close()
	if err != nil {
		return nil, err
	}

	return versions, nil
}

func (r *VersionRepository) Create(v *domain.Version) error {
	query := `INSERT INTO module_version (module_id, filename, filehash, settings) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, v.ModuleId, v.Filename, v.Hash, v.Settings)
	return err
}

func (r *VersionRepository) SetActive(id int) error {
	version, err := r.GetById(id)
	if err != nil {
		return err
	}

	query := `
		UPDATE module_version
		SET is_active = (CASE WHEN id = $1 THEN true ELSE false END)
		WHERE module_id = $2
	`
	_, err = r.db.Exec(query, id, version.ModuleId)
	return err
}

func (r *VersionRepository) GetActiveVersions() (map[string]domain.Version, error) {
	versions := make(map[string]domain.Version)

	query := `
		SELECT v.id, m.name, v.settings
		FROM module_version AS v
		JOIN module AS m ON m.id = v.module_id
		WHERE v.is_active = true
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		v := domain.Version{}
		var moduleName string
		if err := rows.Scan(&v.Id, &moduleName, &v.Settings); err != nil {
			return nil, err
		}

		versions[moduleName] = v
	}

	err = rows.Close()
	if err != nil {
		return nil, err
	}

	return versions, nil
}

func (r *VersionRepository) GetActiveVersionByModule(moduleName string) (*domain.Version, error) {
	v := &domain.Version{}
	query := `
		SELECT v.filehash, v.filename
		FROM module_version AS v
		JOIN module AS m ON v.module_id = m.id
		WHERE m.name = $1 AND v.is_active = true
	`
	err := r.db.QueryRow(query, moduleName).Scan(&v.Hash, &v.Filename)
	if err != nil {
		return nil, err
	}

	return v, nil
}
