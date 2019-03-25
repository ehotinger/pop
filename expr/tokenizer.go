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
		Kind: Unknown,
	}

	var tokenKind TokenKind
	tokenPos := t.position
	switch t.ch {
	case '!':
		t.NextChar()
		if t.ch == '=' {
			t.NextChar()
			tokenKind = ExclamationEqual
		} else {
			tokenKind = Exclamation
		}
		break
	case '%':
		t.NextChar()
		tokenKind = Percent
		break
	case '&':
		t.NextChar()
		if t.ch == '&' {
			t.NextChar()
			tokenKind = DoubleAmphersand
		} else {
			tokenKind = Amphersand
		}
		break
	case '(':
		t.NextChar()
		tokenKind = OpenParenthesis
		break
	case ')':
		t.NextChar()
		tokenKind = CloseParenthesis
		break
	case '*':
		t.NextChar()
		tokenKind = Asterisk
		break
	case '+':
		t.NextChar()
		tokenKind = Plus
		break
	case '-':
		t.NextChar()
		tokenKind = Minus
		break
	case '/':
		t.NextChar()
		tokenKind = Slash
		break
	case '<':
		t.NextChar()
		if t.ch == '=' {
			t.NextChar()
			tokenKind = LessThanEqual
		} else {
			tokenKind = LessThan
		}
		break
	case '=':
		t.NextChar()
		if t.ch == '=' {
			t.NextChar()
			tokenKind = DoubleEqual
		} else {
			tokenKind = Equal
		}
		break
	case '>':
		t.NextChar()
		if t.ch == '=' {
			t.NextChar()
			tokenKind = GreaterThanEqual
		} else {
			tokenKind = GreaterThan
		}
		break
	case ',':
		t.NextChar()
		tokenKind = Comma
		break
	case '.':
		t.NextChar()
		tokenKind = Dot
		break
	case ':':
		t.NextChar()
		tokenKind = Colon
		break
	case '?':
		t.NextChar()
		tokenKind = Question
		break
	case '[':
		t.NextChar()
		tokenKind = OpenBracket
		break
	case ']':
		t.NextChar()
		tokenKind = CloseBracket
		break
	case '|':
		t.NextChar()
		if t.ch == '|' {
			t.NextChar()
			tokenKind = DoubleBar
		} else {
			tokenKind = Bar
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
		tokenKind = StringLiteral
		break
	default:
		if unicode.IsLetter(t.ch) || t.ch == '@' || t.ch == '_' {
			for {
				t.NextChar()
				if !unicode.IsLetter(t.ch) && !unicode.IsDigit(t.ch) && t.ch != '_' {
					break
				}
			}
			tokenKind = Identifier
			break
		}
		if unicode.IsDigit(t.ch) {
			tokenKind = IntegerLiteral

			for {
				t.NextChar()
				if !unicode.IsDigit(t.ch) {
					break
				}
			}
			if t.ch == '.' {
				tokenKind = RealLiteral
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
			tokenKind = End
			break
		}
		return token, fmt.Errorf("parsing error at position %d, rune: %v", t.position, t.ch)
	}
	token.Text = string(t.text[tokenPos:t.position])
	token.Position = tokenPos
	token.Kind = tokenKind
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
