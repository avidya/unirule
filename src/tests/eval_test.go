package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unirule/grammar"
	"unirule/interpreter"
)

func TestEvalAnd(t *testing.T) {
	assert := assert.New(t)
	expr := grammar.Parse("1 & 2")

	eval := &interpreter.EvalVisitor{
		Data: map[string]bool{
			"1": true,
			"2": false,
		},
	}
	expr.Accept(eval)

	assert.True(!eval.Result())

	eval.Data["2"] = true
	expr.Accept(eval)
	assert.True(eval.Result())
}

func TestEvalReal_1(t *testing.T) {
	expr := grammar.Parse("(1|!2)  &3&！ 4& 5  & !（!6|7|8)")
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
	assert := assert.New(t)
	assert.True(!eval.Result())
}

func TestEvalReal_2(t *testing.T) {
	expr := grammar.Parse("(1|!2)  &3&！ 4& 5  & （!6|7|8)")
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
	assert := assert.New(t)
	assert.True(eval.Result())
}
