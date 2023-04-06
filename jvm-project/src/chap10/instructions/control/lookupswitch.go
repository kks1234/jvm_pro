package control

import (
	"jvm-project/src/chap10/instructions/base"
	"jvm-project/src/chap10/rtda"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (L *LOOKUP_SWITCH) FetchOperands(reader *base.ByteCodeReader) {
	reader.SkipPadding()
	L.defaultOffset = reader.ReadInt32()
	L.npairs = reader.ReadInt32()
	L.matchOffsets = reader.ReadInt32s(L.npairs * 2)
}

func (L *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < L.npairs*2; i += 2 {
		if L.matchOffsets[i] == key {
			offset := L.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(L.defaultOffset))
}
