package ast

import (
	"github.com/Vlad-PyDev/AsyncCalculationWebService/internal/models"
)

var (
	id int = 0
)

func ast(tokens []*token) (*models.AstNode, error) {
	var nodeStack []*models.AstNode

	for _, tok := range tokens {
		switch tok.t {
		case operand:
			newNode := &models.AstNode{
				ID:      id,
				AstType: "number",
				Value:   tok.val,
			}
			nodeStack = append(nodeStack, newNode)
			id++

		case operator:
			if len(nodeStack) < 2 {
				return nil, ErrInvalidExpression
			}

			rightNode := nodeStack[len(nodeStack)-1]
			leftNode := nodeStack[len(nodeStack)-2]
			nodeStack = nodeStack[:len(nodeStack)-2]

			newNode := &models.AstNode{
				ID:      id,
				AstType: "operation",
				Value:   tok.val,
				Left:    leftNode,
				Right:   rightNode,
			}
			nodeStack = append(nodeStack, newNode)
			id++

		default:
			return nil, ErrWrongCharacter
		}
	}

	if len(nodeStack) != 1 {
		return nil, ErrInvalidExpression
	}

	return nodeStack[0], nil
}

func priority(op string) (int, error) {
	switch {
	case op == "/" || op == "*":
		return 3, nil
	case op == "+" || op == "-":
		return 2, nil
	case op == "(":
		return 1, nil
	default:
		return 0, ErrUnknownOperator
	}
}
