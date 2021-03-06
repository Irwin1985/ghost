package lexer

import (
	"strings"

	"ghostlang.org/x/ghost/builtins"
	"ghostlang.org/x/ghost/token"
)

// Lexer takes source code as input and outputs the resulting tokens.
type Lexer struct {
	input        string
	position     int
	readPosition int
	character    byte
}

// New creates a new Lexer instance
func New(input string) *Lexer {
	lexer := &Lexer{input: input}

	lexer.readCharacter()

	return lexer
}

func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

// NextToken looks at the current character, and returns
// a token depending on whicharacter character it is.
func (lexer *Lexer) NextToken() token.Token {
	var currentToken token.Token

	lexer.skipWhitespace()

	switch lexer.character {
	case '=':
		if lexer.peekCharacter() == '=' {
			character := lexer.character
			lexer.readCharacter()
			literal := string(character) + string(lexer.character)
			currentToken = token.Token{Type: token.EQ, Literal: literal}
		} else {
			currentToken = newToken(token.ASSIGN, lexer.character)
		}
	case '+':
		if lexer.peekCharacter() == '+' {
			character := lexer.character
			lexer.readCharacter()
			currentToken = token.Token{Type: token.PLUSPLUS, Literal: string(character) + string(lexer.character)}
		} else if lexer.peekCharacter() == '=' {
			character := lexer.character
			lexer.readCharacter()
			currentToken = token.Token{Type: token.PLUSASSIGN, Literal: string(character) + string(lexer.character)}
		} else {
			currentToken = newToken(token.PLUS, lexer.character)
		}
	case '-':
		if lexer.peekCharacter() == '-' {
			character := lexer.character
			lexer.readCharacter()
			currentToken = token.Token{Type: token.MINUSMINUS, Literal: string(character) + string(lexer.character)}
		} else if lexer.peekCharacter() == '=' {
			character := lexer.character
			lexer.readCharacter()
			currentToken = token.Token{Type: token.MINUSASSIGN, Literal: string(character) + string(lexer.character)}
		} else {
			currentToken = newToken(token.MINUS, lexer.character)
		}
	case '!':
		if lexer.peekCharacter() == '=' {
			character := lexer.character
			lexer.readCharacter()
			literal := string(character) + string(lexer.character)
			currentToken = token.Token{Type: token.NOTEQ, Literal: literal}
		} else {
			currentToken = newToken(token.BANG, lexer.character)
		}
	case '#':
		lexer.skipSingleLineComment()

		return lexer.NextToken()
	case '/':
		if lexer.peekCharacter() == '/' {
			lexer.skipSingleLineComment()

			return lexer.NextToken()
		} else if lexer.peekCharacter() == '*' {
			lexer.skipMultiLineComment()

			return lexer.NextToken()
		} else if lexer.peekCharacter() == '=' {
			character := lexer.character
			lexer.readCharacter()
			currentToken = token.Token{Type: token.SLASHASSIGN, Literal: string(character) + string(lexer.character)}
		} else {
			currentToken = newToken(token.SLASH, lexer.character)
		}
	case '*':
		if lexer.peekCharacter() == '=' {
			character := lexer.character
			lexer.readCharacter()
			currentToken = token.Token{Type: token.ASTERISKASSIGN, Literal: string(character) + string(lexer.character)}
		} else {
			currentToken = newToken(token.ASTERISK, lexer.character)
		}
	case '%':
		currentToken = newToken(token.PERCENT, lexer.character)
	case '<':
		if lexer.peekCharacter() == '=' {
			character := lexer.character
			lexer.readCharacter()
			literal := string(character) + string(lexer.character)
			currentToken = token.Token{Type: token.LTE, Literal: literal}
		} else {
			currentToken = newToken(token.LT, lexer.character)
		}
	case '>':
		if lexer.peekCharacter() == '=' {
			character := lexer.character
			lexer.readCharacter()
			literal := string(character) + string(lexer.character)
			currentToken = token.Token{Type: token.GTE, Literal: literal}
		} else {
			currentToken = newToken(token.GT, lexer.character)
		}
	case ';':
		currentToken = newToken(token.SEMICOLON, lexer.character)
	case ',':
		currentToken = newToken(token.COMMA, lexer.character)
	case ':':
		if lexer.peekCharacter() == '=' {
			character := lexer.character
			lexer.readCharacter()
			literal := string(character) + string(lexer.character)
			currentToken = token.Token{Type: token.ASSIGN, Literal: literal}
		} else {
			currentToken = newToken(token.COLON, lexer.character)
		}
	case '(':
		currentToken = newToken(token.LPAREN, lexer.character)
	case ')':
		currentToken = newToken(token.RPAREN, lexer.character)
	case '{':
		currentToken = newToken(token.LBRACE, lexer.character)
	case '}':
		currentToken = newToken(token.RBRACE, lexer.character)
	case '[':
		currentToken = newToken(token.LBRACKET, lexer.character)
	case ']':
		currentToken = newToken(token.RBRACKET, lexer.character)
	case '"':
		currentToken.Type = token.STRING
		currentToken.Literal = lexer.readString('"')
	case '\'':
		currentToken.Type = token.STRING
		currentToken.Literal = lexer.readString('\'')
	case '.':
		if lexer.peekCharacter() == '.' {
			character := lexer.character
			lexer.readCharacter()
			literal := string(character) + string(lexer.character)
			currentToken = token.Token{Type: token.RANGE, Literal: literal}
		} else {
			currentToken = newToken(token.DOT, lexer.character)
		}
	case 0:
		currentToken.Type = token.EOF
		currentToken.Literal = ""
	default:
		if isLetter(lexer.character) {
			currentToken.Literal = lexer.readIdentifier()
			currentToken.Type = token.LookupIdentifier(currentToken.Literal)
			return currentToken
		} else if isDigit(lexer.character) {
			currentToken.Type = token.NUMBER
			currentToken.Literal = lexer.readNumber()
			return currentToken
		} else {
			currentToken = newToken(token.ILLEGAL, lexer.character)
		}
	}

	lexer.readCharacter()

	return currentToken
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.character == ' ' || lexer.character == '\t' || lexer.character == '\n' || lexer.character == '\r' {
		lexer.readCharacter()
	}
}

