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
	visitor.VisitAnd(and)
}

type Or struct {
	Operands []Expr
}

func (or *Or) Accept(visitor ExprVisitor) {
	visitor.VisitOr(or)
}

type Not struct {
	Operand Expr
}

func (not *Not) Accept(visitor ExprVisitor) {
	visitor.VisitNot(not)
}

type Literal struct {
	Value string
}

func (l *Literal) Accept(visitor ExprVisitor) {
	visitor.VisitLiteral(l)
}
