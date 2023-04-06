package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPC       int // the next instruction after the call
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}
func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}
func (f *Frame) Thread() *Thread {
	return f.thread
}
func (f *Frame) NextPC() int {
	return f.nextPC
}
func (f *Frame) SetNextPC(nextPC int) {
	f.nextPC = nextPC
}
