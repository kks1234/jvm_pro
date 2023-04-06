package constants

import (
	"jvm-project/src/chap10/instructions/base"
	"jvm-project/src/chap10/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (N *NOP) Execute(frame *rtda.Frame) {
}
