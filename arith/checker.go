package arith

import (
	"errors"
)

var (
	ErrTypeMismatch = errors.New("type mismatch")
	ErrUnknownTerm  = errors.New("unknown term")
)

func Check(term Term) (Type, error) {
	switch v := term.(type) {
	case TermTrue:
		return TypeBool{}, nil
	case TermFalse:
		return TypeBool{}, nil
	case TermIf:
		if _, err := expectType[TypeBool](v.Cond); err != nil {
			return nil, err
		}
		cons, err := Check(v.Cons)
		if err != nil {
			return nil, err
		}
		alt, err := Check(v.Alt)
		if err != nil {
			return nil, err
		}
		if !cons.Equals(alt) {
			return nil, ErrTypeMismatch
		}
		return cons, nil
	case TermInt:
		return TypeInt{}, nil
	case TermAdd:
		if _, err := expectType[TypeInt](v.Left); err != nil {
			return nil, err
		}
		if _, err := expectType[TypeInt](v.Right); err != nil {
			return nil, err
		}
		return TypeInt{}, nil
	default:
		return nil, ErrUnknownTerm
	}
}

func expectType[T Type](term Term) (T, error) {
	var zero T
	ty, err := Check(term)
	if err != nil {
		return zero, err
	}
	t, ok := (ty).(T)
	if !ok {
		return zero, ErrTypeMismatch
	}
	return t, nil
}
