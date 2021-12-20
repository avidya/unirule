package tests

import (
	"fmt"
	"github.com/avidya/unirule/grammar"
	"github.com/avidya/unirule/interpreter"
	"github.com/stretchr/testify/assert"
	"testing"
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
	assert.New(t).True(eval.Result() == false)
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

func TestToString(t *testing.T) {
	s := grammar.ToString{}
	expr := grammar.Parse("(1|!2) &3& !4& 5  & !（!6|7|8)")
	expr.Accept(&s)
	fmt.Printf("original shape: %s\n", s.String())
	assert.New(t).True(s.String() == "(and (and (and (and (or 1 (not 2)) 3) (not 4)) 5) (not (or (or (not 6) 7) 8)))")
}
