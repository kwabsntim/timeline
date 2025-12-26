package media

import "time"

type Media struct {
	ID         string    `json:"id"`
	WrapUUID   string    `json:"wrap_uuid"`
	Filename   string    `json:"filename"`
	FilePath   string    `json:"file_path"`
	FileSize   int64     `json:"file_size"`
	MimeType   string    `json:"mime_type"`
	UploadedAt time.Time `json:"uploaded_at"`
}
