package references

import (
	"jvm-project/src/chap7/instructions/base"
	"jvm-project/src/chap7/rtda"
	"jvm-project/src/chap7/rtda/heap"
)

type INVOKE_STATIC struct {
	base.Index16Instruction
}

func (I *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(I.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	base.InvokeMethod(frame, resolvedMethod)

}
