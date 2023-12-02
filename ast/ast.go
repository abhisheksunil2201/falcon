package ast

import "falcon/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node 
	expressionNode()
}

//Root node of every AST our parser produces
type Program struct {
	Statements []Statement
}

//TokenLiteral returns the literal value of the token associated with the root node
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return " "
	}
}

//LetStatement represents a let statement
type LetStatement struct {
	Token token.Token //the token.LET token
	Name  *Identifier //the identifier of the binding
	Value Expression  //the expression to be bound to the identifier
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

//Identifier represents an identifier
type Identifier struct {
	Token token.Token //the token.IDENT token
	Value string      //the value of the identifier
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

//ReturnStatement represents a return statement
type ReturnStatement struct {
	Token token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }