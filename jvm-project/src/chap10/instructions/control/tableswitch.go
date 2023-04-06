package control

import (
	"jvm-project/src/chap10/instructions/base"
	"jvm-project/src/chap10/rtda"
)

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (T *TABLE_SWITCH) FetchOperands(reader *base.ByteCodeReader) {
	reader.SkipPadding()
	T.defaultOffset = reader.ReadInt32()
	T.low = reader.ReadInt32()
	T.high = reader.ReadInt32()
	jumpOffsetsCount := T.high - T.low + 1
	T.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (T *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= T.low && index <= T.high {
		offset = int(T.jumpOffsets[index-T.low])
	} else {
		offset = int(T.defaultOffset)
	}
	base.Branch(frame, offset)
}
