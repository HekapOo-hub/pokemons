package modifier

import (
	"pokemons/internal/skill"
)

type statusModifier struct {
}

func (*statusModifier) GetType() Type {
	return Status
}

type normalStatus struct {
	statusModifier
}

func NewNormalStatus() *normalStatus {
	return &normalStatus{}
}

func (ns *normalStatus) Modify(s skill.Skill) {}

type burnStatus struct {
	damagePerTurn float32
	statusModifier
}

func NewBurnStatus(damagePerTurn float32) *burnStatus {
	return &burnStatus{damagePerTurn: damagePerTurn}
}

func (bs *burnStatus) Modify(s skill.Skill) {
	effect := s.GetUserSkillEffect()
	effect.Hp -= bs.damagePerTurn
}
