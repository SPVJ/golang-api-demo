package main

import ("example/controller"
	"example/repository"
	"example/service"
	"example/router"
	"time"
	"net/http")


func main() {
	repository := repository.NewRestaurantRepository()
	restaurantService := service.NewRestaurantServiceImpl(repository)
	restaurantController := controller.NewRestaurantControllerImpl(restaurantService)
	r := router.NewRouter(restaurantController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()


}
