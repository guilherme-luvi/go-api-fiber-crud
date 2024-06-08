package services

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/guilherme-luvi/go-api-fiber-crud/internal/models"
	"github.com/guilherme-luvi/go-api-fiber-crud/pkg/config"
)

func GetRandomStarWarsPeople() (models.StarWarsPerson, error) {
	url := config.StarWarsAPIURL + "/people/"

	resp, err := config.GetHttpClient().Get(url + getRandomNumber())
	if err != nil {
		return models.StarWarsPerson{}, err
	}

	person := models.StarWarsPerson{}
	if err := json.Unmarshal(resp, &person); err != nil {
		return models.StarWarsPerson{}, err
	}

	return person, nil
}

func GetRandomStarWarsPlanet() (models.StarWarsPlanet, error) {
	url := config.StarWarsAPIURL + "/planets/"

	resp, err := config.GetHttpClient().Get(url + getRandomNumber())
	if err != nil {
		return models.StarWarsPlanet{}, err
	}

	planet := models.StarWarsPlanet{}
	if err := json.Unmarshal(resp, &planet); err != nil {
		return models.StarWarsPlanet{}, err
	}

	return planet, nil
}

func getRandomNumber() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	number := r.Intn(20) + 1
	return fmt.Sprint(number)
}
