package heap

// jvms8 6.5.instanceof
// jvms8 6.5.checkcast
func (class *Class) isAssignableFrom(other *Class) bool {
	s, t := other, class

	if s == t {
		return true
	}

	if !t.IsInterface() {
		return s.isSubClassOf(t)
	} else {
		return s.isImplements(t)
	}
}

func (class *Class) isSubClassOf(other *Class) bool {
	for c := class.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (class *Class) isImplements(iface *Class) bool {
	for c := class; c != nil; c = c.superClass {
		for _, inter := range c.interfaces {
			if inter == iface || inter.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

func (class *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range class.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}
