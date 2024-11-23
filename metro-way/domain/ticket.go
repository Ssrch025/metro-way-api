package domain

type Ticket struct {
	ID         uint   `gorm:"primaryKey"`
	ScheduleID uint   `gorm:"foreignKey:TrainSchedule"`
	Class      string // 1A, 2A, SL
	Price      float64
	Status     string // Available, Booked, Waitlisted
}
