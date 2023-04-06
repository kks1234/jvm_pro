package stores

import (
	"jvm-project/src/chap5/instructions/base"
	"jvm-project/src/chap5/rtda"
)

type FSTORE struct {
	base.Index8Instruction
}
type FSTORE_0 struct {
	base.NoOperandsInstruction
}
type FSTORE_1 struct {
	base.NoOperandsInstruction
}
type FSTORE_2 struct {
	base.NoOperandsInstruction
}
type FSTORE_3 struct {
	base.NoOperandsInstruction
}

func _fstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

func (F *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, uint(F.Index))
}

func (F *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}

func (F *FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}

func (F *FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}

func (F *FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}
