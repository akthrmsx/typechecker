package basic

import "errors"

var (
	ErrTypeMismatch   = errors.New("type mismatch")
	ErrParamsMismatch = errors.New("params mismatch")
	ErrUnknownTerm    = errors.New("unknown term")
)

func Check(term Term, env Env) (Type, error) {
	switch v := term.(type) {
	case TermTrue:
		return TypeBool{}, nil
	case TermFalse:
		return TypeBool{}, nil
	case TermIf:
		if _, err := expectType[TypeBool](v.Cond, env); err != nil {
			return nil, err
		}
		cons, err := Check(v.Cons, env)
		if err != nil {
			return nil, err
		}
		alt, err := Check(v.Alt, env)
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
		if _, err := expectType[TypeInt](v.Left, env); err != nil {
			return nil, err
		}
		if _, err := expectType[TypeInt](v.Right, env); err != nil {
			return nil, err
		}
		return TypeInt{}, nil
	case TermVar:
		return env.Get(v.Name)
	case TermFunc:
		e := env.Clone()
		for _, param := range v.Params {
			e.Set(param.Name, param.Type.Clone())
		}
		ty, err := Check(v.Body, e)
		if err != nil {
			return nil, err
		}
		return TypeFunc{v.Params, ty}, nil
	case TermCall:
		f, err := expectType[TypeFunc](v.Func, env)
		if err != nil {
			return nil, err
		}
		if len(v.Args) != len(f.Params) {
			return nil, ErrParamsMismatch
		}
		for i, param := range f.Params {
			if err := matchParamType(v.Args[i], param, env); err != nil {
				return nil, err
			}
		}
		return f.ReturnType, nil
	case TermSeq:
		_, err := Check(v.First, env)
		if err != nil {
			return nil, err
		}
		return Check(v.Second, env)
	case TermConst:
		ty, err := Check(v.Value, env)
		if err != nil {
			return nil, err
		}
		e := env.Clone()
		e.Set(v.Name, ty)
		return Check(v.Next, e)
	default:
		return nil, ErrUnknownTerm
	}
}

func expectType[T Type](term Term, env Env) (T, error) {
	var zero T
	ty, err := Check(term, env)
	if err != nil {
		return zero, err
	}
	t, ok := (ty).(T)
	if !ok {
		return zero, ErrTypeMismatch
	}
	return t, nil
}

func matchParamType(term Term, param Param, env Env) error {
	ty, err := Check(term, env)
	if err != nil {
		return err
	}
	if !ty.Equals(param.Type) {
		return ErrTypeMismatch
	}
	return nil
}
