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
	wrap.UUID = uuid.New().String()
	query := `INSERT INTO wrap(uuid,name,status,created_at,updated_at) VALUES(?,?)`
	_, err := r.db.Exec(query, wrap.Name, wrap.Status)
	if err != nil {
		return err
	}
	return err

}
