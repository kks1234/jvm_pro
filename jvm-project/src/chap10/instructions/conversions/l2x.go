package conversions

import (
	"jvm-project/src/chap10/instructions/base"
	"jvm-project/src/chap10/rtda"
)

type L2I struct {
	base.NoOperandsInstruction
}

type L2F struct {
	base.NoOperandsInstruction
}

type L2D struct {
	base.NoOperandsInstruction
}

func (L *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	result := int32(val)
	stack.PushInt(result)
}

func (L *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	result := float32(val)
	stack.PushFloat(result)
}

func (L *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	result := float64(val)
	stack.PushDouble(result)
}
