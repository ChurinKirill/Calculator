package calculator

import (
	"fmt"
	"math"
	"testing"
)

type Test struct {
	in  string
	out float64
	err error
}

var tests = []Test{
	// {"12+3/5*(3+5)", 16.8},
	// {"14*349/(3+8/2*(45-83)+2)", -33.2380952381},
	// {"2/3+4/5-6/7/(8/9-10/11)+12/13", 44.8183150183},
	// {"28347+9278/(273+943/65)*0.002", 28347.0645409},
	{"1+1", 2, nil},
	{"(2+2)*2", 8, nil},
	{"2+2*2", 6, nil},
	{"1/2", 0.5, nil},
	{"1+1*", 0, fmt.Errorf("error")},
}

func TestCalc(t *testing.T) {
	for i, test := range tests {
		precision := 6
		res, err := Calc(test.in)
		roundRes := math.Round(res*math.Pow(10, float64(precision))) / math.Pow(10, float64(precision))
		roundOut := math.Round(test.out*math.Pow(10, float64(precision))) / math.Pow(10, float64(precision))
		if test.err == nil && err != nil {
			t.Fatalf("#%d: Calc(%q) returns error: %v", i, test.in, err)
		} else if test.err == nil && roundRes != roundOut {
			t.Fatalf("#%d: Calc(%q): got %f, except %f", i, test.in, res, test.out)
		} else if test.err != nil && err == nil {
			t.Fatalf("There's must be error")
		}
	}
}
