package grammar

import "fmt"

/**
 * <pre>
 *
 * grammar unirule;
 *
 * expr
 *   : or ( '|' or )* ;
 *
 * or
 *   : and ( '&' and )* ;
 *
 * and
 *   : DIGIT
 *   | '!' and
 *   | '(' expr ')'
 * ;
 *
 * DIGIT
 *   : [1-9][0-9]*
 * ;
 * </pre>
 */

func Parse(exprString string) Expr {
	if len(exprString) == 0 {
		panic("empty input")
	}
	ctx := &AnalysisContext{
		exprChars:   []rune(exprString),
		pos:         0,
		currentChar: rune(exprString[0]),
	}
	ctx.expr()
	return ctx.current
}

func (ctx *AnalysisContext) expr() {
	ctx.or()
	for ctx.nextToken() == PIPE {
		ctx.match(PIPE)
		ex := &Or{
			Operands: []Expr{},
		}
		ex.Operands = append(ex.Operands, ctx.current)
		ctx.or()
		ex.Operands = append(ex.Operands, ctx.current)
		ctx.current = ex
	}
}

func (ctx *AnalysisContext) or() {
	ctx.and()
	for ctx.nextToken() == AMP {
		ctx.match(AMP)
		ex := &And{
			Operands: []Expr{},
		}
		ex.Operands = append(ex.Operands, ctx.current)
		ctx.and()
		ex.Operands = append(ex.Operands, ctx.current)
		ctx.current = ex
	}
}

func (ctx *AnalysisContext) and() {
	switch ctx.nextToken() {
	case DIGIT:
		ctx.current = &Literal{ctx.scanDigit()}
	case EXLAMATION:
		ctx.match(EXLAMATION)
		ex := &Not{}
		ctx.and()
		ex.Operand = ctx.current
		ctx.current = ex
	case L_PAREN:
		ctx.match(L_PAREN)
		ctx.expr()
		ctx.match(R_PAREN)
	default:
		panic("encounter unexpected factor element: " + string(ctx.nextToken()))
	}
}

func (ctx *AnalysisContext) match(token int) {
	if ctx.nextToken() == token {
		ctx.consume()
	} else {
		panic(fmt.Sprintf("wrong token passed in, expected [%d], actual [%d]", token, ctx.nextToken()))
	}
}
