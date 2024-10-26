package calculator

import (
	"fmt"
)

func createNode(tokens []iToken) (iNode, error) {
	if len(tokens) == 0 || len(tokens) == 1 && tokens[0].getType() != numT {
		return nil, fmt.Errorf("createNode: incorrect input format: %v", tokens)
	} else if len(tokens) == 1 { // single number
		return nodeConst{constVal: tokens[0].getValue().(float64)}, nil
	} else if tokens[0].getValue() == '(' && tokens[len(tokens)-1].getValue() == ')' { // the bracket-value, e.g."(a+b*c/d)" => nodeBrackets
		inner, err := createNode(tokens[1 : len(tokens)-1])
		if err != nil {
			return nil, err
		}
		return nodeBrackets{innerExp: inner}, nil
	}
	i := len(tokens) - 1
	for i >= 0 { // low priority operations first (will be calculated last)
		if tokens[i].getValue() == ')' { // multi-value that has brackets (ignoring all in brackets - this will calculate first)
			i--
			var s int
			ok := false
			for i >= 0 {
				if tokens[i].getValue() == ')' {
					s++
				} else if tokens[i].getValue() == '(' && s != 0 {
					s--
				} else if tokens[i].getValue() == '(' && s == 0 {
					if i > 0 {
						i--
					}
					ok = true
					break
				}
				i--
			}
			if !ok {
				return nil, fmt.Errorf("createNode: incorrect input format: %v", tokens)
			} else if i == 0 {
				continue
			}
		}
		if tokens[i].getType() == opT && (tokens[i].getValue() == '+' || tokens[i].getValue() == '-') {
			leftN, errl := createNode(tokens[:i])
			rightN, errr := createNode(tokens[i+1:])
			var er error = nil
			if errl != nil {
				er = errl
			} else if errr != nil {
				er = errr
			}
			if er != nil {
				return nil, er
			}
			return node{left: leftN, right: rightN, operation: tokens[i].getValue().(rune)}, nil
		}
		i--
	}
	// if low priority operations does not exist
	i = len(tokens) - 1
	for i >= 0 { // high priority operations (will be calculated first)
		if tokens[i].getValue() == ')' { // multi-value that has brackets (ignoring all in brackets - this will calculate firts)
			i--
			var s int
			ok := false
			for i >= 0 {
				if tokens[i].getValue() == ')' {
					s++
				} else if tokens[i].getValue() == '(' && s != 0 {
					s--
				} else if tokens[i].getValue() == '(' && s == 0 {
					i--
					ok = true
					break
				}
				i--
			}
			if !ok {
				return nil, fmt.Errorf("createNode: incorrect input format: %v", tokens)
			}
		}
		if tokens[i].getType() == opT && (tokens[i].getValue() == '*' || tokens[i].getValue() == '/') {
			leftN, errl := createNode(tokens[:i])
			rightN, errr := createNode(tokens[i+1:])
			var er error = nil
			if errl != nil {
				er = errl
			} else if errr != nil {
				er = errr
			}
			if er != nil {
				return nil, er
			}
			return node{left: leftN, right: rightN, operation: tokens[i].getValue().(rune)}, nil
		}
		i--
	}
	return nodeConst{constVal: 0}, nil // dummy return
}

func Calc(expression string) (float64, error) {
	tokens, errt := Tokenize(expression)
	if errt != nil {
		return 0, errt
	}
	rootNode, errn := createNode(tokens)
	if errn != nil {
		return 0, errn
	}
	return rootNode.calculate(), nil
}
