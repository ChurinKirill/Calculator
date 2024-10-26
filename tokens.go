package calculator

const (
	numT int = 0
	opT  int = 1
	brT  int = 2
)

type iToken interface {
	getType() int
	getValue() interface{}
}

type numToken struct {
	value float64
}

type opToken struct {
	value rune
}

type brToken struct {
	value rune
}

func (token numToken) getType() int {
	return numT
}
func (token numToken) getValue() interface{} {
	return token.value
}

func (token opToken) getType() int {
	return opT
}
func (token opToken) getValue() interface{} {
	return token.value
}

func (token brToken) getType() int {
	return brT
}
func (token brToken) getValue() interface{} {
	return token.value
}
