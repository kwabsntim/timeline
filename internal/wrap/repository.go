package wrap

import (
	"database/sql"

	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewSqlRepository(db *sql.DB) wrapinterface {
	return &Repository{db: db}
}

// create wrap
func (r *Repository) CreateWrap(wrap *Wrap) error {
	if wrap.UUID == "" {
		wrap.UUID = uuid.New().String()
	}

	query := `INSERT INTO wrap(uuid,name,status,created_at,updated_at) VALUES(?,?,?,?,?)`
	_, err := r.db.Exec(query, wrap.UUID, wrap.Name, wrap.Status, wrap.Created_at, wrap.Updated_at)
	if err != nil {
		return err
	}
	return err

}
func (r *Repository) Getwrap(uuid string) (*Wrap, error) {
	query := `SELECT uuid,name,status,created_at,updated_at FROM wrap WHERE uuid=?`
	var wrap Wrap
	err := r.db.QueryRow(query, uuid).Scan(
		&wrap.UUID,
		&wrap.Name,
		&wrap.Status,
		&wrap.Created_at,
		&wrap.Updated_at,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No wrap found
		}
		return nil, err // Database error
	}
	return &wrap, nil
}
