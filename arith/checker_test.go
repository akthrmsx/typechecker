package arith

import (
	"fmt"
	"testing"
)

type test struct {
	input Term
	want  Type
	err   bool
}

type tests []test

func TestTrue(t *testing.T) {
	tests := tests{
		{
			input: TermTrue{},
			want:  TypeBool{},
		},
	}
	runTests(t, tests)
}

func TestFalse(t *testing.T) {
	tests := tests{
		{
			input: TermFalse{},
			want:  TypeBool{},
		},
	}
	runTests(t, tests)
}

func TestIf(t *testing.T) {
	tests := tests{
		{
			input: TermIf{TermTrue{}, TermInt{1}, TermInt{2}},
			want:  TypeInt{},
		},
		{
			input: TermIf{TermInt{0}, TermInt{1}, TermInt{2}},
			err:   true,
		},
		{
			input: TermIf{TermTrue{}, TermAdd{TermTrue{}, TermInt{1}}, TermInt{1}},
			err:   true,
		},
		{
			input: TermIf{TermTrue{}, TermInt{1}, TermAdd{TermTrue{}, TermInt{1}}},
			err:   true,
		},
		{
			input: TermIf{TermTrue{}, TermInt{1}, TermTrue{}},
			err:   true,
		},
	}
	runTests(t, tests)
}

func TestInt(t *testing.T) {
	tests := tests{
		{
			input: TermInt{1},
			want:  TypeInt{},
		},
	}
	runTests(t, tests)
}

func TestAdd(t *testing.T) {
	tests := tests{
		{
			input: TermAdd{TermInt{1}, TermInt{2}},
			want:  TypeInt{},
		},
		{
			input: TermAdd{TermTrue{}, TermInt{1}},
			err:   true,
		},
		{
			input: TermAdd{TermAdd{TermTrue{}, TermInt{1}}, TermInt{1}},
			err:   true,
		},
		{
			input: TermAdd{TermInt{1}, TermTrue{}},
			err:   true,
		},
		{
			input: TermAdd{TermInt{1}, TermAdd{TermTrue{}, TermInt{1}}},
			err:   true,
		},
	}
	runTests(t, tests)
}

func runTests(t *testing.T, tests tests) {
	t.Helper()
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got, err := Check(tt.input)
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
