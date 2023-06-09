package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)

	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := cp[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (cp ConstantPool) getNameAndType(index uint16) (string, string) {
	info := cp.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := cp.getUtf8(info.nameIndex)
	_type := cp.getUtf8(info.descriptionIndex)
	return name, _type
}

func (cp ConstantPool) getClassName(index uint16) string {
	info := cp.getConstantInfo(index).(*ConstantClassInfo)
	return cp.getUtf8(info.nameIndex)
}

func (cp ConstantPool) getUtf8(index uint16) string {
	utf8Info := cp.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
