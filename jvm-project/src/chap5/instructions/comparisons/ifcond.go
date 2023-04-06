package comparisons

import (
	"jvm-project/src/chap5/instructions/base"
	"jvm-project/src/chap5/rtda"
)

type IFEQ struct {
	base.BranchInstruction
}
type IFNE struct {
	base.BranchInstruction
}
type IFLT struct {
	base.BranchInstruction
}
type IFLE struct {
	base.BranchInstruction
}
type IFGT struct {
	base.BranchInstruction
}
type IFGE struct {
	base.BranchInstruction
}

func (I *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, I.Offset)
	}
}

func (I *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, I.Offset)
	}
}
func (I *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, I.Offset)
	}
}
func (I *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, I.Offset)
	}
}
func (I *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, I.Offset)
	}
}
func (I *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, I.Offset)
	}
}
