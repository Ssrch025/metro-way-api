package domain

import "time"

type TrainSchedule struct {
	ID                 uint `gorm:"primaryKey"`
	TrainID            uint `gorm:"foreignKey:Train"`
	DepartureStationID uint `gorm:"foreignKey:Station"`
	ArrivalStationID   uint `gorm:"foreignKey:Station"`
	DepartureTime      time.Time
	ArrivalTime        time.Time
	RunsOn             []string `gorm:"type:text[]"`
}
