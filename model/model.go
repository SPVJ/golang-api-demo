package model

type Restaurant struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
	Cuisine string `json:"cuisine"`
	IsOpen bool   `json:"is_open"`}

type GetRestaurantByIdRequest struct {
	id int `validate:"required" json:"id"`
}

type UpsertRestaurantRequest struct {
	ID 	int    `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Cuisine string `json:"cuisine" validate:"required"`
	IsOpen bool   `json:"is_open" validate:"required"`
}

type DeleteRestaurantRequest struct {
	ID int `json:"id" validate:"required"`
}

type GetAllRestaurantResponse struct {
	Restaurants []Restaurant `json:"restaurants"`
}

type GetOpenRestaurantsResponse struct {
	Restaurants []Restaurant `json:"restaurants"`
}

type GetRestaurantByIdResponse struct {
	Restaurant Restaurant `json:"restaurant"`
}

type UpsertRestaurantResponse struct {
	Result  string `json:"result"`
}

type DeleteRestaurantResponse struct {
	Result  string `json:"result"`
}

type ResponseEntity struct {
	Status int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}