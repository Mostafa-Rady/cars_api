package cars

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestAddCar(t *testing.T) {
	// arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	carSvcMock := NewMockCarHandler(mockCtrl)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	carsCtrl := CarController{Svc: carSvcMock}

	// Define the route similar to its definition in the routes file
	r.POST("api/v1/cars", carsCtrl.CreateCar)

	dto := CarPreview{
		Name:         "Renault Megan",
		Type:         "Sedan",
		Color:        "Red",
		SpeedRangeKm: 180,
		CarFeatures: []string{
			"Sunroof",
		},
	}

	carSvcMock.EXPECT().CreateCar(&dto).Return(uint(1), nil).Times(1)
	carSvcMock.EXPECT().ValidateCar(&dto).Return(true, nil).Times(1)

	data, err := json.Marshal(&dto)
	if err != nil {
		t.Error(err)
	}

	payload := strings.NewReader(string(data))
	req, err := http.NewRequest("POST", "/api/v1/cars", payload)
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	// act
	r.ServeHTTP(w, req)

	var tmp struct {
		Data struct {
			CarId int `json:"carId"`
		} `json:"data"`
		Msg string `json:"msg"`
	}

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(body, &tmp)
	if err != nil {
		t.Fatal(err)
	}

	// assert
	// todo: add more asserts for unsuccessful requests

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.Equal(t, tmp.Msg, "Created")
	assert.Equal(t, tmp.Data.CarId, 1)

}

func TestGetCarByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	carSvcMock := NewMockCarHandler(mockCtrl)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router and silence gin logs
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	carsCtrl := CarController{Svc: carSvcMock}

	// Define the route similar to its definition in the routes file
	r.GET("api/v1/cars/:id", carsCtrl.GetCarByID)

	now := time.Now()
	c := CarPreview{
		Id:             1,
		CreatedAt:      now,
		UpdatedAt:      now,
		Name:           "Renault Megan",
		Type:           "Sedan",
		Color:          "Red",
		SpeedRangeKm:   180,
		CarFeaturesIds: nil,
		CarFeatures: []string{
			"Sunroof",
		},
	}

	carSvcMock.EXPECT().FindCarByID(uint(1)).Return(&c, nil).Times(1)

	req, err := http.NewRequest("GET", "/api/v1/cars/1", nil)
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	r.ServeHTTP(w, req)

	var tmp struct {
		Data struct {
			Car struct {
				Id           int         `json:"id"`
				CreatedAt    time.Time   `json:"createdAt"`
				UpdatedAt    time.Time   `json:"updatedAt"`
				DeletedAt    interface{} `json:"deletedAt"`
				Name         string      `json:"name"`
				Type         string      `json:"type"`
				Color        string      `json:"color"`
				SpeedRangeKm int         `json:"speedRangeKm"`
				Features     []string    `json:"features"`
			} `json:"car"`
		} `json:"data"`
	}

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(body, &tmp)
	if err != nil {
		t.Fatal(err)
	}

	// todo: add more asserts for unsuccessful requests

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, tmp.Data.Car.Id, 1)
	assert.Equal(t, len(tmp.Data.Car.Features), 1)

	// todo: add more test cases for uncovered response errors

}

func TestSearchCtrl(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	carSvcMock := NewMockCarHandler(mockCtrl)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	carsCtrl := CarController{Svc: carSvcMock}

	// Define the route similar to its definition in the routes file
	r.POST("api/v1/cars/search", carsCtrl.Search)

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
	results := []CarPreview{
		{
			Id:             1,
			Name:           "BMW X5",
			TypeID:         3,
			Type:           "Suv",
			ColorID:        1,
			Color:          "Red",
			CarFeatures:    []string{"Sunroof"},
			CarFeaturesIds: []int32{1},
		},
	}
	carSvcMock.EXPECT().Search(&p).Return(results, nil).Times(1)
	data, err := json.Marshal(&p)
	if err != nil {
		t.Error(err)
	}

	payload := strings.NewReader(string(data))
	req, err := http.NewRequest("POST", "/api/v1/cars/search", payload)
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	r.ServeHTTP(w, req)

	var tmp struct {
		Data struct {
			Cars []struct {
				Id           int         `json:"id"`
				CreatedAt    time.Time   `json:"createdAt"`
				UpdatedAt    time.Time   `json:"updatedAt"`
				DeletedAt    interface{} `json:"deletedAt"`
				Name         string      `json:"name"`
				Type         string      `json:"type"`
				Color        string      `json:"color"`
				SpeedRangeKm int         `json:"speedRangeKm"`
				Features     []string    `json:"features"`
			} `json:"cars"`
		} `json:"data"`
	}

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(body, &tmp)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, tmp.Data.Cars[0].Id, 1)

	// todo: add more asserts for unsuccessful requests
	// todo: add more asserts for unmatched inputs
}
