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