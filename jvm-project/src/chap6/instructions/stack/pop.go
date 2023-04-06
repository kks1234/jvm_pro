package stack

import (
	"jvm-project/src/chap6/instructions/base"
	"jvm-project/src/chap6/rtda"
)

type POP struct {
	base.NoOperandsInstruction
}

type POP2 struct {
	base.NoOperandsInstruction
}

func (P *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

func (P *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
