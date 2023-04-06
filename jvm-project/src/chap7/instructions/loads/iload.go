package loads

import (
	"jvm-project/src/chap7/instructions/base"
	"jvm-project/src/chap7/rtda"
)

//从局部变量表中读取变量压入操作数栈

type ILOAD struct {
	base.Index8Instruction
}
type ILOAD_0 struct {
	base.NoOperandsInstruction
}

type ILOAD_1 struct {
	base.NoOperandsInstruction
}

type ILOAD_2 struct {
	base.NoOperandsInstruction
}

type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (I *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(I.Index))
}
func (I *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

func (I *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

func (I *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

func (I *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
