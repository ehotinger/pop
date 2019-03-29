package expr

import "testing"

// TestNewExpressionParser tests creating a new ExpressionParser and validates
// that the corresponding tokens generated are correct.
func TestNewExpressionParser(t *testing.T) {
	for _, test := range []struct {
		name           string
		expression     string
		expectedTokens []*Token
	}{
		{
			"arithmetic",
			"1 > 0",
			[]*Token{
				{
					Type: IntegerLiteral,
					Text: "1",
				},
				{
					Type: GreaterThan,
					Text: ">",
				},
				{
					Type: IntegerLiteral,
					Text: "0",
				},
			},
		},
		{
			"symbols",
			"! != % & && ( ) * + - / < <= = == > >=",
			[]*Token{
				{
					Type: Exclamation,
					Text: "!",
				},
				{
					Type: ExclamationEqual,
					Text: "!=",
				},
				{
					Type: Percent,
					Text: "%",
				},
				{
					Type: Ampersand,
					Text: "&",
				},
				{
					Type: DoubleAmpersand,
					Text: "&&",
				},
				{
					Type: OpenParenthesis,
					Text: "(",
				},
				{
					Type: CloseParenthesis,
					Text: ")",
				},
				{
					Type: Asterisk,
					Text: "*",
				},
				{
					Type: Plus,
					Text: "+",
				},
				{
					Type: Minus,
					Text: "-",
				},
				{
					Type: Slash,
					Text: "/",
				},
				{
					Type: LessThan,
					Text: "<",
				},
				{
					Type: LessThanEqual,
					Text: "<=",
				},
				{
					Type: Equal,
					Text: "=",
				},
				{
					Type: DoubleEqual,
					Text: "==",
				},
				{
					Type: GreaterThan,
					Text: ">",
				},
				{
					Type: GreaterThanEqual,
					Text: ">=",
				},
			},
		},
		{
			"double assignment",
			`double x = 1.5E5`,
			[]*Token{
				{
					Type: Identifier,
					Text: "double",
				},
				{
					Type: Identifier,
					Text: "x",
				},
				{
					Type: Equal,
					Text: "=",
				},
				{
					Type: RealLiteral,
					Text: "1.5E5",
				},
			},
		},
		{
			"string literals",
			`| || , . : ? [ ] new 3 4.532 'apple' "double"`,
			[]*Token{
				{
					Type: Bar,
					Text: "|",
				},
				{
					Type: DoubleBar,
					Text: "||",
				},
				{
					Type: Comma,
					Text: ",",
				},
				{
					Type: Dot,
					Text: ".",
				},
				{
					Type: Colon,
					Text: ":",
				},
				{
					Type: Question,
					Text: "?",
				},
				{
					Type: OpenBracket,
					Text: "[",
				},
				{
					Type: CloseBracket,
					Text: "]",
				},
				{
					Type: Identifier,
					Text: "new",
				},
				{
					Type: IntegerLiteral,
					Text: "3",
				},
				{
					Type: RealLiteral,
					Text: "4.532",
				},
				{
					Type: StringLiteral,
					Text: "'apple'",
				},
				{
					Type: StringLiteral,
					Text: `"double"`,
				},
			},
		},
	} {
		parser, err := NewExpressionParser(test.expression)
		if err != nil {
			t.Fatalf("test: %s - unexpected error while creating expression parser: %v", test.name, err)
		}
		tokens, err := parser.parseTokens()
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}

		if len(tokens) != len(test.expectedTokens) {
			t.Fatalf("test: %s - expected length %d but got %d", test.name, len(test.expectedTokens), len(tokens))
		}
		for i := 0; i < len(tokens); i++ {
			if !test.expectedTokens[i].Equals(tokens[i]) {
				t.Fatalf("test: %s - expected %v but got %v", test.name, test.expectedTokens[i], tokens[i])
			}
		}
	}
}

func TestValidateTokens(t *testing.T) {
	for _, test := range []struct {
		parser   *ExpressionParser
		expected error
	}{
		{
			&ExpressionParser{
				tokens: []*Token{
					{
						Type: OpenParenthesis,
						Text: "(",
					},
					{
						Type: OpenParenthesis,
						Text: "(",
					},
					{
						Type: StringLiteral,
						Text: "foo",
					},
					{
						Type: CloseParenthesis,
						Text: ")",
					},
					{
						Type: CloseParenthesis,
						Text: ")",
					},
				},
			},
			nil,
		},
		{
			&ExpressionParser{
				tokens: []*Token{
					{
						Type: OpenParenthesis,
						Text: "(",
					},
					{
						Type: CloseParenthesis,
						Text: ")",
					},
					{
						Type: CloseParenthesis,
						Text: ")",
					},
					{
						Type: StringLiteral,
						Text: "foo",
					},
				},
			},
			errInvalidParenOrder,
		},
		{
			&ExpressionParser{
				tokens: []*Token{
					{
						Type: OpenParenthesis,
						Text: "(",
					},
					{
						Type: OpenParenthesis,
						Text: "(",
					},
					{
						Type: StringLiteral,
						Text: "foo",
					},
					{
						Type: CloseParenthesis,
						Text: ")",
					},
				},
			},
			errUnbalancedParen,
		},
	} {
		if actual := test.parser.validateTokens(); actual != test.expected {
			t.Fatalf("expected %v but got %v", test.expected, actual)
		}
	}
}
