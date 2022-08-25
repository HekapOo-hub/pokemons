package modifier

import (
	"pokemons/internal/element"
	"pokemons/internal/skill"
)

type attackModifier struct {
}

func (am *attackModifier) GetType() Type {
	return Attack
}

type commonAttackModifier struct {
	attack float32
	attackModifier
}

func NewCommonAttackModifier(attack float32) *commonAttackModifier {
	return &commonAttackModifier{attack: attack}
}

func (c *commonAttackModifier) Modify(s skill.Skill) {
	if s.GetType() == skill.Common {
		s.SetPower(s.GetPower() + c.attack)
	}
}

type specialAttackModifier struct {
	specialAttack float32
	attackModifier
}

func NewSpecialAttackModifier(specialAttack float32) *specialAttackModifier {
	return &specialAttackModifier{specialAttack: specialAttack}
}

func (m *specialAttackModifier) Modify(s skill.Skill) {
	if s.GetType() == skill.Special {
		s.SetPower(s.GetPower() + m.specialAttack)
	}
}

type elementAttackModifier struct {
	element element.Element
	attackModifier
}

func NewElementAttackModifier(e element.Element) *elementAttackModifier {
	return &elementAttackModifier{element: e}
}

func (e *elementAttackModifier) Modify(s skill.Skill) {
	factor := s.UserEffectiveness(e.element)
	s.SetPower(s.GetPower() * factor)
}
