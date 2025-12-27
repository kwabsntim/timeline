package media

import (
	"database/sql"
	"errors"

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

// creation of media
func (r *Repository) CreateMedia(media *Media) error {
	if media.UUID == "" {
		media.UUID = uuid.NewString()
	}
	query := `INSERT INTO media (uuid,wrap_uuid,filename,file_path,file_size,mime_type,uploaded_at,photo_taken_at) VALUES (?,?,?,?,?,?,?,?)`
	_, err := r.db.Exec(query, media.UUID, media.WrapUUID, media.Filename, media.FilePath, media.FileSize, media.MimeType, media.UploadedAt, media.Photo_taken_at)
	if err != nil {
		return err
	}
	return err
}

// get the media and its files
// get the media and its files
func (r *Repository) GetMediaByWrap(wrapUUID string) ([]*Media, error) {
	query := `SELECT uuid,wrap_uuid,filename,file_path,file_size,mime_type,uploaded_at,photo_taken_at FROM media WHERE wrap_uuid=? ORDER BY photo_taken_at ASC`

	rows, err := r.db.Query(query, wrapUUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mediaList []*Media
	for rows.Next() {
		var media Media
		err := rows.Scan(
			&media.UUID,
			&media.WrapUUID,
			&media.Filename,
			&media.FilePath,
			&media.FileSize,
			&media.MimeType,
			&media.UploadedAt,
			&media.Photo_taken_at,
		)
		if err != nil {
			return nil, err
		}
		mediaList = append(mediaList, &media)
	}
	return mediaList, nil
}

// delete the media
func (r *Repository) DeleteMedia(mediaUUID string) error {
	// Check if media exists
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM media WHERE uuid=?)`
	err := r.db.QueryRow(checkQuery, mediaUUID).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("media not found")
	}

	// Delete the media
	query := `DELETE FROM media WHERE uuid=?`
	_, err = r.db.Exec(query, mediaUUID)
	return err
}
