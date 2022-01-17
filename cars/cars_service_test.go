package cars

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestCreateCar(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	persistenceMock := NewMockCarPersistence(mockCtrl)

	svc := CarService{Repo: persistenceMock}

	carDto := CarPreview{
		Name:         "BMW X3",
		Type:         "Sedan",
		Color:        "Red",
		SpeedRangeKm: 200,
		CarFeatures: []string{
			"Sunroof",
		},
	}

	car := Car{
		Model:        gorm.Model{},
		Name:         "BMW X3",
		Type:         1,
		Color:        1,
		SpeedRangeKm: 200,
	}

	carRet := Car{
		Model: gorm.Model{
			ID: 1,
		},
		Name:         "BMW X3",
		Type:         1,
		Color:        1,
		SpeedRangeKm: 200,
	}
	persistenceMock.EXPECT().CreateCar(&car).Return(&carRet, nil).Times(1)

	features := []CarsFeature{
		{
			CarID:     1,
			FeatureID: 1,
		},
	}
	persistenceMock.EXPECT().AddFeatures(features).Return(nil).Times(1)

	carID, err := svc.CreateCar(&carDto)
	assert.Nil(t, err)
	assert.Equal(t, carID, uint(1))

	//todo: add asserts for other uncovered error checks
}

func TestFindCarByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	persistenceMock := NewMockCarPersistence(mockCtrl)

	svc := CarService{Repo: persistenceMock}
	now := time.Now()
	p := CarPreview{
		Id:           1,
		CreatedAt:    now,
		UpdatedAt:    now,
		Name:         "Tipo",
		Type:         "Sedan",
		Color:        "Red",
		SpeedRangeKm: 160,
		CarFeatures:  []string{"Sunroof"},
	}

	persistenceMock.EXPECT().FindByID(uint(1)).Return(&p, nil).Times(1)

	car, err := svc.FindCarByID(uint(1))
	assert.Nil(t, err)
	assert.Equal(t, car.Id, uint(1))
	assert.NotEqual(t, len(car.CarFeatures), 0)

	//todo: add more asserts
}

func TestSearch(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	persistenceMock := NewMockCarPersistence(mockCtrl)

	svc := CarService{Repo: persistenceMock}

	since := "2022-01-17"
	before := "2022-01-18"
	name := "BMW X5"
	color := "Red"
	tp := "Suv"
	p := CarSearch{
		Id:          nil,
		Since:       &since,
		Before:      &before,
		Name:        &name,
		Type:        &tp,
		Color:       &color,
		CarFeatures: []string{"Sunroof"},
	}
	now := time.Now()
	results := []CarPreview{
		{
			Id:             1,
			CreatedAt:      now,
			UpdatedAt:      now,
			Name:           "BMW X5",
			TypeID:         3,
			Type:           "Suv",
			ColorID:        1,
			Color:          "Red",
			CarFeatures:    []string{"Sunroof"},
			CarFeaturesIds: []int32{1},
		},
	}

	persistenceMock.EXPECT().Find(&p).Return(results, nil).Times(1)

	found, err := svc.Search(&p)
	assert.Nil(t, err)
	assert.Equal(t, found[0].Id, uint(1))
	assert.NotEqual(t, len(found), 0)

	//todo: add more asserts and implement searching for remaining car properties
}

func TestValidateCar(t *testing.T) {
	carSvc := CarService{}
	v, m := carSvc.ValidateCar(&CarPreview{
		SpeedRangeKm: 460,
		CarFeatures: []string{
			"Is wrong feature",
		},
	})
	assert.Equal(t, v, false)
	assert.NotNil(t, m)
	assert.NotEqual(t, len(m), 0)
	assert.Equal(t, m[0], "Name is required")

	//todo: add asserts for other uncovered validity messages
}
