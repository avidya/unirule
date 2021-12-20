package grammar

import (
	"fmt"
	"strings"
)

type Expr interface {
	Accept(visitor ExprVisitor)
}

type ExprVisitor interface {
	VisitAnd(and *And)
	VisitOr(or *Or)
	VisitNot(not *Not)
	VisitLiteral(l *Literal)
}

type And struct {
	Operands []Expr
}

func (and *And) Accept(visitor ExprVisitor) {
	for _, operand := range and.Operands {
		operand.Accept(visitor)
	}
	visitor.VisitAnd(and)
}

type Or struct {
	Operands []Expr
}

func (or *Or) Accept(visitor ExprVisitor) {
	for _, operand := range or.Operands {
		operand.Accept(visitor)
	}
	visitor.VisitOr(or)
}

type Not struct {
	Operand Expr
}

func (not *Not) Accept(visitor ExprVisitor) {
	not.Operand.Accept(visitor)
	visitor.VisitNot(not)
}

type Literal struct {
	Value string
}

func (l *Literal) Accept(visitor ExprVisitor) {
	visitor.VisitLiteral(l)
}

type ToString struct {
	content []string
}

func (t *ToString) VisitAnd(and *And) {
	orString := strings.Join(t.content[len(t.content)-len(and.Operands):], " ")
	t.content = append(t.content[:len(t.content)-len(and.Operands)], fmt.Sprintf("(and %s)", orString))
}

func (t *ToString) VisitOr(or *Or) {
	orString := strings.Join(t.content[len(t.content)-len(or.Operands):], " ")
	t.content = append(t.content[:len(t.content)-len(or.Operands)], fmt.Sprintf("(or %s)", orString))
}

func (t *ToString) VisitNot(_ *Not) {
	t.content = append(t.content[:len(t.content)-1], fmt.Sprintf("(not %s)", t.content[len(t.content)-1]))
}

func (t *ToString) VisitLiteral(literal *Literal) {
	t.content = append(t.content, literal.Value)
}

func (t *ToString) String() string {
	return t.content[len(t.content)-1]
}