func (lexer *Lexer) skipSingleLineComment() {
	for lexer.character != '\n' && lexer.character != 0 {
		lexer.readCharacter()
	}

	lexer.skipWhitespace()
}

func (lexer *Lexer) skipMultiLineComment() {
	endOfComment := false

	for !endOfComment {
		if lexer.character == 0 {
			endOfComment = true
		}

		if lexer.character == '*' && lexer.peekCharacter() == '/' {
			endOfComment = true
			lexer.readCharacter()
		}

		lexer.readCharacter()
	}

	lexer.skipWhitespace()
}

func isLetter(character byte) bool {
	return 'a' <= character && character <= 'z' || 'A' <= character && character <= 'Z' || character == '_'
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}

func isIdentifier(character byte) bool {
	return 'a' <= character && character <= 'z' || 'A' <= character && character <= 'Z' || character == '_' || character == '.'
}

func (lexer *Lexer) readString(end byte) string {
	position := lexer.position + 1

	for {
		lexer.readCharacter()
		if lexer.character == end || lexer.character == 0 {
			break
		}
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position

	for isDigit(lexer.character) || lexer.character == '.' {
		lexer.readCharacter()
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	readPosition := lexer.readPosition
	identifier := ""

	for isIdentifier(lexer.character) {
		identifier += string(lexer.character)

		lexer.readCharacter()
	}

	if strings.Contains(identifier, ".") {
		if _, ok := builtins.BuiltinFunctions[identifier]; ok {
			return identifier
		}

		index := strings.Index(identifier, ".")
		identifier = identifier[:index]

		lexer.position = position
		lexer.readPosition = readPosition

		for index > 0 {
			lexer.readCharacter()
			index--
		}
	}

	return identifier
}

func (lexer *Lexer) readCharacter() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.character = 0
	} else {
		lexer.character = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition

	lexer.readPosition++
}

func (lexer *Lexer) peekCharacter() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	}

	return lexer.input[lexer.readPosition]
}
