package parser

import (
	"falcon/ast"
	"falcon/lexer"
	"falcon/token"
	"fmt"
)

//Parser represents a parser
type Parser struct {
	l *lexer.Lexer
	errors []string
	curToken  token.Token //current token
	peekToken token.Token //next token
}

//New creates a new parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}} //initialize parser
	//read two tokens to set curToken and peekToken
	p.nextToken()
	p.nextToken()
	return p
}

//nextToken reads the next token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

//ParseProgram parses a program
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		//parse statements
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		//read next token
		p.nextToken()
	}
	return program
}

//parseStatement parses a statement
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

//parseLetStatement parses a let statement
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	//check if next token is an identifier
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	//set the identifier
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	//check if next token is an equal sign
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

//parseReturnStatement parses a return statement
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt

}

//curTokenIs checks if the current token is of a given type
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

//peekTokenIs checks if the next token is of a given type
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

//expectPeek checks if the next token is of a given type and advances the tokens if it is
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		//advance tokens
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

//Errors returns the parser errors
func (p *Parser) Errors() []string {
	return p.errors
}

//peekError adds an error to the parser
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}