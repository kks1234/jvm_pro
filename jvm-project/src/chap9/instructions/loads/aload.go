package loads

import (
	"jvm-project/src/chap9/instructions/base"
	"jvm-project/src/chap9/rtda"
)

//从局部变量表中读取变量压入操作数栈

type ALOAD struct {
	base.Index8Instruction
}
type ALOAD_0 struct {
	base.NoOperandsInstruction
}

type ALOAD_1 struct {
	base.NoOperandsInstruction
}

type ALOAD_2 struct {
	base.NoOperandsInstruction
}

type ALOAD_3 struct {
	base.NoOperandsInstruction
}

func _aload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}

func (A *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, uint(A.Index))
}
func (A *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

func (A *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

func (A *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

func (A *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}
