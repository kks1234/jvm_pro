package control

import "jvm-project/src/chap11/instructions/base"
import "jvm-project/src/chap11/rtda"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
