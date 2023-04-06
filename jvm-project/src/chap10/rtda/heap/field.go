package heap

import "jvm-project/src/chap10/classfile"

type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

func (f *Field) IsVolatile() bool {
	return 0 != f.accessFlags&ACC_VOLATILE
}
func (f *Field) IsTransient() bool {
	return 0 != f.accessFlags&ACC_TRANSIENT
}
func (f *Field) IsEnum() bool {
	return 0 != f.accessFlags&ACC_ENUM
}
func (f *Field) isLongOrDouble() bool {
	return f.descriptor == "J" || f.descriptor == "D"
}

func (f *Field) copyAttributes(field *classfile.MemberInfo) {
	if valAttr := field.ConstantValueAttribute(); valAttr != nil {
		f.constValueIndex = uint(valAttr.ConstantValueindex())
	}
}
func (f *Field) ConstValueIndex() uint {
	return f.constValueIndex
}
func (f *Field) SlotId() uint {
	return f.slotId
}
