package skill

import (
	"pokemons/internal/element"
)

type Prototype string

const (
	Thunderbolt Prototype = "Thunderbolt"
	FireBlast             = "Fire blast"
	Absorb                = "Absorb"
	RockSlide             = "Rock slide"
)

func NewSkill(name Prototype) *skill {
	switch name {
	case Thunderbolt:
		return newThunderbolt()
	case FireBlast:
		return newFireBlast()
	case Absorb:
		return newAbsorb()
	case RockSlide:
		return newRockSlide()
	default:
		panic("skill is not implemented")
	}
}

func newThunderbolt() *skill {
	return &skill{
		name:      "Thunderbolt",
		power:     15,
		useCount:  15,
		skillType: Common,
		Element:   element.Thunder{},
	}
}

func newFireBlast() *skill {
	return &skill{
		name:           "Fire blast",
		power:          20,
		useCount:       5,
		skillType:      Special,
		Element:        element.Fire{},
		userEffect:     skillEffect{},
		opponentEffect: skillEffect{},
	}
}

func newRockSlide() *skill {
	return &skill{
		name:           "Rock slide",
		power:          15,
		useCount:       15,
		skillType:      Special,
		Element:        element.Rock{},
		userEffect:     skillEffect{},
		opponentEffect: skillEffect{},
	}
}

func newAbsorb() *skill {
	return &skill{
		name:           "Absorb",
		power:          10,
		useCount:       20,
		skillType:      Common,
		Element:        element.Grass{},
		userEffect:     skillEffect{Hp: 5},
		opponentEffect: skillEffect{},
	}
}
