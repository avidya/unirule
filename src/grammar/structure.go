package grammar

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
	and.Operands[0].Accept(visitor)
	and.Operands[1].Accept(visitor)
	visitor.VisitAnd(and)
}

type Or struct {
	Operands []Expr
}

func (or *Or) Accept(visitor ExprVisitor) {
	or.Operands[0].Accept(visitor)
	or.Operands[1].Accept(visitor)
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
