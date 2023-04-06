package heap

import "jvm-project/src/chap10/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (cm *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	cm.accessFlags = memberInfo.AccessFlags()
	cm.name = memberInfo.Name()
	cm.descriptor = memberInfo.Descriptor()
}

func (cm *ClassMember) IsPublic() bool {
	return 0 != cm.accessFlags&ACC_PUBLIC
}
func (cm *ClassMember) IsPrivate() bool {
	return 0 != cm.accessFlags&ACC_PRIVATE
}
func (cm *ClassMember) IsProtected() bool {
	return 0 != cm.accessFlags&ACC_PROTECTED
}
func (cm *ClassMember) IsStatic() bool {
	return 0 != cm.accessFlags&ACC_STATIC
}
func (cm *ClassMember) IsFinal() bool {
	return 0 != cm.accessFlags&ACC_FINAL
}
func (cm *ClassMember) IsSynthetic() bool {
	return 0 != cm.accessFlags&ACC_SYNTHETIC
}

// getters
func (cm *ClassMember) Name() string {
	return cm.name
}
func (cm *ClassMember) Descriptor() string {
	return cm.descriptor
}
func (cm *ClassMember) Class() *Class {
	return cm.class
}

// jvms 5.4.4
func (cm *ClassMember) isAccessibleTo(d *Class) bool {
	if cm.IsPublic() {
		return true
	}
	c := cm.class
	if cm.IsProtected() {
		return d == c || d.IsSubClassOf(c) || c.GetPackageName() == d.GetPackageName()
	}
	if !cm.IsPrivate() {
		return c.GetPackageName() == d.GetPackageName()
	}
	return c == d
}
