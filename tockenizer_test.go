package calculator

import "testing"

func TestTokenize(t *testing.T) {
	got, err := Tokenize("(2+2)*2")
	except := []iToken{
		brToken{value: '('},
		numToken{value: 2},
		opToken{value: '+'},
		numToken{value: 2},
		brToken{value: ')'},
		opToken{value: '*'},
		numToken{value: 2},
	}
	if err != nil {
		t.Fatalf("Function Tokenize(%q): error: %q", "12+3/5*(3+5)", err)
	} else if len(got) != len(except) {
		t.Fatalf("Function Tokenize(%q): got %v, except %v", "12+3/5*(3+5)", got, except)
	} else {
		for i := range got {
			if got[i] != except[i] {
				t.Fatalf("Function Tokenize(%q): got %v, except %v", "12+3/5*(3+5)", got, except)
			}
		}
	}
}
