package media

import (
	"database/sql"

	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateMedia(media *Media) error {
	if media.ID == "" {
		media.ID = uuid.New().String() // You'll need to import "github.com/google/uuid"
	}
	query := `INSERT INTO media (id,wrap_uuid,filename,file_path,file_size,mime_type,uploaded_at,photo_taken_at)`
	_, err := r.db.Exec(query)
	if err != nil {
		return err

	}
	return err
}
func (r *Repository) GetMedia(uuid string) (*Media, error) {
	query := `SELECT id,wrap_uuid,filename,file_path,file_size,mime_type,uploaded_at,photo_taken_at FROM media WHERE id=?`
}
