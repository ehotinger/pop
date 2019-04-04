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

func NewExpressionParser(expression string, parameters map[string]interface{}) (ep *ExpressionParser, err error) {
	tokenizer, err := NewTokenizer(expression, parameters)
	if err != nil {
		return nil, err
	}
	ep = &ExpressionParser{
		tokenizer: tokenizer,
	}
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
		err := ep.tokenizer.NextToken()
		if err != nil {
			return tokens, err
		}
		tokens = append(tokens, ep.tokenizer.token)
	}
	return tokens, nil
}

func (ep *ExpressionParser) Evaluate(parameters map[string]interface{}) (interface{}, error) {
	return nil, nil
}

func (ep *ExpressionParser) ParseExpression() (Expression, error) {
	return ep.tokenizer.Parse()
}
