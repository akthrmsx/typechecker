package arith

type Type interface {
	IsType()
	Equals(other Type) bool
}

type TypeBool struct{}

func (t TypeBool) IsType() {}

func (t TypeBool) Equals(other Type) bool {
	_, ok := (other).(TypeBool)
	return ok
}

type TypeInt struct{}

func (t TypeInt) IsType() {}

func (t TypeInt) Equals(other Type) bool {
	_, ok := (other).(TypeInt)
	return ok
}

var (
	_ Type = TypeBool{}
	_ Type = TypeInt{}
)
