package types

// ==================================
type Todo struct {
	ID     string `json:"id" db:"primary_key"`
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

// ==================================

type Pokemon struct {
	Name  string `json:"name"`
	Order int    `json:"order"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
		} `json:"ability"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
}

type PokemonResponse struct {
	Pokemon Pokemon `json:"pokemon"`
}

// ==================================
