package math

import (
	"jvm-project/src/chap7/instructions/base"
	"jvm-project/src/chap7/rtda"
)

type DADD struct {
	base.NoOperandsInstruction
}

type FADD struct {
	base.NoOperandsInstruction
}

type IADD struct {
	base.NoOperandsInstruction
}
type LADD struct {
	base.NoOperandsInstruction
}

func (D *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}

func (F *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}

func (I *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

func (L *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}
