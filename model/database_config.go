package model

import "time"

// DatabaseConfig represents the structure of a database configuration
type DatabaseConfig struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null;default:''" json:"name"`
	Host      string    `gorm:"type:varchar(255);not null;default:''" json:"host"`
	Port      int       `gorm:"type:int;not null;default:0" json:"port"`
	Username  string    `gorm:"type:varchar(255);not null;default:''" json:"username"`
	Password  string    `gorm:"type:varchar(255);not null;default:''" json:"password"`
	Database  string    `gorm:"type:varchar(255)" json:"database"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (obj *DatabaseConfig) TableName() string {
	return "database_config"
}
