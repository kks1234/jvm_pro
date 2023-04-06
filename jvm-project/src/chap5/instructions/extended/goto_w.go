package extended

import (
	"jvm-project/src/chap5/instructions/base"
	"jvm-project/src/chap5/rtda"
)

type GOTO_W struct {
	offset int
}

func (G *GOTO_W) FetchOperands(reader *base.ByteCodeReader) {
	G.offset = int(reader.ReadInt32())
}

func (G *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, G.offset)
}

