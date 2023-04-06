package math

import (
	"jvm-project/src/chap8/instructions/base"
	"jvm-project/src/chap8/rtda"
)

type IINC struct {
	Index uint
	Const int32
}

func (I *IINC) FetchOperands(reader *base.ByteCodeReader) {
	I.Index = uint(reader.ReadUint8())
	I.Const = int32(reader.ReadInt8())
}

func (I *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(I.Index)
	val += I.Const
	localVars.SetInt(I.Index, val)
}
