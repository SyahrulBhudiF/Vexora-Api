package types

import (
	"github.com/google/uuid"
	"time"
)

type Entity struct {
	UUID      uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();index" json:"uuid"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
}
