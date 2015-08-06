package duck

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

//No need to test Lt because it is used internally by Gt
func TestGt(t *testing.T) {
	cases := []struct {
		A1  interface{}
		A2  interface{}
		Ok  bool
		Out bool
	}{
		{"2", true, true, true},
		{"0", true, true, false},
		{"2.0", true, true, true},
		{"0.0", true, true, false},
		{" 2 ", false, false, true},
		{"Inf", true, true, true},
	}

	for _, c := range cases {
		out, ok := Gt(c.A1, c.A2)
		require.Equal(t, c.Ok, ok, JSONString(c))
		if c.Ok {
			require.Equal(t, c.Out, out, JSONString(c))
		}
	}

	//Now reverse the arguments and outputs!
	for _, c := range cases {
		out, ok := Gt(c.A2, c.A1)
		require.Equal(t, c.Ok, ok, fmt.Sprintf("Reversed: %s", JSONString(c)))
		if c.Ok {
			require.Equal(t, !c.Out, out, fmt.Sprintf("Reversed: %s", JSONString(c)))
		}
	}
}

//No need to test Lte because it is used internally by Gte
func TestGte(t *testing.T) {
	cases := []struct {
		A1  interface{}
		A2  interface{}
		Ok  bool
		Out bool
	}{
		{"2", true, true, true},
		{"0", true, true, false},
		{"2.0", true, true, true},
		{"0.0", true, true, false},
		{" 2 ", false, false, true},
		{"Inf", true, true, true},
	}

	for _, c := range cases {
		out, ok := Gte(c.A1, c.A2)
		require.Equal(t, c.Ok, ok, JSONString(c))
		if c.Ok {
			require.Equal(t, c.Out, out, JSONString(c))
		}
	}

	//Now reverse the arguments and outputs!
	for _, c := range cases {
		out, ok := Gte(c.A2, c.A1)
		require.Equal(t, c.Ok, ok, fmt.Sprintf("Reversed: %v", JSONString(c)))
		if c.Ok {
			require.Equal(t, !c.Out, out, fmt.Sprintf("Reversed: %v", JSONString(c)))
		}
	}
}

func TestEqual(t *testing.T) {
	cases := []struct {
		A1  interface{}
		A2  interface{}
		Ok  bool
		Out bool
	}{
		{"2", true, true, false},
		{"0", true, true, false},
		{"2.0", true, true, false},
		{"0.0", true, true, false},
		{" 2 ", false, false, false},
		{"Inf", true, true, false},
		{"NaN", math.NaN(), true, true},
		{"2.0", "2", true, true},
		{"hello", "world", true, false},
	}

	for _, c := range cases {
		out, ok := Eq(c.A1, c.A2)
		require.Equal(t, c.Ok, ok, JSONString(c))
		if c.Ok {
			require.Equal(t, c.Out, out, JSONString(c))
		}
	}

	//Now reverse the arguments!
	for _, c := range cases {
		out, ok := Eq(c.A2, c.A1)
		require.Equal(t, c.Ok, ok, fmt.Sprintf("Reversed: %v", JSONString(c)))
		if c.Ok {
			require.Equal(t, c.Out, out, fmt.Sprintf("Reversed: %v", JSONString(c)))
		}
	}
}

func TestCmp(t *testing.T) {
	cases := []struct {
		A1  interface{}
		A2  interface{}
		Out int
	}{
		{"2", true, GreaterThan},
		{"0", true, LessThan},
		{"2.0", true, GreaterThan},
		{"0.0", true, LessThan},
		{" 2 ", false, CantCompare},
		{"Inf", true, GreaterThan},
		{"NaN", 3.3, CantCompare},
		{"NaN", math.NaN(), Equals},
	}

	for _, c := range cases {
		require.Equal(t, c.Out, Cmp(c.A1, c.A2), JSONString(c))

	}

	//Now reverse the arguments!
	for _, c := range cases {

		if c.Out == LessThan {
			c.Out = GreaterThan
		} else if c.Out == GreaterThan {
			c.Out = LessThan
		}
		require.Equal(t, c.Out, Cmp(c.A2, c.A1), JSONString(c))
	}
}
