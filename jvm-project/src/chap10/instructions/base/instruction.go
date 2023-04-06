package base

import "jvm-project/src/chap10/rtda"

type Instruction interface {
	FetchOperands(reader *ByteCodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
}

func (n *NoOperandsInstruction) FetchOperands(reader *ByteCodeReader) {

}

type BranchInstruction struct {
	Offset int
}

func (b *BranchInstruction) FetchOperands(reader *ByteCodeReader) {
	b.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}

func (i *Index8Instruction) FetchOperands(reader *ByteCodeReader) {
	i.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (i *Index16Instruction) FetchOperands(reader *ByteCodeReader) {
	i.Index = uint(reader.ReadUint16())
}
