package ast

type stack []*token

func rpn(tokens []*token) ([]*token, error) {
	var tokenStack stack
	result := make([]*token, 0)

	for _, tok := range tokens {
		switch tok.t {
		case operand:
			result = append(result, tok)

		case operator:
			currentPriority, err := priority(tok.val)
			if err != nil {
				return nil, err
			}

			for tokenStack.len() > 0 {
				topToken := tokenStack.peek()
				if topToken.t == openBracket {
					break
				}

				topPriority, err := priority(topToken.val)
				if err != nil {
					return nil, err
				}

				if topPriority >= currentPriority {
					poppedToken, _ := tokenStack.pop()
					result = append(result, poppedToken)
				} else {
					break
				}
			}
			tokenStack.push(tok)

		case openBracket:
			tokenStack.push(tok)

		case closeBracket:
			hasOpenBracket := false
			for tokenStack.len() > 0 {
				poppedToken, err := tokenStack.pop()
				if err != nil {
					return nil, ErrInvalidExpression
				}
				if poppedToken.t == openBracket {
					hasOpenBracket = true
					break
				}
				result = append(result, poppedToken)
			}
			if !hasOpenBracket {
				return nil, ErrNotOpenedBracket
			}

		default:
			return nil, ErrUnknownOperator
		}
	}

	for tokenStack.len() > 0 {
		poppedToken, err := tokenStack.pop()
		if err != nil {
			return nil, err
		}
		if poppedToken.t == openBracket {
			return nil, ErrNotClosedBracket
		}
		result = append(result, poppedToken)
	}

	return result, nil
}

func (s *stack) push(t *token) {
	*s = append(*s, t)
}

func (s *stack) pop() (*token, error) {
	if len(*s) == 0 {
		return nil, ErrEmptyStack
	}
	token := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return token, nil
}

func (s *stack) peek() *token {
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}

func (s *stack) len() int {
	return len(*s)
}
