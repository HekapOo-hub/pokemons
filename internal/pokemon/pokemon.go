package pokemon

import (
	"fmt"
	"pokemons/internal/element"
	"pokemons/internal/modifier"
	"pokemons/internal/skill"
)

type Pokemon interface {
	AddModifier(mod modifier.Modifier)
	AddModifierIfNotExist(mod modifier.Modifier)
	UseSkill(skillName string, opponent *pokemon) error
	Struggle(opponent *pokemon)
	PrepareForBattle()
}

type pokemon struct {
	name             string
	hp               float32
	maxHp            float32
	curPP            map[string]int
	maxPP            map[string]int
	attackModifiers  []modifier.Modifier
	defenseModifiers []modifier.Modifier
	skills           []skill.Skill
	statusOK         bool

	attack         float32
	defense        float32
	specialAttack  float32
	specialDefense float32
	element.Element
}

func NewPokemon(name string, e element.Element) *pokemon {
	return &pokemon{
		name:             name,
		hp:               100,
		maxHp:            100,
		curPP:            make(map[string]int),
		maxPP:            make(map[string]int),
		attackModifiers:  make([]modifier.Modifier, 0),
		defenseModifiers: make([]modifier.Modifier, 0),
		skills:           make([]skill.Skill, 0, skill.MaxPossibleNumberOfSkills),
		statusOK:         true,

		attack:         8,
		defense:        3,
		specialAttack:  6,
		specialDefense: 3,
		Element:        e,
	}
}

func (p *pokemon) LearnSkill(name skill.Prototype) error {
	if len(p.maxPP) == skill.MaxPossibleNumberOfSkills {
		return fmt.Errorf("%s can not learn more skills. You can replace existing one", p.name)
	}
	s := skill.NewSkill(name)
	p.curPP[s.GetName()] = s.GetUseCount()
	p.maxPP[s.GetName()] = s.GetUseCount()
	p.skills = append(p.skills, s)
	return nil
}

func (p *pokemon) ReplaceSkill(oldSkill skill.Skill, newSkill skill.Skill) {
	delete(p.curPP, oldSkill.GetName())
	delete(p.maxPP, oldSkill.GetName())
	p.curPP[newSkill.GetName()] = newSkill.GetUseCount()
	p.maxPP[newSkill.GetName()] = newSkill.GetUseCount()
}

func (p *pokemon) UseSkill(skillName skill.Prototype, opponent *pokemon) error {
	s := skill.NewSkill(skillName)
	skillNameStr := s.GetName()
	if p.maxPP[skillNameStr] == 0 {
		return fmt.Errorf("%s doesn't know this skill", p.name)
	}
	if p.curPP[skillNameStr] == 0 {
		return fmt.Errorf("%s doesn't have any power points on this skill, choose another skill", p.name)
	}

	for i := range p.attackModifiers {
		p.attackModifiers[i].Modify(s)
	}
	for i := range opponent.defenseModifiers {
		opponent.defenseModifiers[i].Modify(s)
	}
	applySkillEffect(s, p, opponent)
	p.curPP[skillNameStr]--
	// execute fight result
	return nil
}

func (p *pokemon) Struggle(opponent *pokemon) {
	opponent.hp -= p.attack - opponent.defense
}

func (p *pokemon) PrepareForBattle() {
	commonDefMod := modifier.NewCommonDefenseModifier(p.defense)
	commonAttackMod := modifier.NewCommonAttackModifier(p.attack)
	specialDefMod := modifier.NewSpecialDefenseModifier(p.specialDefense)
	specialAttackMod := modifier.NewSpecialAttackModifier(p.specialAttack)
	elemDefMod := modifier.NewElementDefenseModifier(p.Element)
	elemAttackMod := modifier.NewElementAttackModifier(p.Element)
	normalStatus := modifier.NewNormalStatus()
	p.attackModifiers = append(p.attackModifiers, normalStatus, commonAttackMod, specialAttackMod, elemAttackMod)
	p.defenseModifiers = append(p.defenseModifiers, commonDefMod, specialDefMod, elemDefMod)
}

func (p *pokemon) AddModifier(mod modifier.Modifier) {
	switch mod.GetType() {
	case modifier.Attack:
		p.attackModifiers = append(p.attackModifiers, mod)
	case modifier.Defense:
		p.defenseModifiers = append(p.defenseModifiers, mod)
	case modifier.Status:
		p.attackModifiers[0] = mod
	}
}

func (p *pokemon) AddModifierIfNotExist(mod modifier.Modifier) {
	switch mod.GetType() {
	case modifier.Attack:
		for i := range p.attackModifiers {
			if p.attackModifiers[i] == mod {
				return
			}
		}
		p.attackModifiers = append(p.attackModifiers, mod)
	case modifier.Defense:
		for i := range p.defenseModifiers {
			if p.attackModifiers[i] == mod {
				return
			}
		}
		p.defenseModifiers = append(p.defenseModifiers, mod)
	case modifier.Status:
		if p.statusOK {
			p.attackModifiers[0] = mod
			p.statusOK = false
		}
	}
}

func applySkillEffect(s skill.Skill, user *pokemon, opponent *pokemon) {

	effect := s.GetUserSkillEffect()

	user.hp += effect.Hp
	user.attack += effect.Attack
	user.defense += effect.Defense
	user.specialAttack += effect.SpecialAttack
	user.specialDefense += effect.SpecialDefense

	effect = s.GetOpponentSkillEffect()

	opponent.hp -= s.GetPower()
	opponent.hp += effect.Hp
	opponent.attack += effect.Attack
	opponent.defense += effect.Defense
	opponent.specialAttack += effect.SpecialAttack
	opponent.specialDefense += effect.SpecialDefense

}
