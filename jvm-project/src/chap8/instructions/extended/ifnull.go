package extended

import (
	"jvm-project/src/chap8/instructions/base"
	"jvm-project/src/chap8/rtda"
)

type IFNULL struct {
	base.BranchInstruction
}

type IFNONNULL struct {
	base.BranchInstruction
}

// Branch if reference is null
func (I *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, I.Offset)
	}
}

// Branch if reference not null
func (I *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, I.Offset)
	}
}
