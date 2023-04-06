package stores

import (
	"jvm-project/src/chap10/instructions/base"
	"jvm-project/src/chap10/rtda"
)

type ISTORE struct {
	base.Index8Instruction
}
type ISTORE_0 struct {
	base.NoOperandsInstruction
}
type ISTORE_1 struct {
	base.NoOperandsInstruction
}
type ISTORE_2 struct {
	base.NoOperandsInstruction
}
type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func _istore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}

func (I *ISTORE) Execute(frame *rtda.Frame) {
	_istore(frame, uint(I.Index))
}

func (I *ISTORE_0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}

func (I *ISTORE_1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}

func (I *ISTORE_2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}

func (I *ISTORE_3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}
