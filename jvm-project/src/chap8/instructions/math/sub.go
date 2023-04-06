package math

import (
	"jvm-project/src/chap8/instructions/base"
	"jvm-project/src/chap8/rtda"
)

type DSUB struct {
	base.NoOperandsInstruction
}

type FSUB struct {
	base.NoOperandsInstruction
}

type ISUB struct {
	base.NoOperandsInstruction
}

type LSUB struct {
	base.NoOperandsInstruction
}

func (D *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 - v2
	stack.PushDouble(result)
}

func (F *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 - v2
	stack.PushFloat(result)
}
func (I *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 - v2
	stack.PushInt(result)
}
func (L *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 - v2
	stack.PushLong(result)
}
