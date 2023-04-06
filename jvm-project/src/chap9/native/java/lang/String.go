package lang

import (
	"jvm-project/src/chap9/native"
	"jvm-project/src/chap9/rtda"
	"jvm-project/src/chap9/rtda/heap"
)

func init() {
	native.Register("java/lang/String", "intern", "()Ljava/lang/String;", intern)
}

func intern(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
