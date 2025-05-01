package router

import (
	"github.com/gin-gonic/gin"
	"example/controller"
   )


func NewRouter(restaurant controller.RestaurantController) *gin.Engine {

	service := gin.Default()

	router := service.Group("/restaurant")
	router.GET("/", restaurant.GetAllRestaurant)
	router.GET("/open", restaurant.GetOpenRestaurants)
	router.POST("/", restaurant.UpsertRestaurant)
	router.GET("/:id", restaurant.GetRestaurantById)
	router.DELETE("/:id", restaurant.DeleteRestaurant)

	return service

} 