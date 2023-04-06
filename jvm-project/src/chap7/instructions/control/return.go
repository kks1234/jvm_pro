package control

import (
	"jvm-project/src/chap7/instructions/base"
	"jvm-project/src/chap7/rtda"
)

type RETURN struct {
	base.NoOperandsInstruction
}

type ARETURN struct {
	base.NoOperandsInstruction
}

type IRETURN struct {
	base.NoOperandsInstruction
}

type FRETURN struct {
	base.NoOperandsInstruction
}
type DRETURN struct {
	base.NoOperandsInstruction
}
type LRETURN struct {
	base.NoOperandsInstruction
}

func (R *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

func (I *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(retVal)
}
func (A *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	ref := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(ref)
}
func (D *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(val)
}
func (F *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(val)
}
func (L *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(val)
}
