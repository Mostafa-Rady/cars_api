package cars

import (
	"fmt"
	"strings"
)

// CarHandler abstracts cars business logic (depends on CarPersistence)
//go:generate mockgen -source cars_service.go -destination cars_service_mock.go -package cars
type CarHandler interface {
	CreateCar(car *CarPreview) (uint, error)
	ValidateCar(car *CarPreview) (bool, []string)
	FindCarByID(id uint) (*CarPreview, error)
	Search(c *CarSearch) ([]CarPreview, error)
}

type CarService struct {
	Repo CarPersistence
}

// todo: replace next 3 variables with calls to DB to fetch available data instead of adding more hardcoded items in the future
var carTypes = map[string]int{
	"Sedan":      1,
	"Van":        2,
	"Suv":        3,
	"Motor-bike": 4}

var carColors = map[string]int{
	"Red":   1,
	"Green": 2,
	"Blue":  3}

var carFeatures = map[string]int{
	"Sunroof":         1,
	"Panorama":        2,
	"Auto-parking":    3,
	"Surround-system": 4}

func (d *CarPreview) toDbCarModel() (*Car, []CarsFeature, error) {
	ct, err := getItemKey(carTypes, d.Type, "car type")
	if err != nil {
		return nil, nil, err
	}

	cc, err := getItemKey(carColors, d.Color, "car color")
	if err != nil {
		return nil, nil, err
	}

	var features []CarsFeature
	for _, f := range d.CarFeatures {
		cf, err := getItemKey(carFeatures, f, "car feature")
		if err != nil {
			return nil, nil, err
		}

		features = append(features, CarsFeature{
			FeatureID: cf,
		})
	}

	c := Car{
		Name:         d.Name,
		Type:         ct,
		Color:        cc,
		SpeedRangeKm: d.SpeedRangeKm,
	}

	return &c, features, nil
}

func getItemKey(items map[string]int, item string, itemName string) (int, error) {
	if isEmptyString(item) {
		return 0, fmt.Errorf("%s is required %s", itemName, item)
	}

	if items == nil {
		return 0, fmt.Errorf("unsupported %s: %s", itemName, item)
	}
	ct, ok := items[item]
	if !ok {
		return 0, fmt.Errorf("unsupported %s: %s", itemName, item)
	}

	return ct, nil
}

func isEmptyString(item string) bool {
	return len(strings.ReplaceAll(item, " ", "")) < 1
}

// ValidateCar validates car properties and returns an array of validation errors (if any)
func (s *CarService) ValidateCar(car *CarPreview) (bool, []string) {
	var m []string

	if isEmptyString(car.Name) {
		m = append(m, "Name is required")
	}

	_, err := getItemKey(carTypes, car.Type, "car type")
	if err != nil {
		m = append(m, err.Error())
	}

	_, err = getItemKey(carColors, car.Color, "car color")
	if err != nil {
		m = append(m, err.Error())
	}

	for _, f := range car.CarFeatures {
		_, err = getItemKey(carFeatures, f, "car feature")
		if err != nil {
			m = append(m, err.Error())
		}
	}

	if car.SpeedRangeKm < 0 || car.SpeedRangeKm > 240 {
		m = append(m, "Speed range should be between 0:240")
	}

	return len(m) < 1, m
}

func (s *CarService) CreateCar(car *CarPreview) (uint, error) {
	dbCar, features, err := car.toDbCarModel()
	if err != nil {
		return 0, err
	}
	created, err := s.Repo.CreateCar(dbCar)
	if err != nil {
		return 0, err
	}

	for i := range features {
		features[i].CarID = created.ID
	}

	err = s.Repo.AddFeatures(features)
	if err != nil {
		return 0, err
	}
	return created.ID, nil

}

func (s *CarService) FindCarByID(id uint) (*CarPreview, error) {
	if id < 1 {
		return nil, nil
	}

	car, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return car, nil
}

func (s *CarService) Search(c *CarSearch) ([]CarPreview, error) {
	if c.Color != nil && len(*c.Color) > 0 {
		colorID, err := getItemKey(carColors, *c.Color, "color")
		if err != nil {
			return nil, err
		}
		c.ColorID = &colorID
	}

	if c.Type != nil && len(*c.Type) > 0 {
		typeID, err := getItemKey(carTypes, *c.Type, "type")
		if err != nil {
			return nil, err
		}
		c.TypeID = &typeID
	}

	if c.CarFeatures != nil && len(c.CarFeatures) > 0 {
		for _, f := range c.CarFeatures {
			fID, err := getItemKey(carFeatures, f, "type")
			if err != nil {
				return nil, err
			}
			c.CarFeaturesIds = append(c.CarFeaturesIds, int32(fID))
		}

	}

	results, err := s.Repo.Find(c)
	if err != nil {
		return nil, err
	}

	return results, nil
}
