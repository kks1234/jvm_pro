package conversions

import (
	"jvm-project/src/chap8/instructions/base"
	"jvm-project/src/chap8/rtda"
)

// int to byte
type I2B struct {
	base.NoOperandsInstruction
}

// int to char
type I2C struct {
	base.NoOperandsInstruction
}

// int to short
type I2S struct {
	base.NoOperandsInstruction
}

// int to long
type I2L struct {
	base.NoOperandsInstruction
}

// int to float
type I2F struct {
	base.NoOperandsInstruction
}

// int to double
type I2D struct {
	base.NoOperandsInstruction
}

func (I *I2B) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	result := int32(int8(val))
	stack.PushInt(result)
}

func (I *I2C) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	result := int32(uint16(val))
	stack.PushInt(result)
}

func (I *I2S) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	result := int32(int16(val))
	stack.PushInt(result)
}

func (I *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	result := int64(val)
	stack.PushLong(result)
}

func (I *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	result := float32(val)
	stack.PushFloat(result)
}

func (I *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	result := float64(val)
	stack.PushDouble(result)
}
