package constants

import "jvm-project/src/chap11/instructions/base"
import "jvm-project/src/chap11/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
