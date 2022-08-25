package main

import (
	"fmt"
	"pokemons/internal/element"
	"pokemons/internal/pokemon"
	"pokemons/internal/skill"
)

func main() {
	charizard := pokemon.NewPokemon("charizard", element.Fire{})
	bulbasaur := pokemon.NewPokemon("bulbasaur", element.Grass{})

	charizard.LearnSkill(skill.FireBlast)
	bulbasaur.LearnSkill(skill.Absorb)

	charizard.PrepareForBattle()
	bulbasaur.PrepareForBattle()

	err := charizard.UseSkill(skill.FireBlast, bulbasaur)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", bulbasaur)
	err = charizard.UseSkill(skill.FireBlast, bulbasaur)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", bulbasaur)

	err = bulbasaur.UseSkill(skill.Absorb, charizard)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", bulbasaur)
	fmt.Printf("%+v\n", charizard)
}
