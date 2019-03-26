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
		break
	case '%':
		t.NextChar()
		tokenType = Percent
		break
	case '&':
		t.NextChar()
		if t.ch == '&' {
			t.NextChar()
			tokenType = DoubleAmpersand
		} else {
			tokenType = Ampersand
		}
		break
	case '(':
		t.NextChar()
		tokenType = OpenParenthesis
		break
	case ')':
		t.NextChar()
		tokenType = CloseParenthesis
		break
	case '*':
		t.NextChar()
		tokenType = Asterisk
		break
	case '+':
		t.NextChar()
		tokenType = Plus
		break
	case '-':
		t.NextChar()
		tokenType = Minus
		break
	case '/':
		t.NextChar()
		tokenType = Slash
		break
	case '<':
		t.NextChar()
		if t.ch == '=' {
			t.NextChar()
			tokenType = LessThanEqual
		} else {
			tokenType = LessThan
		}
		break
	case '=':
		t.NextChar()
		if t.ch == '=' {
			t.NextChar()
			tokenType = DoubleEqual
		} else {
			tokenType = Equal
		}
		break
	case '>':
		t.NextChar()
		if t.ch == '=' {
			t.NextChar()
			tokenType = GreaterThanEqual
		} else {
			tokenType = GreaterThan
		}
		break
	case ',':
		t.NextChar()
		tokenType = Comma
		break
	case '.':
		t.NextChar()
		tokenType = Dot
		break
	case ':':
		t.NextChar()
		tokenType = Colon
		break
	case '?':
		t.NextChar()
		tokenType = Question
		break
	case '[':
		t.NextChar()
		tokenType = OpenBracket
		break
	case ']':
		t.NextChar()
		tokenType = CloseBracket
		break
	case '|':
		t.NextChar()
		if t.ch == '|' {
			t.NextChar()
			tokenType = DoubleBar
		} else {
			tokenType = Bar
		}
		break
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
		break
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
					return token, fmt.Errorf("text position: %d - digit expected", t.position)
				}

				for {
					t.NextChar()
					if !unicode.IsDigit(t.ch) {
						break
					}
				}
			}
			// TODO: support for special floating point syntax?
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
