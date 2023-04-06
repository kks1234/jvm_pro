package constants

import (
	"jvm-project/src/chap6/instructions/base"
	"jvm-project/src/chap6/rtda"
)

// Push item from run-time constant pool
type LDC struct {
	base.Index8Instruction
}

// Push item from run-time constant pool (wide index)
type LDC_W struct {
	base.Index16Instruction
}

// Push long or double from run-time constant pool (wide index)
type LDC2_W struct {
	base.Index16Instruction
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index)

	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	// case string:
	// case *heap.ClassRef:
	// case MethodType, MethodHandle
	default:
		panic("todo: ldc!")
	}
}

func (L *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, L.Index)
}

func (L *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, L.Index)
}

func (L *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(L.Index)

	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}
