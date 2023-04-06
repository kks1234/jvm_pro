package conversions

import (
	"jvm-project/src/chap9/instructions/base"
	"jvm-project/src/chap9/rtda"
)

type F2D struct {
	base.NoOperandsInstruction
}

type F2I struct {
	base.NoOperandsInstruction
}

type F2L struct {
	base.NoOperandsInstruction
}

func (F *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	result := float64(val)
	stack.PushDouble(result)
}

func (F *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	result := int32(val)
	stack.PushInt(result)
}

func (F *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	result := int64(val)
	stack.PushLong(result)
}
