package loads

import (
	"jvm-project/src/chap5/instructions/base"
	"jvm-project/src/chap5/rtda"
)

//从局部变量表中读取变量压入操作数栈

type LLOAD struct {
	base.Index8Instruction
}
type LLOAD_0 struct {
	base.NoOperandsInstruction
}

type LLOAD_1 struct {
	base.NoOperandsInstruction
}

type LLOAD_2 struct {
	base.NoOperandsInstruction
}

type LLOAD_3 struct {
	base.NoOperandsInstruction
}

func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

func (L *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, uint(L.Index))
}
func (L *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}

func (L *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

func (L *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

func (L *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}
