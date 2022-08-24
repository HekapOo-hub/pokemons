package skill

import (
	"pokemons/internal/element"
	"pokemons/internal/pokemon"
)

type Observer interface {
	Handle(s *Skill) bool
}

type baseAttackModifier struct {
	pokemon *pokemon.Pokemon
}

func NewBaseAttackModifier(p *pokemon.Pokemon) *baseAttackModifier {
	return &baseAttackModifier{
		pokemon: p,
	}
}

func (e *baseAttackModifier) Handle(s *Skill) bool {
	if s.AttackerID == e.pokemon.ID {
		switch s.Type {
		case Common:
			s.Power += e.pokemon.Attack
		case Special:
			s.Power += e.pokemon.SpecialAttack
		default:
		}
	}
	return true
}

type elementAttackModifier struct {
	pokemon *pokemon.Pokemon
	element.Element
}

func NewElementAttackModifier(p *pokemon.Pokemon, e element.Element) *elementAttackModifier {
	return &elementAttackModifier{
		pokemon: p,
		Element: e,
	}
}

func (e elementAttackModifier) Handle(s *Skill) bool {
	if s.AttackerID == e.pokemon.ID {
		factor := e.Element.UserEffectiveness(s.Element)
		s.Power *= factor
	}
	return true
}

type baseDefenseModifier struct {
	pokemon *pokemon.Pokemon
}

func NewBaseDefenseModifier(p *pokemon.Pokemon) *baseDefenseModifier {
	return &baseDefenseModifier{
		pokemon: p,
	}
}

func (e baseDefenseModifier) Handle(s *Skill) bool {
	if s.DefenderID == e.pokemon.ID {
		switch s.Type {
		case Common:
			s.Power -= e.pokemon.Defense
		case Special:
			s.Power -= e.pokemon.SpecialDefense
		default:
		}
	}
	return true
}

type elementDefenseModifier struct {
	pokemon *pokemon.Pokemon
	element.Element
}

func NewElementDefenseModifier(p *pokemon.Pokemon, e element.Element) *elementDefenseModifier {
	return &elementDefenseModifier{
		pokemon: p,
		Element: e,
	}
}

func (e elementDefenseModifier) Handle(s *Skill) bool {
	if s.DefenderID == e.pokemon.ID {
		factor := s.Element.EffectiveAgainst(e.Element)
		s.Power *= factor
	}
	return true
}
