package expr

import "testing"

// TestNewExpressionParser tests creating a new ExpressionParser and validates
// that the corresponding tokens generated are correct.
func TestNewExpressionParser(t *testing.T) {
	for _, test := range []struct {
		expression     string
		expectedTokens []*Token
	}{
		{"1 > 0",
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
