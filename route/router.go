package route

import (
	"gokemon/abdulsalam/api"

	"github.com/labstack/echo"
)

type Server struct {
	e *echo.Echo
}

func (s *Server) Init() {
	apis := s.e.Group("/api")

	// entry-api-point
	apis.GET("/", api.Home)

	// pokemon-world-api
	pokegroup := apis.Group("/pokemon")
	pokegroup.GET("/list", api.GetPokemons)
	pokegroup.GET("/list/:id", api.GetPokemon)

	// pokemon-local-world-api
	pokegrouplocal := apis.Group("/pokemon-local")
	pokegrouplocal.GET("/list", api.GetPokemonLocal)
	pokegrouplocal.GET("/list/:id", api.GetPokemonLocalById)
	pokegrouplocal.GET("/lucky", api.GetLuckyNumber)
	pokegrouplocal.POST("/create", api.CreatePokemonLocal)
	pokegrouplocal.PATCH("/update/:id", api.UpdatePokemonLocal)
	pokegrouplocal.DELETE("/delete/:id", api.DeletePokemonLocal)
}

func InitServer(e *echo.Echo) *Server {
	return &Server{e: e}
}
