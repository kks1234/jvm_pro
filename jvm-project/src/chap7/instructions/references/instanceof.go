package references

import (
	"jvm-project/src/chap7/instructions/base"
	"jvm-project/src/chap7/rtda"
	"jvm-project/src/chap7/rtda/heap"
)

type INSTANCE_OF struct {
	base.Index16Instruction
}

func (I *INSTANCE_OF) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(I.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}

}
