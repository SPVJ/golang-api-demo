package controller

import (
	"example/model"
	"example/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RestaurantController interface {
	GetAllRestaurant(ctx *gin.Context)
	GetRestaurantById(ctx *gin.Context)
	UpsertRestaurant(ctx *gin.Context)
	DeleteRestaurant(ctx *gin.Context)
	GetOpenRestaurants(ctx *gin.Context)
}

type RestaurantControllerImpl struct {
	restaurantService service.RestaurantService
}

func NewRestaurantControllerImpl(restaurantService service.RestaurantService) RestaurantController {
	return &RestaurantControllerImpl{
		restaurantService: restaurantService,
	}
}

func (controller RestaurantControllerImpl) GetAllRestaurant(ctx *gin.Context) {
	response, err := controller.restaurantService.GetAllRestaurant()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ResponseEntity{
			Status:  500,
			Message: "Error fetching restaurants",
		})
		return
	}

	responseEntity := model.ResponseEntity{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    response,
	}

	ctx.JSON(http.StatusOK, responseEntity)
}

func (controller RestaurantControllerImpl) GetOpenRestaurants(ctx *gin.Context) {
	response, err := controller.restaurantService.GetOpenRestaurants()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ResponseEntity{
			Status:  500,
			Message: "Error fetching open restaurants",
		})
		return
	}

	responseEntity := model.ResponseEntity{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    response,
	}

	ctx.JSON(http.StatusOK, responseEntity)
}

func (controller RestaurantControllerImpl) GetRestaurantById(ctx *gin.Context) {
	restaurantId := ctx.Param("id")
	id, err := strconv.Atoi(restaurantId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseEntity{
			Status:  400,
			Message: "Invalid restaurant ID",
		})
		return
	}
	response, err := controller.restaurantService.GetRestaurantById(id)
	if err != nil {	
		ctx.JSON(http.StatusInternalServerError, model.ResponseEntity{
			Status:  404,
			Message: "Error fetching restaurant",
		})
		return
	}
	responseEntity := model.ResponseEntity{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    response,	
	}
	ctx.JSON(http.StatusOK, responseEntity)
}


func (controller RestaurantControllerImpl) UpsertRestaurant(ctx *gin.Context) {
	var restaurant model.Restaurant
	if err := ctx.ShouldBindJSON(&restaurant); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseEntity{
			Status:  400,
			Message: "Invalid request",
		})
		return
	}

	response, err := controller.restaurantService.UpsertRestaurant(restaurant)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ResponseEntity{
			Status:  500,
			Message: "Error upserting restaurant",
		})
		return
	}

	responseEntity := model.ResponseEntity{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    response,
	}

	ctx.JSON(http.StatusOK, responseEntity)
}
func (controller RestaurantControllerImpl) DeleteRestaurant(ctx *gin.Context) {
	restaurantId := ctx.Param("id")
	id, err := strconv.Atoi(restaurantId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseEntity{
			Status:  400,
			Message: "Invalid restaurant ID",
		})
		return
	}

	response, err := controller.restaurantService.DeleteRestaurant(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ResponseEntity{
			Status:  500,
			Message: "Error deleting restaurant",
		})
		return
	}

	responseEntity := model.ResponseEntity{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    response,
	}

	ctx.JSON(http.StatusOK, responseEntity)
}
