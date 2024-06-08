package services

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/guilherme-luvi/go-api-fiber-crud/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestGetRandomStarWarsPeople(t *testing.T) {
	// Cria um servidor de teste para simular o retorno do client HTTP
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"name": "Luke Skywalker", "height": "172", "mass": "77", "hair_color": "blond", "skin_color": "fair", "eye_color": "blue"}`))
	}))
	defer server.Close()

	// Substitui a URL da API Star Wars pela URL do nosso servidor de teste
	oldURL := config.StarWarsAPIURL
	config.StarWarsAPIURL = server.URL
	defer func() { config.StarWarsAPIURL = oldURL }()

	// Chama a função que queremos testar
	person, err := GetRandomStarWarsPeople()

	// Verifica se a função retornou um erro
	assert.Nil(t, err)

	// Verifica se a função retornou a pessoa correta
	assert.Equal(t, "Luke Skywalker", person.Name)
	assert.Equal(t, "172", person.Height)
	assert.Equal(t, "77", person.Mass)
	assert.Equal(t, "blond", person.HairColor)
	assert.Equal(t, "fair", person.SkinColor)
	assert.Equal(t, "blue", person.EyeColor)
}

func TestGetRandomStarWarsPlanet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"name": "Tatooine", "rotation_period": "23", "orbital_period": "304", "diameter": "10465", "climate": "arid", "gravity": "1 standard", "terrain": "desert"}`))
	}))
	defer server.Close()

	oldURL := config.StarWarsAPIURL
	config.StarWarsAPIURL = server.URL
	defer func() { config.StarWarsAPIURL = oldURL }()

	planet, err := GetRandomStarWarsPlanet()

	assert.Nil(t, err)

	assert.Equal(t, "Tatooine", planet.Name)
	assert.Equal(t, "23", planet.RotationPeriod)
	assert.Equal(t, "304", planet.OrbitalPeriod)
	assert.Equal(t, "10465", planet.Diameter)
	assert.Equal(t, "arid", planet.Climate)
	assert.Equal(t, "1 standard", planet.Gravity)
	assert.Equal(t, "desert", planet.Terrain)
}

func TestGetRandomNumber(t *testing.T) {
	for i := 0; i < 100; i++ {
		numberStr := getRandomNumber()
		number, err := strconv.Atoi(numberStr)

		assert.Nil(t, err)
		assert.True(t, number >= 1 && number <= 20)
	}
}
