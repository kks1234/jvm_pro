package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (stack *Stack) Push(frame *Frame) {
	if stack.size >= stack.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if stack._top != nil {
		frame.lower = stack._top
	}
	stack._top = frame
	stack.size++
}

func (stack *Stack) Pop() *Frame {
	if stack._top == nil {
		panic("jvm stack is empty!")
	}
	top := stack._top
	stack._top = stack._top.lower
	top.lower = nil
	return top
}
