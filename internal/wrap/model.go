package wrap

import "time"

type Wrap struct {
	ID         int64     `json:"id"`
	UUID       string    `json:"uuid"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	Created_at time.Time `json:"create_at"`
	Updated_at time.Time `json:"update_at"`
}

const (
	StatusPending    = "pending"    // Just created
	StatusProcessing = "processing" // Images uploaded, generating
	StatusCompleted  = "completed"  // Video/collage ready
	StatusFailed     = "failed"
)
