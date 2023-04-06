package loads

import (
	"jvm-project/src/chap5/instructions/base"
	"jvm-project/src/chap5/rtda"
)

//从局部变量表中读取变量压入操作数栈

type DLOAD struct {
	base.Index8Instruction
}
type DLOAD_0 struct {
	base.NoOperandsInstruction
}

type DLOAD_1 struct {
	base.NoOperandsInstruction
}

type DLOAD_2 struct {
	base.NoOperandsInstruction
}

type DLOAD_3 struct {
	base.NoOperandsInstruction
}

func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

func (D *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, uint(D.Index))
}
func (D *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

func (D *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

func (D *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

func (D *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}
