package ioc

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

type testInterface interface {
	Print()
}

type testInstance struct {
	Test testInterface
}

func (t testInstance) Print() {}

func newTest() testInterface {
	return &testInstance{}
}

func Test_ioc_Set(t *testing.T) {
	t.Run("ok", func(test *testing.T) {
		instType := (*testInterface)(nil)
		inst := newTest()
		Set(
			instType,
			inst,
		)
		rt := getType(inst)
		_, ok := keyOfInstance[rt]
		assert.Equal(test, ok, true)
	})
}
