package expr

type Expression interface{}

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
	ep.tokens = tokens
	return ep, err
}

func (ep *ExpressionParser) GetTokens() []*Token {
	return ep.tokens
}

func (ep *ExpressionParser) parseTokens() (tokens []*Token, err error) {
	for ep.tokenizer.HasNext() {
		token, err := ep.tokenizer.NextToken()
		if err != nil {
			return tokens, err
		}
		tokens = append(tokens, token)
	}
	return tokens, err
}

func (ep *ExpressionParser) Evaluate(parameters map[string]interface{}) (interface{}, error) {
	return "", nil
}