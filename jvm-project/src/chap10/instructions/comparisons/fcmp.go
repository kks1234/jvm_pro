package comparisons

import (
	"jvm-project/src/chap10/instructions/base"
	"jvm-project/src/chap10/rtda"
)

type FCMPG struct {
	base.NoOperandsInstruction
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
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

func (F *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

func (F *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}
