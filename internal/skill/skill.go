package skill

import (
	"pokemons/internal/element"
)

const MaxPossibleNumberOfSkills = 3

type Type int

const (
	Common = iota
	Special
)

type Skill interface {
	SetPower(p float32)
	GetPower() float32
	GetUseCount() int
	GetUserSkillEffect() *skillEffect
	GetOpponentSkillEffect() *skillEffect
	GetName() string
	GetType() Type
	GetCopy() *skill
	element.Element
}

type skill struct {
	name           string
	power          float32
	useCount       int
	skillType      Type
	userEffect     skillEffect
	opponentEffect skillEffect
	element.Element
}

type skillEffect struct {
	Hp             float32
	Attack         float32
	Defense        float32
	SpecialAttack  float32
	SpecialDefense float32
}

func (s *skill) GetUserSkillEffect() *skillEffect {
	return &s.userEffect
}

func (s *skill) GetOpponentSkillEffect() *skillEffect {
	return &s.opponentEffect
}

func (s *skill) GetUseCount() int {
	return s.useCount
}

func (s *skill) GetName() string {
	return s.name
}

func (s *skill) GetPower() float32 {
	return s.power
}

func (s *skill) SetPower(p float32) {
	s.power = p
}

func (s *skill) GetType() Type {
	return s.skillType
}

func (s *skill) GetCopy() *skill {
	cpy := *s
	return &cpy
}
