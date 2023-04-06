package loads

import (
	"jvm-project/src/chap7/instructions/base"
	"jvm-project/src/chap7/rtda"
)

//从局部变量表中读取变量压入操作数栈

type FLOAD struct {
	base.Index8Instruction
}
type FLOAD_0 struct {
	base.NoOperandsInstruction
}

type FLOAD_1 struct {
	base.NoOperandsInstruction
}

type FLOAD_2 struct {
	base.NoOperandsInstruction
}

type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

func (F *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, uint(F.Index))
}
func (F *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

func (F *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

func (F *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

func (F *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}
