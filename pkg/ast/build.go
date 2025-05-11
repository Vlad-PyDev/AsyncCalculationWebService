package ast

import (
	"strings"
	"sync"

	"github.com/Vlad-PyDev/AsyncCalculationWebService/internal/models"
)

var (
	mu sync.Mutex
)

func Build(expression string) (*models.AstNode, error) {
	mu.Lock()
	defer mu.Unlock()

	cleanedExpression := strings.ReplaceAll(expression, " ", "")

	if err := expErr(cleanedExpression); err != nil {
		return nil, err
	}

	tokenList := tokens(cleanedExpression)

	rpnTokens, err := rpn(tokenList)
	if err != nil {
		return nil, err
	}

	rootNode, err := ast(rpnTokens)
	if err != nil {
		return nil, err
	}

	return rootNode, nil
}
