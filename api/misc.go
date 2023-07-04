package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/hokageCV/go-todo-api/types"
)

func GetRandomPokemon(c *fiber.Ctx) error {
	const pokeUrl = "https://pokeapi.co/api/v2/pokemon"
	randomInt := rand.Intn(200) + 1
	URL := pokeUrl + fmt.Sprintf("/%d", randomInt)
	fmt.Println(URL)

	resp, err := http.Get(URL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	bytesData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var pokemon types.Pokemon
	err = json.Unmarshal(bytesData, &pokemon) // bytes to json
	if err != nil {
		log.Fatal(err)
	}

	// map data to response
	typesArray := make([]string, len(pokemon.Types))
	for i, t := range pokemon.Types {
		typesArray[i] = t.Type.Name
	}

	abilitiesArray := make([]string, len(pokemon.Abilities))
	for i, a := range pokemon.Abilities {
		abilitiesArray[i] = a.Ability.Name
	}

	response := types.PokemonResponse{
		Name:           pokemon.Name,
		Order:          pokemon.Order,
		Types:          typesArray,
		Abilities:      abilitiesArray,
		BaseExperience: pokemon.BaseExperience,
	}

	return c.JSON(response)

}
