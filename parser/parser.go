package parser

import (
	"genlang/lexer"
	"go/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}
