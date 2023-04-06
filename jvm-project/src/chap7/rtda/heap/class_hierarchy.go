package heap

// jvms8 6.5.instanceof
// jvms8 6.5.checkcast
func (class *Class) isAssignableFrom(other *Class) bool {
	s, t := other, class

	if s == t {
		return true
	}

	if !t.IsInterface() {
		return s.IsSubClassOf(t)
	} else {
		return s.IsImplements(t)
	}
}

func (class *Class) IsSubClassOf(other *Class) bool {
	for c := class.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (class *Class) IsImplements(iface *Class) bool {
	for c := class; c != nil; c = c.superClass {
		for _, inter := range c.interfaces {
			if inter == iface || inter.IsSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

func (class *Class) IsSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range class.interfaces {
		if superInterface == iface || superInterface.IsSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

// other extends self
func (class *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(class)
}
