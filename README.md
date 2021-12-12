# unirule-go

unirule is a logical expression parser & evaluation tool written in Go.

a logical expression is composed of a couple of operands, which should always be a valid boolean value,(either `true` or `false`), as well as logical operator: `&`, `|`, `!`, which concat those operands. pairs of parentheses `(` and `)` are also valid. 

usage:
```Go
expr := grammar.Parse("(1|!2) &3& !4& 5  & !（!6|7|8)")

eval := &interpreter.EvalVisitor{
	Data: map[string]bool{
		"1": true,
		"2": true,
		"3": true,
		"4": false,
		"5": true,
		"6": false,
		"7": false,
		"8": false,
	},
}

expr.Accept(eval)

assert.True(eval.Result() == false)
```
