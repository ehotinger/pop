package expr

// Token represents a single parsed token.
type Token struct {
	Kind     TokenKind
	Value    interface{}
	Text     string
	Position int
}

type TokenKind int

const (
	Unknown TokenKind = iota
	Exclamation
	ExclamationEqual
	Percent
	DoubleAmphersand
	Amphersand
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

func (tk TokenKind) ToString() string {
	switch tk {
	case Exclamation:
		return "Exclamation"
	case ExclamationEqual:
		return "ExclamationEqual"
	case Percent:
		return "Percent"
	case DoubleAmphersand:
		return "DoubleAmphersand"
	case Amphersand:
		return "Amphersand"
	case OpenParenthesis:
		return "OpenParenthesis"
	case CloseParenthesis:
		return "CloseParenthesis"
	case Asterisk:
		return "Asterisk"
	case Plus:
		return "Plus"
	case Minus:
		return "Minus"
	case Slash:
		return "Slash"
	case LessThan:
		return "LessThan"
	case LessThanEqual:
		return "LessThanEqual"
	case Equal:
		return "Equal"
	case DoubleEqual:
		return "DoubleEqual"
	case GreaterThan:
		return "GreaterThan"
	case GreaterThanEqual:
		return "GreaterThanEqual"
	case Bar:
		return "Bar"
	case DoubleBar:
		return "DoubleBar"
	case Comma:
		return "Comma"
	case Dot:
		return "Dot"
	case Colon:
		return "Colon"
	case Question:
		return "Question"
	case OpenBracket:
		return "OpenBracket"
	case CloseBracket:
		return "CloseBracket"
	case Identifier:
		return "Identifier"
	case End:
		return "End"
	case IntegerLiteral:
		return "IntegerLiteral"
	case RealLiteral:
		return "RealLiteral"
	case StringLiteral:
		return "StringLiteral"
	default:
		return "Unknown"
	}
}