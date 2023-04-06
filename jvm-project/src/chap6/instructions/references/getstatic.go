package references

import (
	"jvm-project/src/chap6/instructions/base"
	"jvm-project/src/chap6/rtda"
	"jvm-project/src/chap6/rtda/heap"
)

type GET_STATIC struct {
	base.Index16Instruction
}

func (G *GET_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(G.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()
	// todo: init class
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	default:
		// todo
	}

}
