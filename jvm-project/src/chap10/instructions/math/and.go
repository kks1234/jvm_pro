package math

import (
	"jvm-project/src/chap10/instructions/base"
	"jvm-project/src/chap10/rtda"
)

type IAND struct {
	base.NoOperandsInstruction
}

type LAND struct {
	base.NoOperandsInstruction
}

func (I *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

func (L *LAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}
