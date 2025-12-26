package media

import "time"

type Media struct {
	UUID           string    `json:"uuid"`
	WrapUUID       string    `json:"wrap_uuid"`
	Filename       string    `json:"filename"`
	FilePath       string    `json:"file_path"`
	FileSize       int64     `json:"file_size"`
	MimeType       string    `json:"mime_type"`
	UploadedAt     time.Time `json:"uploaded_at"`
	Photo_taken_at time.Time `json:"photo_taken_at"`
}
