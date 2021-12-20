package tests

import (
	"fmt"
	"github.com/avidya/unirule/grammar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseReal(t *testing.T) {
	expr := grammar.Parse("(1|!2)  &3&！ 4& 5  & !（!6|7|8)")
	assert := assert.New(t)
	assert.True(expr != nil)

	if op1, ok := expr.(*grammar.And); ok {
		if op2_2, ok := op1.Operands[1].(*grammar.Not); ok {
			if op3_3, ok := op2_2.Operand.(*grammar.Or); ok {
				assert.True(op3_3.Operands[1].(*grammar.Literal).Value == "8")
			} else {
				assert.True(false)
			}
		} else {
			assert.True(false)
		}
	} else {
		assert.True(false)
	}
}

func TestToString(t *testing.T) {
	s := grammar.ToString{}
	expr := grammar.Parse("(1|!2) &3& !4& 5  & !（!6|7|8)")
	expr.Accept(&s)
	fmt.Printf("original shape: %s\n", s.String())
	assert.New(t).True(s.String() == "(and (and (and (and (or 1 (not 2)) 3) (not 4)) 5) (not (or (or (not 6) 7) 8)))")
}

