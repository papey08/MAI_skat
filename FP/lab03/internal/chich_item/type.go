package chich_item

type Type int

const (
	EOF Type = iota
	Error
	Null
	Ident
	Int
	String
	Assign
	Plus
	Minus
	Slash
	Asterisk
	Modulus
	ModulusAssign
	Equals
	NotEquals
	LT
	GT
	LTEQ
	GTEQ
	And
	Or
	Comma
	Semicolon
	NewLine
	LParen
	RParen
	LBrace
	RBrace
	LBracket
	RBracket
	Function
	If
	Else
	True
	False
	Return
)

var typemap = map[Type]string{
	EOF:       "eof",
	Error:     "error",
	Null:      "null",
	Ident:     "IDENT",
	Int:       "int",
	String:    "string",
	Assign:    "=",
	Plus:      "+",
	Minus:     "*",
	Slash:     "/",
	Asterisk:  "*",
	Modulus:   "%",
	Equals:    "==",
	NotEquals: "!=",
	LT:        "<",
	GT:        ">",
	LTEQ:      "<=",
	GTEQ:      ">=",
	And:       "&&",
	Or:        "||",
	Semicolon: ";",
	NewLine:   "new line",
	LParen:    "(",
	RParen:    ")",
	LBrace:    "{",
	RBrace:    "}",
	LBracket:  "[",
	RBracket:  "]",
	Function:  "function",
	If:        "if",
	Else:      "else",
	True:      "true",
	False:     "false",
}

var keywords = map[string]Type{
	"chich":  Function,
	"if":     If,
	"else":   Else,
	"true":   True,
	"false":  False,
	"return": Return,
	"null":   Null,
}

func (t Type) String() string {
	return typemap[t]
}

func Lookup(ident string) Type {
	if t, ok := keywords[ident]; ok {
		return t
	}
	return Ident
}
