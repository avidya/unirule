package grammar

import (
	"unicode"
)

type AnalysisContext struct {
	exprChars   []rune
	pos         int
	currentChar rune
	current     Expr
}

func (ctx *AnalysisContext) nextToken() int {
	switch ctx.currentChar {
	case ' ', '　', '\n', '\r', '\t':
		ctx.consume()
		return ctx.nextToken()
	case EOF:
		return EOF
	case '!', '！':
		return EXLAMATION
	case '(', '（':
		return L_PAREN
	case ')', '）':
		return R_PAREN
	case '&':
		return AMP
	case '|', '｜':
		return PIPE
	case '1', '2', '3', '4', '5', '6', '7', '8', '9': // zero can't be put head
		return DIGIT
	default:
		panic("unsupported character: " + string(ctx.currentChar))
	}
}

func (ctx *AnalysisContext) scanDigit() string {
	var l []rune
	for ; unicode.IsDigit(ctx.currentChar); ctx.consume() {
		l = append(l, ctx.currentChar)
	}
	return string(l)
}

func (ctx *AnalysisContext) consume() {
	ctx.pos++
	if ctx.pos >= len(ctx.exprChars) {
		ctx.currentChar = EOF
	} else {
		ctx.currentChar = ctx.exprChars[ctx.pos]
	}
}
