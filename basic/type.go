package basic

type Type interface {
	IsType()
	Equals(other Type) bool
	Clone() Type
}

type TypeBool struct{}

type TypeInt struct{}

type TypeFunc struct {
	Params     Params
	ReturnType Type
}

func (t TypeBool) IsType() {}
func (t TypeInt) IsType()  {}
func (t TypeFunc) IsType() {}

func (t TypeBool) Equals(other Type) bool {
	_, ok := (other).(TypeBool)
	return ok
}

func (t TypeInt) Equals(other Type) bool {
	_, ok := (other).(TypeInt)
	return ok
}

func (t TypeFunc) Equals(other Type) bool {
	tt, ok := (other).(TypeFunc)
	if !ok {
		return false
	}
	if len(t.Params) != len(tt.Params) {
		return false
	}
	for i, param := range t.Params {
		if !param.Type.Equals(tt.Params[i].Type) {
			return false
		}
	}
	return t.ReturnType.Equals(tt.ReturnType)
}

func (t TypeBool) Clone() Type {
	return TypeBool{}
}

func (t TypeInt) Clone() Type {
	return TypeInt{}
}

func (t TypeFunc) Clone() Type {
	return TypeFunc{t.Params.Clone(), t.ReturnType.Clone()}
}

var (
	_ Type = TypeBool{}
	_ Type = TypeInt{}
	_ Type = TypeFunc{}
)
