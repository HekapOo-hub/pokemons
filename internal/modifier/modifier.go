package modifier

import "pokemons/internal/skill"

type Type int

const (
	Attack = iota
	Defense
	Status
)

type Modifier interface {
	Modify(s skill.Skill)
	GetType() Type
}
