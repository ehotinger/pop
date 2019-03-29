package expr

import (
	"errors"
	"fmt"
	"unicode"
)

var (
	errMissingExpression = errors.New("expression is required")
)

type Tokenizer struct {
	text     []rune
	position int
	length   int

	ch rune
}

// NewTokenizer creates a new tokenizer for the provided expression.
func NewTokenizer(expression string) (t *Tokenizer, err error) {
	if expression == "" {
		return t, errMissingExpression
	}

	var text []rune
	for _, ch := range expression {
		text = append(text, ch)
	}
	t = &Tokenizer{
		text:   text,
		length: len(text),
	}
	t.SetPosition(0)
	return t, nil
}

func (t *Tokenizer) NextChar() {
	if t.position < t.length {
		t.position++
	}
	if t.position < t.length {
		t.ch = t.text[t.position]
	} else {
		t.ch = '\000'
	}
}

// NextToken returns the next available token and advances the cursor.
func (t *Tokenizer) NextToken() (token *Token, err error) {
	for unicode.IsSpace(t.ch) {
		t.NextChar()
	}

	token = &Token{
		Type: Unknown,
	}

	var tokenType TokenType
	tokenPos := t.position
	switch t.ch {
	case '!':
		t.NextChar()
		if t.ch == '=' {
			t.NextChar()
			tokenType = ExclamationEqual
		} else {
			tokenType = Exclamation
		}
	case '%':
		t.NextChar()
		tokenType = Percent
	case '&':
		t.NextChar()
		if t.ch == '&' {
			t.NextChar()
			tokenType = DoubleAmpersand
		} else {
			tokenType = Ampersand
		}
	case '(':
		t.NextChar()
		tokenType = OpenParenthesis
	case ')':
		t.NextChar()
		tokenType = CloseParenthesis
	case '*':
		t.NextChar()
		tokenType = Asterisk
	case '+':
		t.NextChar()
		tokenType = Plus
	case '-':
		t.NextChar()
		tokenType = Minus
	case '/':
		t.NextChar()
		tokenType = Slash
	case '<':
		t.NextChar()
		if t.ch == '=' {
			t.NextChar()
			tokenType = LessThanEqual
		} else {
			tokenType = LessThan
		}
	case '=':
		t.NextChar()
		if t.ch == '=' {
			t.NextChar()
			tokenType = DoubleEqual
		} else {
			tokenType = Equal
		}
	case '>':
		t.NextChar()
		if t.ch == '=' {
			t.NextChar()
			tokenType = GreaterThanEqual
		} else {
			tokenType = GreaterThan
		}
	case ',':
		t.NextChar()
		tokenType = Comma
	case '.':
		t.NextChar()
		tokenType = Dot
	case ':':
		t.NextChar()
		tokenType = Colon
	case '?':
		t.NextChar()
		tokenType = Question
	case '[':
		t.NextChar()
		tokenType = OpenBracket
	case ']':
		t.NextChar()
		tokenType = CloseBracket
	case '|':
		t.NextChar()
		if t.ch == '|' {
			t.NextChar()
			tokenType = DoubleBar
		} else {
			tokenType = Bar
		}
	case '"':
		fallthrough
	case '\'':
		quote := t.ch
		for {
			t.NextChar()
			for tokenPos < t.length && t.ch != quote {
				t.NextChar()
			}
			if tokenPos == t.length {
				return token, errors.New("unterminated string literal")
			}
			t.NextChar()
			if t.ch != quote {
				break
			}
		}
		tokenType = StringLiteral
	default:
		if unicode.IsLetter(t.ch) || t.ch == '@' || t.ch == '_' {
			for {
				t.NextChar()
				if !unicode.IsLetter(t.ch) && !unicode.IsDigit(t.ch) && t.ch != '_' {
					break
				}
			}
			tokenType = Identifier
			break
		}
		if unicode.IsDigit(t.ch) {
			tokenType = IntegerLiteral

			for {
				t.NextChar()
				if !unicode.IsDigit(t.ch) {
					break
				}
			}
			if t.ch == '.' {
				tokenType = RealLiteral
				t.NextChar()

				if !unicode.IsDigit(t.ch) {
					return token, fmt.Errorf("text position: %d - digit expected - char: %s", t.position, string(t.ch))
				}

				for {
					t.NextChar()
					if !unicode.IsDigit(t.ch) {
						break
					}
				}
			}
			if t.ch == 'E' || t.ch == 'e' {
				tokenType = RealLiteral
				t.NextChar()
				if !unicode.IsDigit(t.ch) {
					return token, fmt.Errorf("text position: %d - digit expected - char: %s", t.position, string(t.ch))
				}
				for {
					t.NextChar()
					if !unicode.IsDigit(t.ch) {
						break
					}
				}
				if t.ch == 'F' || t.ch == 'f' {
					t.NextChar()
				}
			}
			break
		}
		if t.position == t.length {
			tokenType = End
			break
		}
		return token, fmt.Errorf("parsing error at position %d, rune: %v", t.position, t.ch)
	}
	token.Text = string(t.text[tokenPos:t.position])
	token.Position = tokenPos
	token.Type = tokenType
	return token, nil
}

// SetPosition sets the tokenizer's position.
func (t *Tokenizer) SetPosition(position int) {
	t.position = position
	if t.position < t.length {
		t.ch = t.text[t.position]
	} else {
		t.ch = '\000'
	}
}

// HasNext returns true or false if the tokenizer can be advanced.
func (t *Tokenizer) HasNext() bool {
	return t.position < t.length
}
