package main

import (
	"fmt"
	"pokemons/internal/battle"
	"pokemons/internal/element"
	"pokemons/internal/pokemon"
	"pokemons/internal/skill"
)

func main() {
	p1 := pokemon.NewPokemon("charizard", element.Fire{})
	p2 := pokemon.NewPokemon("bulbasaur", element.Grass{})
	b := battle.NewBattle(p1, p2)

	firePunch := skill.NewFirePunch(p1.ID, p2.ID)
	p1.LearnSkill(firePunch.Name, 20)
	thunderBolt := skill.NewThunderBolt(p2.ID, p1.ID)
	p2.LearnSkill(thunderBolt.Name, 15)
	b.UseSkill(firePunch)
	fmt.Printf("%+v\n", p1)
	fmt.Printf("%+v\n", p2)

	b.UseSkill(thunderBolt)
	fmt.Printf("%+v\n", p1)
	fmt.Printf("%+v\n", p2)
}
