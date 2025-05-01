package repository

import (
	"errors"
	"example/model"
	"fmt"
)

type RestaurantRepositoryInterface interface {
	GetAllRestaurant() ([]model.Restaurant, error)
	GetRestaurantById(id int) (model.Restaurant, error)
	UpsertRestaurant(restaurant model.Restaurant) (error)
	DeleteRestaurant(id int) (error)
	GetOpenRestaurants() ([]model.Restaurant, error)
}

type RestaurantRepository struct {

}
var allRestaurants = []model.Restaurant{
	model.Restaurant{ID: 1, Name: "Hot Hot Mala", Address: "Payathai Road", Cuisine: "Chinese", IsOpen: true},
	model.Restaurant{ID: 2, Name: "Sushi Bar", Address: "Sukhumvit Road", Cuisine: "Japanese", IsOpen: false},
	model.Restaurant{ID: 3, Name: "Pasta Palace", Address: "Siam Square", Cuisine: "Italian", IsOpen: true},
	model.Restaurant{ID: 4, Name: "Taco Town", Address: "Chinatown", Cuisine: "Mexican", IsOpen: false},
	model.Restaurant{ID: 5, Name: "Curry Corner", Address: "Silom Road", Cuisine: "Indian", IsOpen: true},
}

func NewRestaurantRepository() RestaurantRepository {
	return RestaurantRepository{}
}


func (repo *RestaurantRepository) GetAllRestaurant() ([]model.Restaurant, error) {
	return allRestaurants, nil
}

func (repo *RestaurantRepository) GetRestaurantById(id int) (model.Restaurant, error) {
	for _, restaurant := range allRestaurants {
		if restaurant.ID == id {
			return restaurant, nil
		}
	}
	return model.Restaurant{}, errors.New("restaurant not found")
}

func (repo *RestaurantRepository) UpsertRestaurant(restaurant model.Restaurant) (error) {

	for i, elem := range allRestaurants {
		if elem.ID == restaurant.ID {
			allRestaurants[i].Name = restaurant.Name
			allRestaurants[i].Address = restaurant.Address
			allRestaurants[i].Cuisine = restaurant.Cuisine
			allRestaurants[i].IsOpen = restaurant.IsOpen
			fmt.Println("Restaurant updated:", elem)
			return nil
		}
	}

	allRestaurants = append(allRestaurants, restaurant)
	fmt.Println("Restaurant added:", restaurant)
	return nil
}


func (repo *RestaurantRepository) DeleteRestaurant(id int) (error) {
	for i, elem := range allRestaurants {
		if elem.ID == id {
			allRestaurants = append(allRestaurants[:i], allRestaurants[i+1:]...)
			return nil
		}
	}

	return errors.New("restaurant not exists")
}

func (repo *RestaurantRepository) GetOpenRestaurants(c chan model.Restaurant) {
	for _, restaurant := range allRestaurants {
		if restaurant.IsOpen {
			c <- restaurant
		}
	}
	close(c)
	
}