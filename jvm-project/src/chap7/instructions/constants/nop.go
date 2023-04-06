package constants

import (
	"jvm-project/src/chap7/instructions/base"
	"jvm-project/src/chap7/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (N *NOP) Execute(frame *rtda.Frame) {
}
