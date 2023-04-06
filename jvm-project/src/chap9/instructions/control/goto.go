package control

import (
	"jvm-project/src/chap9/instructions/base"
	"jvm-project/src/chap9/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (G *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, G.Offset)
}
