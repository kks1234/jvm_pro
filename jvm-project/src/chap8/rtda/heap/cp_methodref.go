package heap

import "jvm-project/src/chap8/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (mr *MethodRef) ResolvedMethod() *Method {
	if mr.method == nil {
		mr.resolvedMethodRef()
	}
	return mr.method
}

func (mr *MethodRef) resolvedMethodRef() {
	d := mr.cp.class
	c := mr.ResolvedClass()
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupMethod(c, mr.name, mr.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	mr.method = method
}

func lookupMethod(class *Class, name string, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = LookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
