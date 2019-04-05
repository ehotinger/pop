package expr

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
