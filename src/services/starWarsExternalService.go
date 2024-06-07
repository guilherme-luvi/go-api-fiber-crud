package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/guilherme-luvi/go-api-fiber-crud/src/config"
)

func GetRandomStarWarsPeople() ([]byte, error) {
	url := config.StarWarsAPIURL + "/people/"

	resp, err := config.GetHttpClient().Get(url + getRandomNumber())
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetRandomStarWarsPlanet() ([]byte, error) {
	url := config.StarWarsAPIURL + "/planets/"

	resp, err := config.GetHttpClient().Get(url + getRandomNumber())
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getRandomNumber() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	number := r.Intn(20) + 1
	return fmt.Sprint(number)
}
