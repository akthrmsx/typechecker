package basic

import (
	"fmt"
	"testing"
)

type test struct {
	input Term
	env   Env
	want  Type
	err   bool
}

type tests []test

func TestTrue(t *testing.T) {
	tests := tests{
		{
			input: TermTrue{},
			env:   Env{},
			want:  TypeBool{},
		},
	}
	runTests(t, tests)
}

func TestFalse(t *testing.T) {
	tests := tests{
		{
			input: TermFalse{},
			env:   Env{},
			want:  TypeBool{},
		},
	}
	runTests(t, tests)
}

func TestIf(t *testing.T) {
	tests := tests{
		{
			input: TermIf{TermTrue{}, TermInt{1}, TermInt{2}},
			env:   Env{},
			want:  TypeInt{},
		},
		{
			input: TermIf{TermAdd{TermTrue{}, TermInt{0}}, TermInt{1}, TermInt{2}},
			env:   Env{},
			err:   true,
		},
		{
			input: TermIf{TermInt{0}, TermInt{1}, TermInt{2}},
			env:   Env{},
			err:   true,
		},
		{
			input: TermIf{TermTrue{}, TermAdd{TermTrue{}, TermInt{1}}, TermInt{2}},
			env:   Env{},
			err:   true,
		},
		{
			input: TermIf{TermTrue{}, TermInt{1}, TermAdd{TermTrue{}, TermInt{2}}},
			env:   Env{},
			err:   true,
		},
		{
			input: TermIf{TermTrue{}, TermInt{1}, TermTrue{}},
			env:   Env{},
			err:   true,
		},
	}
	runTests(t, tests)
}

func TestInt(t *testing.T) {
	tests := tests{
		{
			input: TermInt{1},
			env:   Env{},
			want:  TypeInt{},
		},
	}
	runTests(t, tests)
}

func TestAdd(t *testing.T) {
	tests := tests{
		{
			input: TermAdd{TermInt{1}, TermInt{2}},
			env:   Env{},
			want:  TypeInt{},
		},
		{
			input: TermAdd{TermTrue{}, TermInt{1}},
			env:   Env{},
			err:   true,
		},
		{
			input: TermAdd{TermAdd{TermTrue{}, TermInt{1}}, TermInt{1}},
			env:   Env{},
			err:   true,
		},
		{
			input: TermAdd{TermInt{1}, TermTrue{}},
			env:   Env{},
			err:   true,
		},
		{
			input: TermAdd{TermInt{1}, TermAdd{TermTrue{}, TermInt{1}}},
			env:   Env{},
			err:   true,
		},
	}
	runTests(t, tests)
}

func TestVar(t *testing.T) {
	tests := tests{
		{
			input: TermVar{"a"},
			env:   Env{"a": TypeBool{}},
			want:  TypeBool{},
		},
		{
			input: TermVar{"a"},
			env:   Env{},
			err:   true,
		},
	}
	runTests(t, tests)
}

func TestFunc(t *testing.T) {
	tests := tests{
		{
			input: TermFunc{Params{}, TermTrue{}},
			env:   Env{},
			want:  TypeFunc{Params{}, TypeBool{}},
		},
		{
			input: TermFunc{Params{Param{"a", TypeInt{}}, Param{"b", TypeInt{}}}, TermAdd{TermVar{"a"}, TermVar{"b"}}},
			env:   Env{},
			want:  TypeFunc{Params{Param{"a", TypeInt{}}, Param{"b", TypeInt{}}}, TypeInt{}},
		},
		{
			input: TermFunc{Params{}, TermVar{"a"}},
			env:   Env{},
			err:   true,
		},
	}
	runTests(t, tests)
}

func TestCall(t *testing.T) {
	tests := tests{
		{
			input: TermCall{TermVar{"f"}, Args{}},
			env:   Env{"f": TypeFunc{Params{}, TypeBool{}}},
			want:  TypeBool{},
		},
		{
			input: TermCall{TermVar{"f"}, Args{TermInt{1}, TermInt{2}}},
			env:   Env{"f": TypeFunc{Params{Param{"a", TypeInt{}}, Param{"b", TypeInt{}}}, TypeBool{}}},
			want:  TypeBool{},
		},
		{
			input: TermCall{TermVar{"f"}, Args{}},
			env:   Env{},
			err:   true,
		},
		{
			input: TermCall{TermVar{"f"}, Args{}},
			env:   Env{"f": TypeBool{}},
			err:   true,
		},
		{
			input: TermCall{TermVar{"f"}, Args{}},
			env:   Env{"f": TypeFunc{Params{Param{"a", TypeInt{}}}, TypeBool{}}},
			err:   true,
		},
		{
			input: TermCall{TermVar{"f"}, Args{TermVar{"a"}}},
			env:   Env{"f": TypeFunc{Params{Param{"a", TypeInt{}}}, TypeBool{}}},
			err:   true,
		},
		{
			input: TermCall{TermVar{"f"}, Args{TermTrue{}}},
			env:   Env{"f": TypeFunc{Params{Param{"a", TypeInt{}}}, TypeBool{}}},
			err:   true,
		},
	}
	runTests(t, tests)
}

func TestSeq(t *testing.T) {
	tests := tests{
		{
			input: TermSeq{TermInt{1}, TermSeq{TermInt{2}, TermTrue{}}},
			env:   Env{},
			want:  TypeBool{},
		},
		{
			input: TermSeq{TermVar{"a"}, TermTrue{}},
			env:   Env{},
			err:   true,
		},
		{
			input: TermSeq{TermTrue{}, TermVar{"a"}},
			env:   Env{},
			err:   true,
		},
	}
	runTests(t, tests)
}

func TestConst(t *testing.T) {
	tests := tests{
		{
			input: TermConst{"a", TermInt{1}, TermAdd{TermVar{"a"}, TermInt{1}}},
			env:   Env{},
			want:  TypeInt{},
		},
		{
			input: TermConst{"a", TermVar{"a"}, TermAdd{TermVar{"a"}, TermInt{1}}},
			env:   Env{},
			err:   true,
		},
		{
			input: TermConst{"a", TermInt{1}, TermVar{"b"}},
			env:   Env{},
			err:   true,
		},
	}
	runTests(t, tests)
}

func runTests(t *testing.T, tests tests) {
	t.Helper()
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got, err := Check(tt.input, tt.env)
			if tt.err {
				if err == nil {
					t.Fatal("error expected")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if !tt.want.Equals(got) {
				t.Errorf("want: %#v, but got: %#v", tt.want, got)
			}
		})
	}
}
