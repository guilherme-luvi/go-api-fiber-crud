package services

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/guilherme-luvi/go-api-fiber-crud/src/config"
	"github.com/guilherme-luvi/go-api-fiber-crud/src/schemas"
)

func GetRandomStarWarsPeople() (schemas.StarWarsPerson, error) {
	url := config.StarWarsAPIURL + "/people/"

	resp, err := config.GetHttpClient().Get(url + getRandomNumber())
	if err != nil {
		return schemas.StarWarsPerson{}, err
	}

	person := schemas.StarWarsPerson{}
	if err := json.Unmarshal(resp, &person); err != nil {
		return schemas.StarWarsPerson{}, err
	}

	return person, nil
}

func GetRandomStarWarsPlanet() (schemas.StarWarsPlanet, error) {
	url := config.StarWarsAPIURL + "/planets/"

	resp, err := config.GetHttpClient().Get(url + getRandomNumber())
	if err != nil {
		return schemas.StarWarsPlanet{}, err
	}

	planet := schemas.StarWarsPlanet{}
	if err := json.Unmarshal(resp, &planet); err != nil {
		return schemas.StarWarsPlanet{}, err
	}

	return planet, nil
}

func getRandomNumber() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	number := r.Intn(20) + 1
	return fmt.Sprint(number)
}
