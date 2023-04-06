package math

import (
	"jvm-project/src/chap6/instructions/base"
	"jvm-project/src/chap6/rtda"
)

type DNEG struct {
	base.NoOperandsInstruction
}

type FNEG struct {
	base.NoOperandsInstruction
}

type INEG struct {
	base.NoOperandsInstruction
}

type LNEG struct {
	base.NoOperandsInstruction
}

func (D *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

func (F *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

func (I *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

func (L *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}
