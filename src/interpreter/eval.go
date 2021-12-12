package interpreter

import (
	"unirule/grammar"
)

type EvalVisitor struct {
	Data  map[string]bool
	stack []bool
}

func (v *EvalVisitor) VisitAnd(and *grammar.And) {
	v.visit(and.Operands[0])
	v.visit(and.Operands[1])
	op1 := v.stack[len(v.stack)-2]
	op2 := v.stack[len(v.stack)-1]
	v.stack = append(v.stack[:len(v.stack)-2], op1 && op2)

}

func (v *EvalVisitor) VisitOr(or *grammar.Or) {
	v.visit(or.Operands[0])
	v.visit(or.Operands[1])
	op1 := v.stack[len(v.stack)-2]
	op2 := v.stack[len(v.stack)-1]
	v.stack = append(v.stack[:len(v.stack)-2], op1 || op2)
}

func (v *EvalVisitor) VisitNot(not *grammar.Not) {
	v.visit(not.Operand)
	op := v.stack[len(v.stack)-1]
	v.stack = append(v.stack[:len(v.stack)-1], !op)
}

func (v *EvalVisitor) VisitLiteral(l *grammar.Literal) {
	v.stack = append(v.stack, v.Data[l.Value])
}

func (v *EvalVisitor) visit(expr grammar.Expr) {
	switch expr.(type) {
	case *grammar.Or:
		v.VisitOr(expr.(*grammar.Or))
	case *grammar.And:
		v.VisitAnd(expr.(*grammar.And))
	case *grammar.Not:
		v.VisitNot(expr.(*grammar.Not))
	case *grammar.Literal:
		v.VisitLiteral(expr.(*grammar.Literal))
	}
}

func (v *EvalVisitor) Result() bool {
	return v.stack[len(v.stack)-1]
}
