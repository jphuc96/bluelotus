package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Consignment is a struct for domain consignment
type Consignment struct {
	ID               string     `json:"id,omitempty"`
	VesselID         *string    `json:"vessel_id,omitempty"`
	ContainersNumber int32      `json:"containers_number,omitempty"`
	Weight           int32      `json:"weight,omitempty"`
	Description      string     `json:"description,omitempty"`
	CreatedAt        time.Time  `json:"created_at,omitempty"`
	UpdatedAt        time.Time  `json:"updated_at,omitempty"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}

// BeforeCreate prepare data before create data
func (d *Consignment) BeforeCreate(scope *gorm.Scope) error {
	d.ID = uuid.New().String()
	return nil
}
