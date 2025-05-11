package ast

func expErr(expression string) error {
	exprLength := len(expression)
	hasNumber := false
	openBrackets := 0
	closeBrackets := 0

	for i := 0; i < exprLength; i++ {
		currentChar := expression[i]
		nextChar := byte(0)
		if i < exprLength-1 {
			nextChar = expression[i+1]
		}

		if currentChar == '(' {
			openBrackets++
		}
		if currentChar == ')' {
			closeBrackets++
		}
		if currentChar >= '0' && currentChar <= '9' && !hasNumber {
			hasNumber = true
		}

		switch {
		case i == 0 && (currentChar == ')' || currentChar == '*' || currentChar == '+' || currentChar == '-' || currentChar == '/'):
			return ErrOperatorFirst
		case i == exprLength-1 && (currentChar == '*' || currentChar == '+' || currentChar == '-' || currentChar == '/'):
			return ErrOperatorLast
		case currentChar == '(' && nextChar == ')':
			return ErrEmptyBrackets
		case currentChar == ')' && nextChar == '(':
			return ErrMergedBrackets
		case (currentChar == '*' || currentChar == '+' || currentChar == '-' || currentChar == '/') && (nextChar == '*' || nextChar == '+' || nextChar == '-' || nextChar == '/'):
			return ErrMergedOperators
		case currentChar < '(' || currentChar > '9':
			return ErrWrongCharacter
		case exprLength <= 2:
			return ErrInvalidExpression
		case currentChar == '/' && nextChar == '0':
			return ErrDivisionByZero
		}
	}

	if openBrackets > closeBrackets {
		return ErrNotClosedBracket
	} else if closeBrackets > openBrackets {
		return ErrNotOpenedBracket
	}

	if !hasNumber {
		return ErrNoOperators
	}
	return nil
}
