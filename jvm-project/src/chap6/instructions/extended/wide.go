package extended

import (
	"jvm-project/src/chap6/instructions/base"
	"jvm-project/src/chap6/instructions/loads"
	"jvm-project/src/chap6/instructions/math"
	"jvm-project/src/chap6/instructions/stores"
	"jvm-project/src/chap6/rtda"
)

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (W *WIDE) FetchOperands(reader *base.ByteCodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15: //iload
		iload := &loads.ILOAD{}
		iload.Index = uint(reader.ReadUint16())
		W.modifiedInstruction = iload
	case 0x16: //lload
		lload := &loads.LLOAD{}
		lload.Index = uint(reader.ReadUint16())
		W.modifiedInstruction = lload
	case 0x17: //fload
		fload := &loads.FLOAD{}
		fload.Index = uint(reader.ReadUint16())
		W.modifiedInstruction = fload
	case 0x18: // dlaod
		dload := &loads.DLOAD{}
		dload.Index = uint(reader.ReadUint16())
		W.modifiedInstruction = dload
	case 0x19: // aload
		aload := &loads.ALOAD{}
		aload.Index = uint(reader.ReadUint16())
		W.modifiedInstruction = aload
	case 0x36: // istore
		istore := &stores.ISTORE{}
		istore.Index = uint(reader.ReadUint16())
		W.modifiedInstruction = istore
	case 0x37: // lstore
		lstore := &stores.LSTORE{}
		lstore.Index = uint(reader.ReadUint16())
		W.modifiedInstruction = lstore
	case 0x38: // fstore
		fstore := &stores.FSTORE{}
		fstore.Index = uint(reader.ReadUint16())
		W.modifiedInstruction = fstore
	case 0x39: // dstore
		dstore := &stores.DSTORE{}
		dstore.Index = uint(reader.ReadUint16())
		W.modifiedInstruction = dstore
	case 0x3a: //astore
		astore := &stores.ASTORE{}
		astore.Index = uint(reader.ReadUint16())
		W.modifiedInstruction = astore
	case 0x84: // iinc
		iinc := &math.IINC{}
		iinc.Index = uint(reader.ReadUint16())
		iinc.Const = int32(reader.ReadInt16())
		W.modifiedInstruction = iinc
	case 0xa9: //ret
		panic("Unsupported opcode: 0xa9!")

	}
}

func (W *WIDE) Execute(frame *rtda.Frame) {
	W.modifiedInstruction.Execute(frame)
}
