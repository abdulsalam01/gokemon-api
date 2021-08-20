package api

import (
	"encoding/json"
	"fmt"
	"gokemon/abdulsalam/model"
	"math/big"
	"math/rand"
	"net/http"

	"github.com/labstack/echo"
)

func GetPokemonLocal(c echo.Context) error {
	db := DatabaseManager()
	pokemon := []model.Pokemon{}

	db.Find(&pokemon)
	// init-array of pokextra
	pokeExtra := make([]model.PokemonExtra, len(pokemon))
	for i, v := range pokemon {
		pokeExtra[i] = model.PokemonExtra{
			Pokemon:   v,
			FiboCount: FibonacyConverter(v.Count),
			NameFibo:  fmt.Sprintf("%s - %d", v.Name, pokeExtra[i].FiboCount),
		}
	}

	return c.JSON(http.StatusOK, pokeExtra)
}

func GetPokemonLocalById(c echo.Context) error {
	id := c.Param("id")
	db := DatabaseManager()
	pokemon := model.Pokemon{}

	db.First(&pokemon, id).Take(&pokemon)

	// init-array of pokextra
	pokeExtra := model.PokemonExtra{}
	pokeExtra = model.PokemonExtra{
		Pokemon:   pokemon,
		FiboCount: FibonacyConverter(pokemon.Count),
		NameFibo:  fmt.Sprintf("%s - %d", pokemon.Name, pokeExtra.FiboCount),
	}

	return c.JSON(http.StatusOK, pokeExtra)
}

func CreatePokemonLocal(c echo.Context) error {
	db := DatabaseManager()
	jsonBody := model.Pokemon{}
	json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	res := db.Create(&jsonBody)

	return c.JSON(http.StatusCreated, res)
}

func UpdatePokemonLocal(c echo.Context) error {
	id := c.Param("id")
	db := DatabaseManager()

	jsonBody := model.Pokemon{}
	respBody := model.Pokemon{}
	json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	db.First(&model.Pokemon{}, id).Take(&respBody)

	respBody.Count = respBody.Count + 1
	respBody.Name = jsonBody.Name

	db.Save(&respBody)
	return c.JSON(http.StatusOK, respBody)
}

func DeletePokemonLocal(c echo.Context) error {
	id := c.Param("id")
	db := DatabaseManager()
	res := db.Delete(&model.Pokemon{}, id)

	return c.JSON(http.StatusOK, res)
}

func GetLuckyNumber(c echo.Context) error {
	i := RandomizeNumber()
	over50 := true
	isPrime := false

	if i <= 10 {
		over50 = false
	}

	if big.NewInt(i).ProbablyPrime(0) {
		isPrime = true
	}

	return c.JSON(http.StatusOK, struct {
		Catched bool  `json:"catch"`
		IsPrime bool  `json:"is_prime"`
		Number  int64 `json:"number"`
	}{
		Catched: over50,
		IsPrime: isPrime,
		Number:  i,
	})
}

// private number return here
func RandomizeNumber() int64 {
	return rand.Int63n(20)
}

func FibonacyConverter(n int) int {
	if n < 2 {
		return n
	}

	return FibonacyConverter(n-2) + FibonacyConverter(n-1)
}
