package skill

import (
	"github.com/google/uuid"
	"pokemons/internal/element"
	"pokemons/internal/pokemon"
)

type SkillType int

const (
	Common SkillType = iota
	Special
)

// should contain modifiers

type Skill struct {
	Name       string
	AttackerID uuid.UUID
	DefenderID uuid.UUID
	Power      float32
	Type       SkillType
	Element    element.Element
	skillEffect
}

type skillEffect struct {
	SelfHp             float32
	SelfAttack         float32
	SelfDefense        float32
	SelfSpecialAttack  float32
	SelfSpecialDefense float32
	SelfSpeed          float32
	SelfStatus         Observer

	OpAttack         float32
	OpDefense        float32
	OpSpecialAttack  float32
	OpSpecialDefense float32
	OpSpeed          float32
	OpStatus         Observer
}

func (s *Skill) Apply(self, opponent *pokemon.Pokemon) {
	self.Hp += s.SelfHp
	self.Attack += s.SelfAttack
	self.Defense += s.SelfDefense
	self.Speed += s.SelfSpeed
	self.SpecialAttack += s.SelfSpecialAttack
	self.SpecialDefense += s.SelfSpecialDefense

	opponent.Hp -= s.Power
	opponent.Attack += s.OpAttack
	opponent.Defense += s.OpDefense
	opponent.Speed += s.OpSpeed
	opponent.SpecialAttack += s.OpSpecialAttack
	opponent.SpecialDefense += s.OpSpecialDefense
}

func NewThunderBolt(attackerID, defenderID uuid.UUID) *Skill {
	return &Skill{
		AttackerID: attackerID,
		DefenderID: defenderID,
		Type:       Common,
		Power:      10,
		Element:    element.Thunder{},
	}
}

func NewFirePunch(attackerID, defenderID uuid.UUID) *Skill {
	return &Skill{
		AttackerID: attackerID,
		DefenderID: defenderID,
		Type:       Common,
		Power:      9,
		Element:    element.Fire{},
	}
}

func NewRockSlide(attackerID, defenderID uuid.UUID) *Skill {
	return &Skill{
		AttackerID: attackerID,
		DefenderID: defenderID,
		Type:       Common,
		Power:      8,
		Element:    element.Rock{},
	}
}
