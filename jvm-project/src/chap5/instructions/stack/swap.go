package stack

import (
	"jvm-project/src/chap5/instructions/base"
	"jvm-project/src/chap5/rtda"
)

type SWAP struct {
	base.NoOperandsInstruction
}

func (S *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
