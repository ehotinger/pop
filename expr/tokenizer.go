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

	token *Token
	ch    rune
}

// NewTokenizer creates a new Tokenizer for the provided expression.
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
	return t, err
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
func (t *Tokenizer) NextToken() (err error) {
	for unicode.IsSpace(t.ch) {
		t.NextChar()
	}

	token := &Token{
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
				return errors.New("unterminated string literal")
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
					return fmt.Errorf("text position: %d - digit expected - char: %s", t.position, string(t.ch))
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
					return fmt.Errorf("text position: %d - digit expected - char: %s", t.position, string(t.ch))
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
		return fmt.Errorf("parsing error at position %d, rune: %v", t.position, t.ch)
	}
	token.Text = string(t.text[tokenPos:t.position])
	token.Position = tokenPos
	token.Type = tokenType
	t.token = token
	return nil
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

func (t *Tokenizer) Parse() (*AbstractExpression, error) {
	err := t.NextToken()
	if err != nil {
		return nil, err
	}
	return t.ParseExpression()
}

// ? : ternary operator
func (t *Tokenizer) ParseExpression() (*AbstractExpression, error) {
	var err error
	expr, err := t.ParseLogicalOr()
	if err != nil {
		return expr, err
	}
	errPos := t.token.Position
	if t.token.Type == Question {
		t.NextToken()
		var expr1 *AbstractExpression
		expr1, err = t.ParseExpression()
		if err != nil {
			return nil, err
		}
		if t.token.Type != Colon {
			return nil, errors.New("expected colon")
		}
		t.NextToken()
		var expr2 *AbstractExpression
		expr2, err = t.ParseExpression()
		if err != nil {
			return nil, err
		}

		expr, err = t.GenerateConditional(expr, expr1, expr2, errPos)
	}
	return expr, err
}

// ||, or
func (t *Tokenizer) ParseLogicalOr() (*AbstractExpression, error) {
	left, err := t.ParseLogicalAnd()
	if err != nil {
		return left, err
	}
	for t.token.Type == DoubleBar || t.token.IsIdentifierWithName(orIdentifier) {
		t.NextToken()
		right, err := t.ParseLogicalAnd()
		if err != nil {
			return right, err
		}
		// TODO: CheckAndPromote for short-circuiting
	}
	return left, nil
}

// &&, and
func (t *Tokenizer) ParseLogicalAnd() (*AbstractExpression, error) {
	left, err := t.ParseComparison()
	if err != nil {
		return nil, err
	}
	for t.token.Type == DoubleAmpersand || t.token.IsIdentifierWithName(andIdentifier) {
		t.NextToken()
		var right *AbstractExpression
		right, err = t.ParseComparison()
		if err != nil {
			return right, err
		}
		// TODO: CheckAndPromote for short-circuiting
	}
	return left, err
}

// =, ==, !=, >, >=, <, <= operators
func (t *Tokenizer) ParseComparison() (*AbstractExpression, error) {
	left, err := t.ParseAdditive()
	if err != nil {
		return left, err
	}
	for t.token.Type == Equal ||
		t.token.Type == DoubleEqual ||
		t.token.Type == ExclamationEqual ||
		t.token.Type == GreaterThan ||
		t.token.Type == GreaterThanEqual ||
		t.token.Type == LessThan ||
		t.token.Type == LessThanEqual {
		// TODO
		t.NextToken()
	}

	return left, err
}

// +, -, &
func (t *Tokenizer) ParseAdditive() (*AbstractExpression, error) {
	left, err := t.ParseMultiplicative()
	if err != nil {
		return left, err
	}
	for t.token.Type == Plus ||
		t.token.Type == Minus ||
		t.token.Type == Ampersand {
		// TODO
		t.NextToken()
	}

	return left, err
}

// *, /, %
func (t *Tokenizer) ParseMultiplicative() (*AbstractExpression, error) {
	left, err := t.ParseUnary()
	if err != nil {
		return left, err
	}
	for t.token.Type == Asterisk ||
		t.token.Type == Slash ||
		t.token.Type == Percent {
		// TODO
		t.NextToken()
	}
	return left, err
}

// -, !
func (t *Tokenizer) ParseUnary() (*AbstractExpression, error) {
	// if t.token.Type == Minus || t.token.Type == Exclamation {
	// 	// TODO
	// }

	return t.ParsePrimary()
}

func (t *Tokenizer) ParsePrimary() (*AbstractExpression, error) {
	expr, err := t.ParsePrimaryStart()
	if err != nil {
		return expr, err
	}
	// TODO
	return expr, nil
}

func (t *Tokenizer) ParsePrimaryStart() (*AbstractExpression, error) {
	switch t.token.Type {
	case Identifier:
		return t.ParseIdentifier()
	case StringLiteral:
		return t.ParseStringLiteral()
	case IntegerLiteral:
		return t.ParseIntegerLiteral()
	case RealLiteral:
		return t.ParseRealLiteral()
	case OpenParenthesis:
		return t.ParseParenthesesExpression()
	default:
		break
	}
	return nil, errors.New("expression expected")
}

func (t *Tokenizer) ParseIdentifier() (*AbstractExpression, error) {
	if t.token.Type != Identifier {
		return nil, fmt.Errorf("expected %v as the token type but got %v", Identifier.ToString(), t.token.Type.ToString())
	}
	return nil, nil
}

func (t *Tokenizer) ParseStringLiteral() (*AbstractExpression, error) {
	if t.token.Type != StringLiteral {
		return nil, fmt.Errorf("expected %v as the token type but got %v", StringLiteral.ToString(), t.token.Type.ToString())
	}
	return nil, nil
}

func (t *Tokenizer) ParseIntegerLiteral() (*AbstractExpression, error) {
	if t.token.Type != IntegerLiteral {
		return nil, fmt.Errorf("expected %v as the token type but got %v", IntegerLiteral.ToString(), t.token.Type.ToString())
	}
	return nil, nil
}

func (t *Tokenizer) ParseRealLiteral() (*AbstractExpression, error) {
	if t.token.Type != RealLiteral {
		return nil, fmt.Errorf("expected %v as the token type but got %v", RealLiteral.ToString(), t.token.Type.ToString())
	}
	return nil, nil
}

func (t *Tokenizer) ParseParenthesesExpression() (*AbstractExpression, error) {
	if t.token.Type != OpenParenthesis {
		return nil, fmt.Errorf("expected %v as the token type but got %v", OpenParenthesis.ToString(), t.token.Type.ToString())
	}
	t.NextToken()
	expr, err := t.ParseExpression()
	if err != nil {
		return expr, err
	}
	if t.token.Type != CloseParenthesis {
		return nil, fmt.Errorf("expected %v as the token type but got %v", CloseParenthesis.ToString(), t.token.Type.ToString())
	}
	t.NextToken()
	return expr, nil
}

func (t *Tokenizer) GenerateConditional(
	expr *AbstractExpression,
	expr1 *AbstractExpression,
	expr2 *AbstractExpression,
	errorPos int) (*AbstractExpression, error) {
	return nil, errors.New("unimplemented")
}
