package cars

import (
	"fmt"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"strings"
	"time"
)

// CarPersistence abstracts db persistence
//go:generate mockgen -source cars_persistence.go -destination cars_persistence_mock.go -package cars
type CarPersistence interface {
	// CreateCar adds new car
	CreateCar(car *Car) (*Car, error)
	// AddFeatures adds supplied car features
	AddFeatures(car []CarsFeature) error
	// FindByID finds car by id
	FindByID(id uint) (*CarPreview, error)
	// Find cars matching search inputs
	Find(c *CarSearch) (results []CarPreview, err error)
}

// CarRepo provides postgres implementation for car persistence
type CarRepo struct {
	DB *gorm.DB
}

// CarPreview dto model for adding new cars
type CarPreview struct {
	Id             uint           `json:"id"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      *time.Time     `json:"deletedAt"`
	Name           string         `json:"name"`
	TypeID         int            `json:"-"`
	Type           string         `json:"type"`
	ColorID        int            `json:"-"`
	Color          string         `json:"color"`
	SpeedRangeKm   int            `json:"speedRangeKm"`
	CarFeaturesIds pq.Int32Array  `json:"-" gorm:"type:integer[]"`
	CarFeatures    pq.StringArray `json:"features" gorm:"type:type:text[]"`
}

// CarSearch to receive search options
type CarSearch struct {
	Id             *uint          `json:"id"`
	Since          *string        `json:"since"`
	Before         *string        `json:"before"`
	Name           *string        `json:"name"`
	Type           *string        `json:"type"`
	TypeID         *int           `json:"-"`
	Color          *string        `json:"color"`
	ColorID        *int           `json:"-"`
	SpeedRangeKm   *int           `json:"speedRangeKm"`
	CarFeaturesIds pq.Int32Array  `json:"-" gorm:"type:integer[]"`
	CarFeatures    pq.StringArray `json:"features" gorm:"type:type:text[]"`
}

func (CarPreview) TableName() string {
	return "car_previews"
}

// CreateCar adds new car
func (cr *CarRepo) CreateCar(car *Car) (*Car, error) {
	err := cr.DB.Create(car).Error
	if err != nil {
		return nil, err
	}
	return car, nil
}

// AddFeatures adds supplied car features
func (cr *CarRepo) AddFeatures(features []CarsFeature) error {
	// todo: replace with a sql statement for bulk insert (performance) or use gorm bulk insert
	for _, feature := range features {
		err := cr.DB.Create(&feature).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (cr *CarRepo) FindByID(id uint) (*CarPreview, error) {
	var c CarPreview
	err := cr.DB.First(&c, id).Error
	if err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}

func (cr *CarRepo) Find(c *CarSearch) (results []CarPreview, err error) {
	if c == nil {
		return
	}

	var condList []string

	if c.Name != nil && len(*c.Name) > 0 {
		condList = append(condList,
			fmt.Sprintf("lower(name) like '%s%s%s'", "%", strings.ToLower(*c.Name), "%"))
	}
	layoutISO := "2006-01-02"

	if c.Since != nil && len(*c.Since) > 0 {
		_, err = time.Parse(layoutISO, *c.Since)
		if err != nil {
			return nil, fmt.Errorf("invalid input 'since', required format %s", layoutISO)
		}
		condList = append(condList, fmt.Sprintf("created_at >= '%s'", *c.Since))
	}

	if c.Before != nil && len(*c.Before) > 0 {
		_, err = time.Parse(layoutISO, *c.Before)
		if err != nil {
			return nil, fmt.Errorf("invalid input 'before', required format %s", layoutISO)
		}
		condList = append(condList, fmt.Sprintf("created_at <= '%s'", *c.Before))
	}

	if c.TypeID != nil {
		condList = append(condList, fmt.Sprintf("type_id = %v", *c.TypeID))
	}

	if c.ColorID != nil {
		condList = append(condList, fmt.Sprintf("color_id = %v", *c.ColorID))
	}

	if c.CarFeaturesIds != nil && len(c.CarFeaturesIds) > 0 {
		for _, f := range c.CarFeaturesIds {
			condList = append(condList, fmt.Sprintf("%v = ANY (car_features_ids)", f))
		}
	}

	q := ""
	if len(condList) < 1 {
		return
	}

	if len(condList) == 1 {
		q = condList[0]
	}

	if len(condList) > 1 {
		q = strings.Join(condList, " and ")
	}

	err = cr.DB.Find(&results, q).Error
	if err != nil {
		return nil, err
	}

	return
}
