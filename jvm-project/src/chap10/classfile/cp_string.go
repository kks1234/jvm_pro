package classfile

/*
CONSTANT_String_info {
u1 tag;
u2 string_index;
}
*/

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (ci *ConstantStringInfo) readInfo(reader *ClassReader) {
	ci.stringIndex = reader.readUint16()
}

func (ci *ConstantStringInfo) String() string {
	return ci.cp.getUtf8(ci.stringIndex)
}
