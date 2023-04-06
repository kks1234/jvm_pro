package reserved

import (
	"jvm-project/src/chap10/instructions/base"
	"jvm-project/src/chap10/native"
	_ "jvm-project/src/chap10/native/java/lang"
	_ "jvm-project/src/chap10/native/sun/misc"
	"jvm-project/src/chap10/rtda"
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
