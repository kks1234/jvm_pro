package constants

import (
	"jvm-project/src/chap6/instructions/base"
	"jvm-project/src/chap6/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (N *NOP) Execute(frame *rtda.Frame) {
}
