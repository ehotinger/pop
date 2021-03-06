package expr

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

var (
	errMissingExpression     = errors.New("expression is required")
	errMiscomputedExpression = errors.New("unable to create expression")
	errUnimplemented         = errors.New("unimplemented")
)

type Tokenizer struct {
	text     []rune
	position int
	length   int

	parameters map[string]interface{}
	token      *Token
	ch         rune
}

// NewTokenizer creates a new Tokenizer for the provided expression.
func NewTokenizer(expression string, parameters map[string]interface{}) (t *Tokenizer, err error) {
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
	if parameters == nil {
		t.parameters = make(map[string]interface{})
	} else {
		t.parameters = parameters
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

func (t *Tokenizer) Parse() (Expression, error) {
	if err := t.NextToken(); err != nil {
		return nil, err
	}
	return t.ParseExpression()
}

// ? : ternary operator
func (t *Tokenizer) ParseExpression() (Expression, error) {
	var err error
	expr, err := t.ParseLogicalOr()
	if err != nil {
		return expr, err
	}
	errPos := t.token.Position
	if t.token.Type == Question {
		if err = t.NextToken(); err != nil {
			return nil, err
		}
		var expr1 Expression
		expr1, err = t.ParseExpression()
		if err != nil {
			return nil, err
		}
		if t.token.Type != Colon {
			return nil, fmt.Errorf("expected colon, got %v", t.token.Type)
		}
		if err = t.NextToken(); err != nil {
			return nil, err
		}
		var expr2 Expression
		expr2, err = t.ParseExpression()
		if err != nil {
			return nil, err
		}

		expr, err = t.GenerateConditional(expr, expr1, expr2, errPos)
	}
	return expr, err
}

// ||, or
func (t *Tokenizer) ParseLogicalOr() (Expression, error) {
	left, err := t.ParseLogicalAnd()
	if err != nil {
		return left, err
	}
	for t.token.Type == DoubleBar || t.token.IsIdentifierWithName(orIdentifier) {
		if err := t.NextToken(); err != nil {
			return nil, err
		}
		right, err := t.ParseLogicalAnd()
		if err != nil {
			return right, err
		}
		// TODO: CheckAndPromote for short-circuiting
	}
	return left, nil
}

// &&, and
func (t *Tokenizer) ParseLogicalAnd() (Expression, error) {
	left, err := t.ParseComparison()
	if err != nil {
		return nil, err
	}
	for t.token.Type == DoubleAmpersand || t.token.IsIdentifierWithName(andIdentifier) {
		if err = t.NextToken(); err != nil {
			return nil, err
		}
		var right Expression
		right, err = t.ParseComparison()
		if err != nil {
			return right, err
		}
		// TODO: CheckAndPromote for short-circuiting
	}
	return left, err
}

// =, ==, !=, >, >=, <, <= operators
func (t *Tokenizer) ParseComparison() (Expression, error) {
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

		operator := t.token
		if err = t.NextToken(); err != nil {
			return nil, err
		}

		var right Expression
		right, err = t.ParseAdditive()
		if err != nil {
			return nil, err
		}

		// TODO: validation of left and right types
		switch operator.Type {
		case Equal:
			fallthrough
		case DoubleEqual:
			left, err = CreateEqual(left, right)
		case ExclamationEqual:
			left, err = CreateNotEqual(left, right)
		case GreaterThan:
			left, err = CreateGreaterThan(left, right)
		case GreaterThanEqual:
			left, err = CreateGreaterThanOrEqual(left, right)
		case LessThan:
			left, err = CreateLessThan(left, right)
		case LessThanEqual:
			left, err = CreateLessThanOrEqual(left, right)
		}

		if err != nil {
			return left, err
		}
	}
	return left, err
}

// +, -
func (t *Tokenizer) ParseAdditive() (Expression, error) {
	left, err := t.ParseMultiplicative()
	if err != nil {
		return left, err
	}
	for t.token.Type == Plus ||
		t.token.Type == Minus {
		operator := t.token
		if err = t.NextToken(); err != nil {
			return nil, err
		}
		var right Expression
		right, err = t.ParseMultiplicative()
		if err != nil {
			return nil, err
		}
		switch operator.Type {
		case Plus:
			left, err = CreateAdd(left, right)
		case Minus:
			left, err = CreateSubtract(left, right)
		}
		if err != nil {
			return left, err
		}
	}
	return left, err
}

// *, /, %, mod
func (t *Tokenizer) ParseMultiplicative() (Expression, error) {
	left, err := t.ParseUnary()
	if err != nil {
		return left, err
	}
	for t.token.Type == Asterisk ||
		t.token.Type == Slash ||
		t.token.Type == Percent ||
		t.token.IsIdentifierWithName(modIdentifier) {
		operator := t.token
		if err = t.NextToken(); err != nil {
			return nil, err
		}
		var right Expression
		right, err = t.ParseUnary()
		if err != nil {
			return left, err
		}
		// TODO: Promote
		switch operator.Type {
		case Asterisk:
			left, err = CreateMultiply(left, right)
		case Slash:
			left, err = CreateDivide(left, right)
		case Percent:
			fallthrough
		case Identifier:
			left, err = CreateModulus(left, right)
		}
		if err != nil {
			return left, err
		}
	}
	return left, err
}

// -, !, +
func (t *Tokenizer) ParseUnary() (Expression, error) {
	if t.token.Type == Minus || t.token.Type == Exclamation || t.token.Type == Plus {
		operator := t.token
		if err := t.NextToken(); err != nil {
			return nil, err
		}
		if operator.Type == Minus &&
			(t.token.Type == IntegerLiteral || t.token.Type == RealLiteral) {
			t.token.Text = "-" + t.token.Text
			t.token.Position = operator.Position
			return t.ParsePrimary()
		}
		expr, err := t.ParseUnary()
		if err != nil {
			return expr, err
		}
		// TODO: Promote
		if operator.Type == Minus {
			expr, err = CreateUnaryNegate(expr)
		} else if operator.Type == Plus {
			expr, err = CreateUnaryPlus(expr)
		} else {
			expr, err = CreateUnaryNot(expr)
		}
		return expr, err
	}

	return t.ParsePrimary()
}

func (t *Tokenizer) ParsePrimary() (Expression, error) {
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
	return nil, errMiscomputedExpression
}

func (t *Tokenizer) ParseIdentifier() (Expression, error) {
	if t.token.Type != Identifier {
		return nil, fmt.Errorf("expected %v as the token type but got %v", Identifier, t.token.Type)
	}
	text := t.token.Text
	if err := t.NextToken(); err != nil {
		return nil, err
	}
	if val, ok := t.parameters[text]; ok {
		return CreateLiteral(val, text), nil
	}
	return nil, fmt.Errorf("unknown identifier: %s", text)
}

func (t *Tokenizer) ParseStringLiteral() (Expression, error) {
	if t.token.Type != StringLiteral {
		return nil, fmt.Errorf("expected %v as the token type but got %v", StringLiteral, t.token.Type)
	}
	s := strings.Trim(t.token.Text, "'")
	s = strings.Trim(s, `"`)
	if err := t.NextToken(); err != nil {
		return nil, err
	}
	return CreateLiteral(s, s), nil
}

func (t *Tokenizer) ParseIntegerLiteral() (Expression, error) {
	if t.token.Type != IntegerLiteral {
		return nil, fmt.Errorf("expected %v as the token type but got %v", IntegerLiteral, t.token.Type)
	}
	text := t.token.Text
	var value interface{}
	var err error
	if text[0] != '-' {
		// TODO: support parsing as 32, 16, etc.
		value, err = strconv.ParseUint(text, 10, 64)
	} else {
		value, err = strconv.ParseInt(text, 10, 64)
	}
	if err != nil {
		return nil, err
	}
	if err := t.NextToken(); err != nil {
		return nil, err
	}
	return CreateLiteral(value, text), nil
}

func (t *Tokenizer) ParseRealLiteral() (Expression, error) {
	if t.token.Type != RealLiteral {
		return nil, fmt.Errorf("expected %v as the token type but got %v", RealLiteral, t.token.Type)
	}
	var value interface{}
	text := t.token.Text
	f, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return nil, err
	}
	value = f

	if value == nil {
		return nil, errors.New("failed to parse real literal")
	}
	if err := t.NextToken(); err != nil {
		return nil, err
	}

	return CreateLiteral(value, text), nil
}

func (t *Tokenizer) ParseParenthesesExpression() (Expression, error) {
	if t.token.Type != OpenParenthesis {
		return nil, fmt.Errorf("expected %v as the token type but got %v", OpenParenthesis, t.token.Type)
	}
	if err := t.NextToken(); err != nil {
		return nil, err
	}
	expr, err := t.ParseExpression()
	if err != nil {
		return expr, err
	}
	if t.token.Type != CloseParenthesis {
		return nil, fmt.Errorf("expected %v as the token type but got %v", CloseParenthesis, t.token.Type)
	}
	if err := t.NextToken(); err != nil {
		return nil, err
	}
	return expr, nil
}

func (t *Tokenizer) GenerateConditional(
	expr Expression,
	expr1 Expression,
	expr2 Expression,
	errorPos int) (Expression, error) {
	// TODO: impl
	return nil, errUnimplemented
}

// TODO: text is unused for now -- maintain a literals map?
func CreateLiteral(value interface{}, text string) Expression {
	return NewConstantExpression(value, reflect.TypeOf(value).Kind())
}
