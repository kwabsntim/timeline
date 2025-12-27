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

	//creating a limit for the files uploaded
	var count int
	countQuery := `SELECT COUNT(*)FROM media WHERE wrap_uuid=?`
	err := r.db.QueryRow(countQuery, media.WrapUUID).Scan(&count)
	if err != nil {
		return err
	}
	//if the file limit is greater than 9 stop
	if count >= 9 {
		return errors.New("The maxmimum file limit is 9 per wrap")
	}
	//if the uuid is empty generate one for it
	if media.UUID == "" {
		media.UUID = uuid.New().String() // You'll need to import "github.com/google/uuid"
	}
	query := `INSERT INTO media (id,wrap_uuid,filename,file_path,file_size,mime_type,uploaded_at,photo_taken_at)VALUES(?,?,?,?,?,?,?,?)`
	_, err = r.db.Exec(query, media.UUID, media.WrapUUID, media.Filename, media.FilePath, media.FileSize, media.MimeType, media.UploadedAt, media.Photo_taken_at)
	if err != nil {
		return err
	}
	return err
}

// get the media and its files
func (r *Repository) GetMediaWrap(WrapUUID string) ([]*Media, error) {
	//get the images based on the limit set
	query := `SELECT uuid,wrap_uuid,filename,file_path,file_size,mime_type,uploaded_at,photo_taken_at FROM media WHERE uuid=?ORDER BY uploaded_at DESC LIMIT 9`
	rows,err:=r.db.Query(query,WrapUUID)
	if err!=nil{
		return nil,err
	}
	defer rows.Close()
	var mediaList []*Media
	for rows.Next(){
		var media Media
		err := r.db.QueryRow(query, WrapUUID).Scan(
			&media.UUID,
			&media.WrapUUID,
			&media.Filename,
			&media.FilePath,
			&media.MimeType,
			&media.FileSize,
			&media.UploadedAt,
			&media.Photo_taken_at,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return []*Media
}
