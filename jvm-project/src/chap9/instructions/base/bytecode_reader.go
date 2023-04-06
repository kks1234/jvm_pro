package base

type ByteCodeReader struct {
	code []byte
	pc   int
}

func (bcr *ByteCodeReader) Reset(code []byte, pc int) {
	bcr.code = code
	bcr.pc = pc
}

//实现一系列read方法

func (bcr *ByteCodeReader) ReadUint8() uint8 {
	i := bcr.code[bcr.pc]
	bcr.pc++
	return i
}
func (bcr *ByteCodeReader) PC() int {
	return bcr.pc
}

func (bcr *ByteCodeReader) ReadInt8() int8 {
	return int8(bcr.ReadUint8())
}

func (bcr *ByteCodeReader) ReadUint16() uint16 {
	byte1 := uint16(bcr.ReadUint8())
	byte2 := uint16(bcr.ReadUint8())
	return (byte1 << 8) | byte2
}

func (bcr *ByteCodeReader) ReadInt16() int16 {
	return int16(bcr.ReadUint16())
}

func (bcr *ByteCodeReader) ReadInt32() int32 {
	b1 := int32(bcr.ReadUint8())
	b2 := int32(bcr.ReadUint8())
	b3 := int32(bcr.ReadUint8())
	b4 := int32(bcr.ReadUint8())
	return (b1 << 24) | (b2 << 16) | (b3 << 8) | b4
}

func (bcr *ByteCodeReader) SkipPadding() {
	for bcr.pc%4 != 0 {
		bcr.ReadUint8()
	}
}

func (bcr *ByteCodeReader) ReadInt32s(n int32) []int32 {
	res := make([]int32, n)
	for i := range res {
		res[i] = bcr.ReadInt32()
	}
	return res
}
