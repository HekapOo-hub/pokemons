package pokemon

import (
	"fmt"
	"github.com/google/uuid"
	"pokemons/internal/element"
)

type Pokemon struct {
	ID             uuid.UUID
	Name           string
	MaxHp          float32
	Hp             float32
	Attack         float32
	Defense        float32
	SpecialAttack  float32
	SpecialDefense float32
	Speed          float32
	SkillsCurPP    map[string]int
	SkillsMaxPP    map[string]int
	Element        element.Element
}

// should subscribe on attack, defense modifiers

func NewPokemon(name string, element element.Element) *Pokemon {
	return &Pokemon{
		ID:             uuid.New(),
		Name:           name,
		MaxHp:          100,
		Hp:             100,
		Attack:         10,
		Defense:        5,
		SpecialAttack:  8,
		SpecialDefense: 2,
		Speed:          10,
		Element:        element,
		SkillsCurPP:    make(map[string]int),
		SkillsMaxPP:    make(map[string]int),
	}
}

func (p *Pokemon) Use(skill string) error {
	curPP, ok := p.SkillsCurPP[skill]
	if !ok {
		return fmt.Errorf("%s with ID: %s doesn't know this move", p.Name, p.ID)
	}
	if curPP > 0 {
		p.SkillsCurPP[skill]--
		return nil
	}
	return fmt.Errorf("%s with ID: %s has no moer power points on this move", p.Name, p.ID)
}

// TODO: implement max number of skills, if number exceeded replace old skill with the new one

func (p *Pokemon) LearnSkill(skill string, maxPP int) error {
	_, ok := p.SkillsMaxPP[skill]
	if ok {
		return fmt.Errorf("%s with ID: %s already knows this move", p.Name, p.ID)
	}
	p.SkillsMaxPP[skill] = maxPP
	p.SkillsCurPP[skill] = maxPP
	return nil
}

func Recover(p *Pokemon) {
	p.Hp = p.MaxHp
	p.SkillsCurPP = p.SkillsMaxPP
}
