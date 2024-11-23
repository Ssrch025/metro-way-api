package domain

type Train struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Type string // Express, Local
}
