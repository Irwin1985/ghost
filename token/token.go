package token

// TokenType represents the type of token.
type TokenType string

// Token represents the scanned token from source.
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENTIFIER = "IDENTIFIER"
	NUMBER     = "NUMBER"
	STRING     = "STRING"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	PERCENT  = "%"
	DOT      = "."
	HASH     = "#"

	LT  = "<"
	GT  = ">"
	LTE = "<="
	GTE = ">="

	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	FUNCTION = "FUNCTION"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	WHILE    = "WHILE"
	AND      = "AND"
	OR       = "OR"
	EXPORT   = "EXPORT"
	IMPORT   = "IMPORT"

	EQ    = "=="
	NOTEQ = "!="

	BIND           = ":="
	PLUSASSIGN     = "+="
	MINUSASSIGN    = "-="
	ASTERISKASSIGN = "*="
	SLASHASSIGN    = "/="

	PLUSPLUS   = "++"
	MINUSMINUS = "--"
)

var keywords = map[string]TokenType{
	"function": FUNCTION,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
	"while":    WHILE,
	"and":      AND,
	"or":       OR,
	"export":   EXPORT,
	"import":   IMPORT,
}

// LookupIdentifier checks the `keywords` table to see whether
// the given identifier is in fact a keyword.
func LookupIdentifier(identifier string) TokenType {
	if token, ok := keywords[identifier]; ok {
		return token
	}

	return IDENTIFIER
}
