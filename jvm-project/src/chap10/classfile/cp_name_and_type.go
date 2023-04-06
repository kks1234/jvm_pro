package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex        uint16
	descriptionIndex uint16
}

func (ci *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	ci.nameIndex = reader.readUint16()
	ci.descriptionIndex = reader.readUint16()
}
