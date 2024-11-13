package schemas

import "time"

type Booking struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"column:"`
	User      User      `gorm:"foreignKey:UserID"`
	Date      time.Time `gorm:"column:date"`
	Time      time.Time `gorm:"column:time"`
	Status    string    `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"index;column:deleted_at"`
}
