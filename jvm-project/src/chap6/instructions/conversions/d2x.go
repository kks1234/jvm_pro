package conversions

import (
	"jvm-project/src/chap6/instructions/base"
	"jvm-project/src/chap6/rtda"
)

type D2F struct {
	base.NoOperandsInstruction
}
type D2I struct {
	base.NoOperandsInstruction
}

type D2L struct {
	base.NoOperandsInstruction
}

func (d *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	result := float32(val)
	stack.PushFloat(result)
}

func (d *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	result := int32(val)
	stack.PushInt(result)
}

func (d *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	result := int64(val)
	stack.PushLong(result)
}
