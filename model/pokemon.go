package model

import (
	"github.com/jinzhu/gorm"
)

type PokemonExtra struct {
	Pokemon   `json:"pokemon"`
	FiboCount int    `json:"fibo_count"`
	NameFibo  string `json:"name_fibo"`
}

type Pokemon struct {
	gorm.Model
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Count   int    `json:"count" gorm:"default:-1"`
}

type PokeAPI struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

type PokeListAPI struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Sprites struct {
		Back_default  string `json:"back_default"`
		Front_default string `json:"front_default"`
	} `json:"sprites"`
}

type PokeListTypeAPI struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Game_indices []struct {
		Game_index int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	Pokemon []struct {
		Slot    int `json:"slot"`
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon"`
	Moves []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"moves"`
}

type PokeListMovementAPI struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Accuracy       int    `json:"accuracy"`
	Effect_chance  string `json:"effect_chance"`
	Pp             int    `json:"pp"`
	Priority       int    `json:"priority"`
	Power          int    `json:"power"`
	Effect_entries []struct {
		Effect       string `json:"effect"`
		Short_effect string `json:"short_effect"`
	} `json:"effect_entries"`
}

type PokeDetailAPI struct {
	PokeListAPI         PokeListAPI           `json:"poke_list"`
	PokeListTypeAPI     PokeListTypeAPI       `json:"poke_list_type"`
	PokeListMovementAPI []PokeListMovementAPI `json:"poke_list_movement"`
}
