package classfile

/*
CONSTANT_Fieldref_info {
	u1 tag;
	u2 class_index;
	u2 name_and_type_index;
}
*/

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (ci *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	ci.classIndex = reader.readUint16()
	ci.nameAndTypeIndex = reader.readUint16()
}

func (ci *ConstantMemberrefInfo) ClassName() string {
	return ci.cp.getClassName(ci.classIndex)
}

func (ci *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return ci.cp.getNameAndType(ci.nameAndTypeIndex)
}

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}
