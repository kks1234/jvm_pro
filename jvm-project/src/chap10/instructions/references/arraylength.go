package references

import (
	"jvm-project/src/chap10/instructions/base"
	"jvm-project/src/chap10/rtda"
)

// Get length of array
type ARRAY_LENGTH struct{ base.NoOperandsInstruction }

func (A *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}

	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
