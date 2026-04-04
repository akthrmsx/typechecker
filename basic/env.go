package basic

import "errors"

var ErrUndefinedVar = errors.New("undefined var")

type Env map[string]Type

func (e Env) Get(k string) (Type, error) {
	v, ok := e[k]
	if !ok {
		return nil, ErrUndefinedVar
	}
	return v, nil
}

func (e Env) Set(k string, v Type) {
	e[k] = v
}

func (e Env) Clone() Env {
	env := make(Env, len(e))
	for k, v := range e {
		env[k] = v.Clone()
	}
	return env
}
