package expr

// Token represents a single parsed token.
type Token struct {
	Type     TokenType
	Value    interface{}
	Text     string
	Position int
}

func (t *Token) Equals(u *Token) bool {
	return t.Type == u.Type &&
		t.Text == u.Text
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

func (t TokenType) ToString() string {
	switch t {
	case Exclamation:
		return "Exclamation"
	case ExclamationEqual:
		return "ExclamationEqual"
	case Percent:
		return "Percent"
	case DoubleAmpersand:
		return "DoubleAmpersand"
	case Ampersand:
		return "Ampersand"
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
