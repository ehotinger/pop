package expr

import (
	"testing"
)

func TestIsIdentifierWithName(t *testing.T) {
	for _, test := range []struct {
		token    *Token
		name     string
		expected bool
	}{
		{
			token:    nil,
			name:     "",
			expected: false,
		},
		{
			token: &Token{
				Type: Identifier,
				Text: "foo",
			},
			name:     "foo",
			expected: true,
		},
		{
			token: &Token{
				Type: StringLiteral,
				Text: "foo",
			},
			name:     "foo",
			expected: false,
		},
		{
			token: &Token{
				Type: Identifier,
				Text: "foo",
			},
			name:     "bar",
			expected: false,
		},
	} {
		if actual := test.token.IsIdentifierWithName(test.name); actual != test.expected {
			t.Fatalf("expected %v but got %v", test.expected, actual)
		}
	}
}
