package models

import "time"

// Region represents a geographic region (e.g. Teso, Central, Eastern)
type Region struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null;uniqueIndex" json:"name"`
	Code      string    `gorm:"size:50;uniqueIndex" json:"code"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`

	Zones []Zone `gorm:"foreignKey:RegionID" json:"zones,omitempty"`
}

func (Region) TableName() string {
	return "regions"
}

// Zone belongs to a Region (e.g. Zone 1, Zone 2 within a region)
type Zone struct {
	ID         uint64    `gorm:"primaryKey" json:"id"`
	RegionID   uint64    `gorm:"not null;index" json:"region_id"`
	RegionCode string    `gorm:"size:50" json:"region_code"`
	Name       string    `gorm:"size:100;not null" json:"name"`
	Code       string    `gorm:"size:50" json:"code"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at"`

	Region    Region     `gorm:"foreignKey:RegionID" json:"region,omitempty"`
	Districts []District `gorm:"foreignKey:ZoneID" json:"districts,omitempty"`
}

func (Zone) TableName() string {
	return "zones"
}

// District belongs to a Zone (and thus to a Region)
type District struct {
	ID         uint64    `gorm:"primaryKey" json:"id"`
	ZoneID     uint64    `gorm:"not null;index" json:"zone_id"`
	RegionCode string    `gorm:"size:50" json:"region_code"`
	Name       string    `gorm:"size:100;not null" json:"name"`
	Code       string    `gorm:"size:50" json:"code"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at"`

	Zone Zone `gorm:"foreignKey:ZoneID" json:"zone,omitempty"`
}

func (District) TableName() string {
	return "districts"
}

// LevelOfCare represents facility level (HC II, HC III, HC IV, Hospital)
type LevelOfCare struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null;uniqueIndex" json:"name"`
	Code      string    `gorm:"size:50;uniqueIndex" json:"code"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}

func (LevelOfCare) TableName() string {
	return "levels_of_care"
}
