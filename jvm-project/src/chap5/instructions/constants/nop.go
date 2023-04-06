package constants

import (
	"jvm-project/src/chap5/instructions/base"
	"jvm-project/src/chap5/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (N *NOP) Execute(frame *rtda.Frame) {
}
