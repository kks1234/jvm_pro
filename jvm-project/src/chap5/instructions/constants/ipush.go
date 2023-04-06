package constants

import (
	"jvm-project/src/chap5/instructions/base"
	"jvm-project/src/chap5/rtda"
)

type BIPUSH struct {
	val int8
}
type SIPUSH struct {
	val int16
}

func (B *BIPUSH) FetchOperands(reader *base.ByteCodeReader) {
	B.val = reader.ReadInt8()
}

func (B *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(B.val)
	frame.OperandStack().PushInt(i)
}

func (B *SIPUSH) FetchOperands(reader *base.ByteCodeReader) {
	B.val = reader.ReadInt16()
}
func (B *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(B.val)
	frame.OperandStack().PushInt(i)
}
