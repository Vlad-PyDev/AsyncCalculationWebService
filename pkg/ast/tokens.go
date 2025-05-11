package ast

import (
	"log"
	"regexp"
)

type token struct {
	t   string
	val string
}

func tokens(str string) []*token {
	tokenList := make([]*token, 0)

	index := 0
	for index < len(str) {
		switch {
		case typeCheck(string(str[index])):
			tokenList = append(tokenList, &token{t: operator, val: string(str[index])})
			index++

		case str[index] >= '0' && str[index] <= '9':
			numberStr := ""
			for index < len(str) && ((str[index] >= '0' && str[index] <= '9') || str[index] == ',' || str[index] == '.') {
				numberStr += string(str[index])
				index++
			}
			tokenList = append(tokenList, &token{t: operand, val: numberStr})

		case str[index] == '(' || str[index] == ')':
			bracketType := openBracket
			if str[index] == ')' {
				bracketType = closeBracket
			}
			tokenList = append(tokenList, &token{t: bracketType, val: string(str[index])})
			index++

		default:
			index++
		}
	}

	return tokenList
}

func typeCheck(symbol string) bool {
	pattern := "[+\\-*/]"
	regex, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatalf("./pkg/ast (func: typeCheck); regexp compilation failed: %s", err)
	}

	return regex.MatchString(symbol)
}
