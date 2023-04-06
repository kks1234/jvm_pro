package classfile

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (cva *ConstantValueAttribute) readInfo(reader *ClassReader) {
	cva.constantValueIndex = reader.readUint16()
}

func (cva *ConstantValueAttribute) ConstantValueindex() uint16 {
	return cva.constantValueIndex
}
