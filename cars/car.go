package cars

import "gorm.io/gorm"

// Car holds car specifications
type Car struct {
	gorm.Model
	Name         string
	Type         int
	Color        int
	SpeedRangeKm int
}

// CarsFeature holds car features-set
type CarsFeature struct {
	CarID     uint
	FeatureID int
}
