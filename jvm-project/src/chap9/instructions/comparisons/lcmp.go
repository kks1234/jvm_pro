package comparisons

import (
	"jvm-project/src/chap9/instructions/base"
	"jvm-project/src/chap9/rtda"
)

type LCMP struct {
	base.NoOperandsInstruction
}

func (L *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
