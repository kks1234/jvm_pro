package comparisons

import (
	"jvm-project/src/chap8/instructions/base"
	"jvm-project/src/chap8/rtda"
)

type DCMPG struct {
	base.NoOperandsInstruction
}

type DCMPL struct {
	base.NoOperandsInstruction
}

func _dcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

func (D *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}

func (D *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}
