package service

import (
	"example/model"
	"example/repository"
	"fmt"
)

type RestaurantService interface {
	GetAllRestaurant() (model.GetAllRestaurantResponse, error)
	GetRestaurantById(id int) (model.GetRestaurantByIdResponse, error)
	UpsertRestaurant(restaurant model.Restaurant) (model.UpsertRestaurantResponse, error)
	DeleteRestaurant(id int) (model.DeleteRestaurantResponse, error)
	GetOpenRestaurants() (model.GetOpenRestaurantsResponse, error)
}

type RestaurantServiceImpl struct {
	restaurantRepository repository.RestaurantRepository
}

func NewRestaurantServiceImpl(restaurantRepository repository.RestaurantRepository) RestaurantService {
	return &RestaurantServiceImpl{
		restaurantRepository: restaurantRepository,
	}
}
func (service RestaurantServiceImpl) GetAllRestaurant() (model.GetAllRestaurantResponse, error) {
	restaurants, err := service.restaurantRepository.GetAllRestaurant()
	if err != nil {
		fmt.Println("Error fetching restaurants:", err)
		return model.GetAllRestaurantResponse{}, err
	}

	response := model.GetAllRestaurantResponse{
		Restaurants: restaurants,
	}

	return response, nil
}
func (service RestaurantServiceImpl) GetRestaurantById(id int) (model.GetRestaurantByIdResponse, error) {
	restaurant, err := service.restaurantRepository.GetRestaurantById(id)
	if err != nil {
		fmt.Println("Error fetching restaurant by ID:", err)
		return model.GetRestaurantByIdResponse{}, err
	}

	response := model.GetRestaurantByIdResponse{
		Restaurant: restaurant,
	}
	return response, nil
}
func (service RestaurantServiceImpl) UpsertRestaurant(restaurant model.Restaurant) (model.UpsertRestaurantResponse, error) {
	err := service.restaurantRepository.UpsertRestaurant(restaurant)
	if err != nil {
		fmt.Println("Error upserting restaurant:", err)
		return model.UpsertRestaurantResponse{},err
	}
	return model.UpsertRestaurantResponse{Result: "success",}, nil
}

func (service RestaurantServiceImpl) DeleteRestaurant(id int) (model.DeleteRestaurantResponse, error) {
	err := service.restaurantRepository.DeleteRestaurant(id)
	if err != nil {
		fmt.Println("Error deleting restaurant:", err)
		return model.DeleteRestaurantResponse{}, err
	}
	return model.DeleteRestaurantResponse{Result: "Success",},nil
}


func (service RestaurantServiceImpl) GetOpenRestaurants() (model.GetOpenRestaurantsResponse, error) {

	c := make(chan model.Restaurant)
	go service.restaurantRepository.GetOpenRestaurants(c)

	restaurants := []model.Restaurant{}

	for restaurant := range c {
		restaurants = append(restaurants, restaurant)
	}

	return model.GetOpenRestaurantsResponse{Restaurants: restaurants}, nil
}