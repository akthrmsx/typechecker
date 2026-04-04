package basic

type Param struct {
	Name string
	Type Type
}

func (p Param) Clone() Param {
	return Param{p.Name, p.Type.Clone()}
}

type Params []Param

func (p Params) Clone() Params {
	params := make(Params, 0, len(p))
	for _, param := range p {
		params = append(params, param.Clone())
	}
	return params
}
