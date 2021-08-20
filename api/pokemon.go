package api

import (
	"encoding/json"
	"fmt"
	"gokemon/abdulsalam/model"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetPokemons(c echo.Context) error {
	pokeReq := fmt.Sprintf("%s/pokemon?limit=10", BaseUrl())
	res, err := http.Get(pokeReq)

	if err != nil {
		return c.JSON(http.StatusNotAcceptable, model.BaseResponse{
			Data:    nil,
			Success: false,
			Message: "error request!",
		})
	}

	response := model.PokeAPI{}
	pokechan := make(chan model.PokeListAPI)
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &response)

	for _, v := range response.Results {
		go GetPokemonList(v.Url, pokechan)
	}

	responseArr := make([]model.PokeListAPI, len(response.Results))
	for i := 0; i < len(response.Results); i++ {
		responseArr[i] = <-pokechan
	}

	return c.JSON(http.StatusOK, model.BaseResponse{
		Data:    responseArr,
		Success: true,
		Message: "success!",
	})
}

func GetPokemon(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	pokeReqHome := fmt.Sprintf("%s/pokemon/%d", BaseUrl(), id)
	pokeReqDetail := fmt.Sprintf("%s/type/%d", BaseUrl(), id)

	res, err := http.Get(pokeReqDetail)

	if err != nil {
		return c.JSON(http.StatusNotAcceptable, model.BaseResponse{
			Data:    nil,
			Success: false,
			Message: "error request!",
		})
	}

	// get the detail of pokechan
	pokecham := make(chan model.PokeListAPI)
	go GetPokemonList(pokeReqHome, pokecham)

	response := model.PokeListTypeAPI{}
	pokechan := make(chan model.PokeListMovementAPI)
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &response)

	for _, v := range response.Moves {
		go GetPokemonMovement(v.Url, pokechan)
	}

	// get res from the channel - gouroutine
	responseArr := make([]model.PokeListMovementAPI, len(response.Moves))
	for i := 0; i < len(response.Moves); i++ {
		responseArr[i] = <-pokechan
	}

	// tight all together
	respData := model.PokeDetailAPI{
		PokeListAPI:         <-pokecham,
		PokeListTypeAPI:     response,
		PokeListMovementAPI: responseArr,
	}

	return c.JSON(http.StatusOK, model.BaseResponse{
		Data:    respData,
		Success: true,
		Message: "success!",
	})
}

// private method
func GetPokemonList(url string, ch chan<- model.PokeListAPI) {
	req, _ := http.Get(url)

	model := model.PokeListAPI{}
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &model)

	ch <- model
}

func GetPokemonMovement(url string, ch chan<- model.PokeListMovementAPI) {
	req, _ := http.Get(url)

	model := model.PokeListMovementAPI{}
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &model)

	ch <- model
}
