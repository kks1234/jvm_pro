package reserved

import (
	"jvm-project/src/chap9/instructions/base"
	"jvm-project/src/chap9/native"
	_ "jvm-project/src/chap9/native/java/lang"
	_ "jvm-project/src/chap9/native/sun/misc"
	"jvm-project/src/chap9/rtda"
)

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (I *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
