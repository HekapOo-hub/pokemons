package skill

import (
	"math/rand"
	"time"
)

type NormalStatus struct{}

func (ns *NormalStatus) Handle(s *Skill) bool {
	return true
}

type BurnStatus struct {
	DamagePerTurn float32
}

func (bs *BurnStatus) Handle(s *Skill) bool {
	s.SelfHp -= bs.DamagePerTurn
	return true
}

type FreezeStatus struct {
}

func (fs *FreezeStatus) Handle(s *Skill) bool {
	return false
}

type ParalysisStatus struct {
}

func (ps *ParalysisStatus) Handle(s *Skill) bool {
	rand.Seed(time.Now().Unix())
	probability := rand.Float64()
	if probability > 0.5 {
		return false
	}
	return true
}

type PoisonStatus struct {
	DamagePerTurn float32
}

func (ps *PoisonStatus) Handle(s *Skill) bool {
	s.SelfHp -= ps.DamagePerTurn
	return true
}

type ConfuseStatus struct{}

func (cs *ConfuseStatus) Handle(s *Skill) bool {
	rand.Seed(time.Now().Unix())
	probability := rand.Float64()
	if probability > 0.5 {
		// deal damage to itself
		s.SelfHp -= s.SelfAttack
		return false
	}
	return true
}

/*
type Status int

const (
	Burn Status = iota
	Freeze
	Paralysis
	Poisoned
	Sleep
)
*/
