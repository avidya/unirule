package interpreter

import (
	"unirule/grammar"
)

type EvalVisitor struct {
	Data  map[string]bool
	stack []bool
}

func (v *EvalVisitor) VisitAnd(_ *grammar.And) {
	op1 := v.stack[len(v.stack)-2]
	op2 := v.stack[len(v.stack)-1]
	v.stack = append(v.stack[:len(v.stack)-2], op1 && op2)

}

func (v *EvalVisitor) VisitOr(_ *grammar.Or) {
	op1 := v.stack[len(v.stack)-2]
	op2 := v.stack[len(v.stack)-1]
	v.stack = append(v.stack[:len(v.stack)-2], op1 || op2)
}

func (v *EvalVisitor) VisitNot(_ *grammar.Not) {
	op := v.stack[len(v.stack)-1]
	v.stack = append(v.stack[:len(v.stack)-1], !op)
}

func (v *EvalVisitor) VisitLiteral(l *grammar.Literal) {
	v.stack = append(v.stack, v.Data[l.Value])
}

func (v *EvalVisitor) Result() bool {
	return v.stack[len(v.stack)-1]
}
