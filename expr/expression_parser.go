package expr

import (
	"errors"
)

var (
	errInvalidParenOrder = errors.New("closed parenthesis without an open parenthesis")
	errUnbalancedParen   = errors.New("unbalanced parentheses")
)

type ExpressionParser struct {
	tokenizer *Tokenizer
	tokens    []*Token
}

func NewExpressionParser(expression string) (ep *ExpressionParser, err error) {
	tokenizer, err := NewTokenizer(expression)
	if err != nil {
		return nil, err
	}
	ep = &ExpressionParser{
		tokenizer: tokenizer,
	}
	tokens, err := ep.parseTokens()
	if err != nil {
		return ep, err
	}
	err = ep.validateTokens()
	if err != nil {
		return ep, err
	}

	ep.tokens = tokens
	return ep, err
}

func (ep *ExpressionParser) GetTokens() []*Token {
	return ep.tokens
}

func (ep *ExpressionParser) validateTokens() error {
	paren := 0
	for _, token := range ep.GetTokens() {
		if token.Type == OpenParenthesis {
			paren++
		} else if token.Type == CloseParenthesis {
			paren--
		}
		if paren < 0 {
			return errInvalidParenOrder
		}
	}
	if paren != 0 {
		return errUnbalancedParen
	}
	return nil
}

func (ep *ExpressionParser) parseTokens() ([]*Token, error) {
	var tokens []*Token
	for ep.tokenizer.HasNext() {
		token, err := ep.tokenizer.NextToken()
		if err != nil {
			return tokens, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func (ep *ExpressionParser) Evaluate(parameters map[string]interface{}) (interface{}, error) {
	return "", nil
}
