package native

import "jvm-project/src/chap9/rtda"

type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}

func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	if methodDescriptor == "()V" && methodName == "registerNatives" {
		// todo 用空方法代替各个类中的本地注册方法， 因为自己实现本地方法的注册
		return emptyNativeMethod
	}
	return nil
}

func emptyNativeMethod(frame *rtda.Frame) {
	// do nothing
}
