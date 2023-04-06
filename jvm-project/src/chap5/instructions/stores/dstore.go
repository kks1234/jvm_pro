package stores

import (
	"jvm-project/src/chap5/instructions/base"
	"jvm-project/src/chap5/rtda"
)

type DSTORE struct {
	base.Index8Instruction
}
type DSTORE_0 struct {
	base.NoOperandsInstruction
}
type DSTORE_1 struct {
	base.NoOperandsInstruction
}
type DSTORE_2 struct {
	base.NoOperandsInstruction
}
type DSTORE_3 struct {
	base.NoOperandsInstruction
}

func _dstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

func (D *DSTORE) Execute(frame *rtda.Frame) {
	_dstore(frame, uint(D.Index))
}

func (D *DSTORE_0) Execute(frame *rtda.Frame) {
	_dstore(frame, 0)
}

func (D *DSTORE_1) Execute(frame *rtda.Frame) {
	_dstore(frame, 1)
}

func (D *DSTORE_2) Execute(frame *rtda.Frame) {
	_dstore(frame, 2)
}

func (D *DSTORE_3) Execute(frame *rtda.Frame) {
	_dstore(frame, 3)
}
