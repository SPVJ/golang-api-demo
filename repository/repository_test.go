package repository_test

import (
	"example/model"
	"example/repository"
	"testing"
	"github.com/stretchr/testify/assert")

func TestGetAllRestaurant(t *testing.T){
	

	t.Run("Testing getting all restuarant", func (t *testing.T) {
		repo := repository.NewRestaurantRepository()
		actual,err := repo.GetAllRestaurant()

		if err != nil {
			t.Errorf("GetAllRestaurant throws an error %s", err)
		}
		
		assertion := assert.New(t)
		assertion.Equal(5, len(actual))

	})
}

var getRestaurantByIdtestCase = []struct{
	name string
	id int
	expect model.Restaurant
}{
	{name: "get restaurant id 1", id:1, expect: model.Restaurant{ID: 1, Name: "Hot Hot Mala", Address: "Payathai Road", Cuisine: "Chinese", IsOpen: true}},
	{name: "get restaurant id 2", id:2, expect: model.Restaurant{ID: 2, Name: "Sushi Bar", Address: "Sukhumvit Road", Cuisine: "Japanese", IsOpen: false}},
	{name: "get restaurant id 3", id:3, expect: model.Restaurant{ID: 3, Name: "Pasta Palace", Address: "Siam Square", Cuisine: "Italian", IsOpen: true}},
}


func TestGetRestaurantById(t *testing.T){
	for _,tc := range getRestaurantByIdtestCase{
		t.Run(tc.name, func (t *testing.T){
			repo := repository.NewRestaurantRepository()

			actual,err := repo.GetRestaurantById(tc.id)

			if err != nil {
				t.Errorf("GetAllRestaurant throws an error %s at id : %d", err, tc.id)
			}
			
			assertion := assert.New(t)
			assertion.Equal(tc.expect.Name, actual.Name)
			assertion.Equal(tc.expect.Address, actual.Address)

			
		})
	}



}