package expr

import "testing"

// TestNewExpressionParser tests creating a new ExpressionParser and validates
// that the corresponding tokens generated are correct.
func TestNewExpressionParser(t *testing.T) {
	for _, test := range []struct {
		expression     string
		expectedTokens []*Token
	}{
		{
			"1 > 0",
			[]*Token{
				&Token{
					Type: IntegerLiteral,
					Text: "1",
				},
				&Token{
					Type: GreaterThan,
					Text: ">",
				},
				&Token{
					Type: IntegerLiteral,
					Text: "0",
				},
			},
		},
		{
			"! != % & && ( ) * + - / < <= = == > >=",
			[]*Token{
				&Token{
					Type: Exclamation,
					Text: "!",
				},
				&Token{
					Type: ExclamationEqual,
					Text: "!=",
				},
				&Token{
					Type: Percent,
					Text: "%",
				},
				&Token{
					Type: Ampersand,
					Text: "&",
				},
				&Token{
					Type: DoubleAmpersand,
					Text: "&&",
				},
				&Token{
					Type: OpenParenthesis,
					Text: "(",
				},
				&Token{
					Type: CloseParenthesis,
					Text: ")",
				},
				&Token{
					Type: Asterisk,
					Text: "*",
				},
				&Token{
					Type: Plus,
					Text: "+",
				},
				&Token{
					Type: Minus,
					Text: "-",
				},
				&Token{
					Type: Slash,
					Text: "/",
				},
				&Token{
					Type: LessThan,
					Text: "<",
				},
				&Token{
					Type: LessThanEqual,
					Text: "<=",
				},
				&Token{
					Type: Equal,
					Text: "=",
				},
				&Token{
					Type: DoubleEqual,
					Text: "==",
				},
				&Token{
					Type: GreaterThan,
					Text: ">",
				},
				&Token{
					Type: GreaterThanEqual,
					Text: ">=",
				},
			},
		},
		{
			`| || , . : ? [ ] new 3 4.532 'apple' "double"`,
			[]*Token{
				&Token{
					Type: Bar,
					Text: "|",
				},
				&Token{
					Type: DoubleBar,
					Text: "||",
				},
				&Token{
					Type: Comma,
					Text: ",",
				},
				&Token{
					Type: Dot,
					Text: ".",
				},
				&Token{
					Type: Colon,
					Text: ":",
				},
				&Token{
					Type: Question,
					Text: "?",
				},
				&Token{
					Type: OpenBracket,
					Text: "[",
				},
				&Token{
					Type: CloseBracket,
					Text: "]",
				},
				&Token{
					Type: Identifier,
					Text: "new",
				},
				&Token{
					Type: IntegerLiteral,
					Text: "3",
				},
				&Token{
					Type: RealLiteral,
					Text: "4.532",
				},
				&Token{
					Type: StringLiteral,
					Text: "'apple'",
				},
				&Token{
					Type: StringLiteral,
					Text: `"double"`,
				},
			},
		},
	} {
		parser, err := NewExpressionParser(test.expression)
		if err != nil {
			t.Fatalf("unexpected error while creating expression parser: %v", err)
		}

		tokens := parser.tokens
		if len(tokens) != len(test.expectedTokens) {
			t.Fatalf("expected length %d but got %d", len(test.expectedTokens), len(tokens))
		}
		for i := 0; i < len(tokens); i++ {
			if !test.expectedTokens[i].Equals(tokens[i]) {
				t.Fatalf("expected %v but got %v", test.expectedTokens[i], tokens[i])
			}
		}
	}
}
