package modifier

import (
	"pokemons/internal/element"
	"pokemons/internal/skill"
)

type defenseModifier struct {
}

func (*defenseModifier) GetType() Type {
	return Defense
}

type commonDefenseModifier struct {
	defenseModifier
	defense float32
}

func NewCommonDefenseModifier(defense float32) *commonDefenseModifier {
	return &commonDefenseModifier{defense: defense}
}
func (c *commonDefenseModifier) Modify(s skill.Skill) {
	if s.GetType() == skill.Common {
		s.SetPower(s.GetPower() - c.defense)
	}
}

type specialDefenseModifier struct {
	defenseModifier
	specialDefense float32
}

func NewSpecialDefenseModifier(specialDefense float32) *specialDefenseModifier {
	return &specialDefenseModifier{specialDefense: specialDefense}
}

func (sdm *specialDefenseModifier) Modify(s skill.Skill) {
	if s.GetType() == skill.Special {
		s.SetPower(s.GetPower() - sdm.specialDefense)
	}
}

type elementDefenseModifier struct {
	defenseModifier
	element element.Element
}

func NewElementDefenseModifier(e element.Element) *elementDefenseModifier {
	return &elementDefenseModifier{element: e}
}

func (e *elementDefenseModifier) Modify(s skill.Skill) {
	factor := s.EffectiveAgainst(e.element)
	s.SetPower(s.GetPower() * factor)
}
