package cars

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CarController contain handlers for cars end-points
type CarController struct {
	Svc CarHandler
}

// BaseResponse wraps success/failure responses
// todo: when app grows move "BaseResponse" struct to a "common" package for reuse
type BaseResponse struct {
	Data    gin.H       `json:"data,omitempty"`
	Message string      `json:"msg,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// CreateCar validates and registers a new car
func (ctrl *CarController) CreateCar(c *gin.Context) {
	var (
		car CarPreview
		res BaseResponse
	)

	err := c.BindJSON(&car)
	if err != nil {
		res.Message = "Posted data is invalid"
		res.Error = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	v, m := ctrl.Svc.ValidateCar(&car)
	if !v {
		res.Message = "Posted data is invalid"
		res.Error = m
		c.JSON(http.StatusBadRequest, res)
	}

	carID, err := ctrl.Svc.CreateCar(&car)
	if err != nil {
		res.Message = "Error creating car"
		res.Error = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res.Message = "Created"
	res.Data = gin.H{"carId": carID}
	c.JSON(http.StatusCreated, res)
}

// GetCarByID finds specific car by its id
func (ctrl *CarController) GetCarByID(c *gin.Context) {
	var res BaseResponse
	id := c.Param("id")

	if isEmptyString(id) {
		res.Message = "Item not found"
		c.JSON(http.StatusNotFound, res)
		return
	}

	carID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		res.Message = "Invalid id"
		res.Error = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if carID < 1 {
		res.Message = "Item not found"
		c.JSON(http.StatusNotFound, res)
		return
	}

	car, err := ctrl.Svc.FindCarByID(uint(carID))
	if err != nil {
		res.Message = "Error finding item"
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if car == nil {
		res.Message = "Item not found"
		c.JSON(http.StatusNotFound, res)
		return
	}
	res.Data = gin.H{"car": car}
	c.JSON(http.StatusOK, res)
}

// Search cars matching given inputs
func (ctrl *CarController) Search(c *gin.Context) {
	var (
		car CarSearch
		res BaseResponse
	)

	err := c.BindJSON(&car)
	if err != nil {
		res.Message = "Posted data is invalid"
		res.Error = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	results, err := ctrl.Svc.Search(&car)
	if err != nil {
		res.Message = "Error finding item"
		res.Error = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res.Data = gin.H{"cars": results}
	c.JSON(http.StatusOK, res)
}
