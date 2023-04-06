package references

import (
	"jvm-project/src/chap9/instructions/base"
	"jvm-project/src/chap9/rtda"
	"jvm-project/src/chap9/rtda/heap"
)

type CHECK_CAST struct {
	base.Index16Instruction
}

func (C *CHECK_CAST) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(C.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
