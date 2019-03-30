package expr

import (
	"strings"
)

const (
	orIdentifier  = "or"
	andIdentifier = "and"
	modIdentifier = "mod"
)

// Token represents a single parsed token.
type Token struct {
	Type     TokenType
	Text     string
	Position int
}

func (t *Token) Equals(u *Token) bool {
	return t.Type == u.Type &&
		t.Text == u.Text
}

// IsIdentifierWithName determines whether or not the token
// is an identifier with the provided name.
func (t *Token) IsIdentifierWithName(name string) bool {
	if t == nil {
		return false
	}

	return t.Type == Identifier && strings.EqualFold(t.Text, name)
}

type TokenType int

// TODO: Suffix with Token to avoid package conflict or
// remove suffix from expression type and split package.
const (
	Unknown TokenType = iota
	Exclamation
	ExclamationEqual
	Percent
	DoubleAmpersand
	Ampersand
	OpenParenthesis
	CloseParenthesis
	Asterisk
	Plus
	Minus
	Slash
	LessThan
	LessThanEqual
	Equal
	DoubleEqual
	GreaterThan
	GreaterThanEqual
	Bar
	DoubleBar
	Comma
	Dot
	Colon
	Question
	OpenBracket
	CloseBracket
	Identifier
	End
	IntegerLiteral
	RealLiteral
	StringLiteral
)

const (
	ExclamationString      = "Exclamation"
	ExclamationEqualString = "ExclamationEqual"
	PercentString          = "Percent"
	DoubleAmpersandString  = "DoubleAmpersand"
	AmpersandString        = "Ampersand"
	OpenParenthesisString  = "OpenParenthesis"
	CloseParenthesisString = "CloseParenthesis"
	AsteriskString         = "Asterisk"
	PlusString             = "Plus"
	MinusString            = "Minus"
	SlashString            = "Slash"
	LessThanString         = "LessThan"
	LessThanEqualString    = "LessThanEqual"
	EqualString            = "Equal"
	DoubleEqualString      = "DoubleEqual"
	GreaterThanString      = "GreaterThan"
	GreaterThanEqualString = "GreaterThanEqual"
	BarString              = "Bar"
	DoubleBarString        = "DoubleBar"
	CommaString            = "Comma"
	DotString              = "Dot"
	ColonString            = "Colon"
	QuestionString         = "Question"
	OpenBracketString      = "OpenBracket"
	CloseBracketString     = "CloseBracket"
	IdentifierString       = "Identifier"
	EndString              = "End"
	IntegerLiteralString   = "IntegerLiteral"
	RealLiteralString      = "RealLiteral"
	StringLiteralString    = "StringLiteral"
	UnknownString          = "Unknown"
)

func (t TokenType) ToString() string {
	switch t {
	case Exclamation:
		return ExclamationString
	case ExclamationEqual:
		return ExclamationEqualString
	case Percent:
		return PercentString
	case DoubleAmpersand:
		return DoubleAmpersandString
	case Ampersand:
		return AmpersandString
	case OpenParenthesis:
		return OpenParenthesisString
	case CloseParenthesis:
		return CloseParenthesisString
	case Asterisk:
		return AsteriskString
	case Plus:
		return PlusString
	case Minus:
		return MinusString
	case Slash:
		return SlashString
	case LessThan:
		return LessThanString
	case LessThanEqual:
		return LessThanEqualString
	case Equal:
		return EqualString
	case DoubleEqual:
		return DoubleEqualString
	case GreaterThan:
		return GreaterThanString
	case GreaterThanEqual:
		return GreaterThanEqualString
	case Bar:
		return BarString
	case DoubleBar:
		return DoubleBarString
	case Comma:
		return CommaString
	case Dot:
		return DotString
	case Colon:
		return ColonString
	case Question:
		return QuestionString
	case OpenBracket:
		return OpenBracketString
	case CloseBracket:
		return CloseBracketString
	case Identifier:
		return IdentifierString
	case End:
		return EndString
	case IntegerLiteral:
		return IntegerLiteralString
	case RealLiteral:
		return RealLiteralString
	case StringLiteral:
		return StringLiteralString
	default:
		return UnknownString
	}
}
