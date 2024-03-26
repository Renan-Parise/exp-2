package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2/pokemon/"

type Pokemon struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Types     []Type    `json:"types"`
	Abilities []Ability `json:"abilities"`
	Stats     []Stat    `json:"stats"`
}

type Type struct {
	Type TypeInfo `json:"type"`
}

type TypeInfo struct {
	Name string `json:"name"`
}

type Ability struct {
	AbilityInfo AbilityInfo `json:"ability"`
}

type AbilityInfo struct {
	Name string `json:"name"`
}

type Stat struct {
	StatInfo StatInfo `json:"stat"`
}

type StatInfo struct {
	Name string `json:"name"`
}

func main() {
	pokemonName := "pikachu"

	resp, err := http.Get(baseURL + pokemonName)
	if err != nil {
		fmt.Println("Erro ao fazer solicitação:", err)
		return
	}
	defer resp.Body.Close()

	var pokemon Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		fmt.Println("Erro ao decodificar resposta JSON:", err)
		return
	}

	fmt.Println("ID:", pokemon.ID)

	fmt.Println("Nome:", pokemon.Name)

	fmt.Println("Tipos:")
	for i, t := range pokemon.Types {
		fmt.Println(i, t.Type.Name)
	}

	fmt.Println("Habilidades:")
	for i, a := range pokemon.Abilities {
		fmt.Println(i, a.AbilityInfo.Name)
	}

	fmt.Println("Stats:")
	for i, s := range pokemon.Stats {
		fmt.Println(i, s.StatInfo.Name)
	}
}
