package calculator

type iNode interface {
	calculate() float64
}

type node struct {
	left      iNode
	right     iNode
	operation rune
}

type nodeConst struct {
	constVal float64
}

type nodeBrackets struct {
	innerExp iNode
}

func (nc nodeConst) calculate() float64 {
	return nc.constVal
}

func (n nodeBrackets) calculate() float64 {
	return n.innerExp.calculate()
}

func (n node) calculate() float64 {
	switch n.operation {
	case '+':
		return n.left.calculate() + n.right.calculate()
	case '-':
		return n.left.calculate() - n.right.calculate()
	case '*':
		return n.left.calculate() * n.right.calculate()
	}
	return n.left.calculate() / n.right.calculate()
}
