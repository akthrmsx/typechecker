package basic

type Term interface {
	IsTerm()
}

type TermTrue struct{}

type TermFalse struct{}

type TermIf struct {
	Cond Term
	Cons Term
	Alt  Term
}

type TermInt struct {
	Value int
}

type TermAdd struct {
	Left  Term
	Right Term
}

type TermVar struct {
	Name string
}

type TermFunc struct {
	Params Params
	Body   Term
}

type TermCall struct {
	Func Term
	Args Args
}

type TermSeq struct {
	First  Term
	Second Term
}

type TermConst struct {
	Name  string
	Value Term
	Next  Term
}

func (t TermTrue) IsTerm()  {}
func (t TermFalse) IsTerm() {}
func (t TermIf) IsTerm()    {}
func (t TermInt) IsTerm()   {}
func (t TermAdd) IsTerm()   {}
func (t TermVar) IsTerm()   {}
func (t TermFunc) IsTerm()  {}
func (t TermCall) IsTerm()  {}
func (t TermSeq) IsTerm()   {}
func (t TermConst) IsTerm() {}

var (
	_ Term = TermTrue{}
	_ Term = TermFalse{}
	_ Term = TermIf{}
	_ Term = TermInt{}
	_ Term = TermAdd{}
	_ Term = TermVar{}
	_ Term = TermFunc{}
	_ Term = TermCall{}
	_ Term = TermSeq{}
	_ Term = TermConst{}
)

type Args []Term
