package heap

import "jvm-project/src/chap9/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (imr *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if imr.method == nil {
		imr.ResolvedInterfaceMethodRef()
	}
	return imr.method
}

func (imr *InterfaceMethodRef) ResolvedInterfaceMethodRef() {
	d := imr.cp.class
	c := imr.ResolvedClass()
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupInterfaceMethod(c, imr.name, imr.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	imr.method = method
}

func lookupInterfaceMethod(class *Class, name string, descriptor string) *Method {
	for _, method := range class.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return LookupMethodInInterfaces(class.interfaces, name, descriptor)
}
