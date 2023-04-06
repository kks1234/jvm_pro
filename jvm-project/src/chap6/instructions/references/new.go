package references

import (
	"jvm-project/src/chap6/instructions/base"
	"jvm-project/src/chap6/rtda"
	"jvm-project/src/chap6/rtda/heap"
)

type NEW struct {
	base.Index16Instruction
}

func (new *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(new.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
