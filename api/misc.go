package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var pokemon types.Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		log.Fatal(err)
	}

	response := types.PokemonResponse{
		Pokemon: pokemon,
	}

	return c.JSON(response)

}
