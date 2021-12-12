package interpreter

import (
	"unirule/grammar"
)

type EvalVisitor struct {
	Data  map[string]bool
	stack []bool
}

func (v *EvalVisitor) VisitAnd(and *grammar.And) {
	and.Operands[0].Accept(v)
	and.Operands[1].Accept(v)
	op1 := v.stack[len(v.stack)-2]
	op2 := v.stack[len(v.stack)-1]
	v.stack = append(v.stack[:len(v.stack)-2], op1 && op2)

}

func (v *EvalVisitor) VisitOr(or *grammar.Or) {
	or.Operands[0].Accept(v)
	or.Operands[1].Accept(v)
	op1 := v.stack[len(v.stack)-2]
	op2 := v.stack[len(v.stack)-1]
	v.stack = append(v.stack[:len(v.stack)-2], op1 || op2)
}

func (v *EvalVisitor) VisitNot(not *grammar.Not) {
	not.Operand.Accept(v)
	op := v.stack[len(v.stack)-1]
	v.stack = append(v.stack[:len(v.stack)-1], !op)
}

func (v *EvalVisitor) VisitLiteral(l *grammar.Literal) {
	v.stack = append(v.stack, v.Data[l.Value])
}

func (v *EvalVisitor) Result() bool {
	return v.stack[len(v.stack)-1]
}
