package entity

import "time"

type Place struct {
	PlaceId          string     `gorm:"primaryKey;not_null" json:"place_id"`
	Name             string    `gorm:"type:varchar(200);" json:"place_name"`
	FormattedAddress string    `gorm:"type:varchar(200);" json:"formatted_address"`
	CreatedAt        time.Time `gorm:""DEFAULT:current_timestamp; type:timestamp"" json:"created_at"`
	UpdatedAt        time.Time `gorm:""DEFAULT:current_timestamp; type:timestamp"" json:"updated_at"`
}
