package stores

import (
	"jvm-project/src/chap6/instructions/base"
	"jvm-project/src/chap6/rtda"
)

type ASTORE struct {
	base.Index8Instruction
}
type ASTORE_0 struct {
	base.NoOperandsInstruction
}
type ASTORE_1 struct {
	base.NoOperandsInstruction
}
type ASTORE_2 struct {
	base.NoOperandsInstruction
}
type ASTORE_3 struct {
	base.NoOperandsInstruction
}

func _astore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, val)
}

func (A *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, uint(A.Index))
}

func (A *ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}

func (A *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

func (A *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

func (A *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}
