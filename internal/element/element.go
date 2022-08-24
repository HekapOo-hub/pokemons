package element

const (
	SuperEffective   = 2
	Effective        = 1
	NotVeryEffective = 0.5
	NotEffective     = 0
)

type Element interface {
	EffectiveAgainst(e Element) float32
	UserEffectiveness(e Element) float32
}

type Fire struct{}

func (f Fire) EffectiveAgainst(e Element) float32 {
	switch e.(type) {
	case Fire:
		return NotVeryEffective
	case Water:
		return NotVeryEffective
	case Grass:
		return SuperEffective
	case Rock:
		return NotVeryEffective
	default:
		return Effective
	}
}

func (f Fire) UserEffectiveness(e Element) float32 {
	switch e.(type) {
	case Fire:
		return Effective
	default:
		return NotVeryEffective
	}
}

type Water struct{}

func (w Water) EffectiveAgainst(e Element) float32 {
	switch e.(type) {
	case Fire:
		return SuperEffective
	case Water:
		return NotVeryEffective
	case Grass:
		return NotVeryEffective
	case Rock:
		return SuperEffective
	default:
		return Effective
	}
}

func (w Water) UserEffectiveness(e Element) float32 {
	switch e.(type) {
	case Water:
		return Effective
	default:
		return NotVeryEffective
	}
}

type Thunder struct{}

func (t Thunder) EffectiveAgainst(e Element) float32 {
	switch e.(type) {
	case Water:
		return SuperEffective
	case Grass:
		return NotVeryEffective
	case Rock:
		return NotEffective
	default:
		return Effective
	}
}

func (t Thunder) UserEffectiveness(e Element) float32 {
	switch e.(type) {
	case Thunder:
		return Effective
	default:
		return NotVeryEffective
	}
}

type Rock struct{}

func (r Rock) EffectiveAgainst(e Element) float32 {
	switch e.(type) {
	case Water:
		return NotVeryEffective
	case Thunder:
		return SuperEffective
	case Rock:
		return NotVeryEffective
	default:
		return Effective
	}
}

func (r Rock) UserEffectiveness(e Element) float32 {
	switch e.(type) {
	case Rock:
		return Effective
	default:
		return NotVeryEffective
	}
}

type Grass struct{}

func (g Grass) EffectiveAgainst(e Element) float32 {
	switch e.(type) {
	case Water:
		return SuperEffective
	case Rock:
		return SuperEffective
	case Fire:
		return NotVeryEffective
	default:
		return 1
	}
}

func (g Grass) UserEffectiveness(e Element) float32 {
	switch e.(type) {
	case Grass:
		return Effective
	default:
		return NotVeryEffective
	}
}
