package battle

import (
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"pokemons/internal/pokemon"
	skill2 "pokemons/internal/skill"
	"sync"
)

type Observable interface {
	PushBack(pokemonID uuid.UUID, o skill2.Observer)
	PushFront(pokemonID uuid.UUID, o skill2.Observer)
	ReplaceFront(pokemonID uuid.UUID, o skill2.Observer)
	Unsubscribe(pokemonID uuid.UUID, o skill2.Observer)
}

type Battle interface {
	UseSkill(s *skill2.Skill) error
	JoinBattle(p *pokemon.Pokemon)
}

type pokemonModifier struct {
	modifier skill2.Observer
	next     *pokemonModifier
}

type pokemonModifiersList struct {
	head *pokemonModifier
}

type battle struct {
	modifiers      map[uuid.UUID]pokemonModifiersList
	participantsMu sync.RWMutex
	participants   map[uuid.UUID]*pokemon.Pokemon
}

func NewBattle(pokemons ...*pokemon.Pokemon) *battle {
	participants := make(map[uuid.UUID]*pokemon.Pokemon)
	for _, p := range pokemons {
		participants[p.ID] = p
	}

	b := &battle{
		modifiers:    make(map[uuid.UUID]pokemonModifiersList),
		participants: participants,
	}
	for id, p := range participants {
		b.PushFront(id, &skill2.NormalStatus{})
		b.PushBack(id, skill2.NewBaseAttackModifier(p))
		b.PushBack(id, skill2.NewBaseDefenseModifier(p))
		b.PushBack(id, skill2.NewElementAttackModifier(p, p.Element))
		b.PushBack(id, skill2.NewElementDefenseModifier(p, p.Element))
	}
	return b
}

// PushBack base modifiers should be in newBattle
func (b *battle) PushBack(pokemonID uuid.UUID, o skill2.Observer) {
	list, ok := b.modifiers[pokemonID]
	if !ok {
		list = pokemonModifiersList{head: &pokemonModifier{modifier: o}}
		b.modifiers[pokemonID] = list
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = &pokemonModifier{modifier: o}
}

func (b *battle) PushFront(pokemonID uuid.UUID, o skill2.Observer) {
	list, ok := b.modifiers[pokemonID]
	if !ok {
		list = pokemonModifiersList{head: &pokemonModifier{modifier: o}}
		b.modifiers[pokemonID] = list
		return
	}
	current := list.head
	list.head = &pokemonModifier{modifier: o, next: current}
}

func (b *battle) ReplaceFront(pokemonID uuid.UUID, o skill2.Observer) {
	list, ok := b.modifiers[pokemonID]
	if !ok {
		list = pokemonModifiersList{head: &pokemonModifier{modifier: o}}
		b.modifiers[pokemonID] = list
		return
	}
	list.head = &pokemonModifier{modifier: o, next: list.head.next}
}

func (b *battle) Unsubscribe(pokemonID uuid.UUID, o skill2.Observer) {
	list, ok := b.modifiers[pokemonID]
	if !ok {
		return
	}
	if list.head.modifier == o {
		list.head = list.head.next
		return
	}
	prev := list.head
	cur := list.head.next
	for cur != nil {
		if cur.modifier == o {
			prev.next = cur.next
			cur.next = nil
			return
		}
		cur = cur.next
		prev = prev.next
	}
}

func (b *battle) UseSkill(s *skill2.Skill) error {
	attacker, ok := b.participants[s.AttackerID]
	if !ok {
		return fmt.Errorf("pokemon with ID: %s does not participate in this battle", s.AttackerID)
	}
	defender, ok := b.participants[s.DefenderID]
	if !ok {
		return fmt.Errorf("pokemon with ID: %s does not participate in this battle", s.DefenderID)
	}
	err := attacker.Use(s.Name)
	if err != nil {
		return err
	}
	for curModifier := b.modifiers[s.AttackerID].head; curModifier != nil; curModifier = curModifier.next {
		ok = curModifier.modifier.Handle(s)
		if !ok {
			return nil
		}
	}
	for curModifier := b.modifiers[s.DefenderID].head; curModifier != nil; curModifier = curModifier.next {
		ok = curModifier.modifier.Handle(s)
		if !ok {
			return nil
		}
	}
	s.Apply(attacker, defender)
	b.ReplaceFront(attacker.ID, s.SelfStatus)
	b.ReplaceFront(defender.ID, s.OpStatus)
	b.analyzeTurnResult(attacker, defender)
	return nil
}

func (b *battle) JoinBattle(p *pokemon.Pokemon) {
	b.participantsMu.Lock()
	b.participants[p.ID] = p
	b.PushFront(p.ID, &skill2.NormalStatus{})
	b.PushBack(p.ID, skill2.NewBaseAttackModifier(p))
	b.PushBack(p.ID, skill2.NewBaseDefenseModifier(p))
	b.PushBack(p.ID, skill2.NewElementAttackModifier(p, p.Element))
	b.PushBack(p.ID, skill2.NewElementDefenseModifier(p, p.Element))
	b.participantsMu.Unlock()
}

func (b *battle) analyzeTurnResult(attacker, defender *pokemon.Pokemon) {
	if defender.Hp <= 0 {
		b.participantsMu.Lock()
		delete(b.participants, defender.ID)
		b.participantsMu.Unlock()
		log.Infof("%s with ID: %s fainted", defender.Name, defender.ID)
	}
	if attacker.Hp <= 0 {
		b.participantsMu.Lock()
		delete(b.participants, attacker.ID)
		b.participantsMu.Unlock()
		log.Infof("pokemon with ID: %s fainted", attacker.ID)
	}
	b.participantsMu.RLock()
	switch len(b.participants) {
	case 1:
		for id, p := range b.participants {
			log.Infof("%s with ID: %s wins", p.Name, id)
		}
	case 0:
		log.Info("The battle ended in a draw")
	}
	b.participantsMu.RUnlock()
}
