package arith

type Term interface {
	IsTerm()
}

type TermTrue struct{}

func (t TermTrue) IsTerm() {}

type TermFalse struct{}

func (t TermFalse) IsTerm() {}

type TermIf struct {
	Cond Term
	Cons Term
	Alt  Term
}

func (t TermIf) IsTerm() {}

type TermInt struct {
	Value int
}

func (t TermInt) IsTerm() {}

type TermAdd struct {
	Left  Term
	Right Term
}

func (t TermAdd) IsTerm() {}

var (
	_ Term = TermTrue{}
	_ Term = TermFalse{}
	_ Term = TermIf{}
	_ Term = TermInt{}
	_ Term = TermAdd{}
)
