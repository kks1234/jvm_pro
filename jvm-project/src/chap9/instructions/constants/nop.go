package constants

import (
	"jvm-project/src/chap9/instructions/base"
	"jvm-project/src/chap9/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (N *NOP) Execute(frame *rtda.Frame) {
}
