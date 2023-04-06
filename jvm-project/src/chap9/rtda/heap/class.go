package heap

import (
	"jvm-project/src/chap9/classfile"
	"strings"
)

type Class struct {
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initStarted       bool
	jClass            *Object
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (class *Class) IsPublic() bool {
	return 0 != class.accessFlags&ACC_PUBLIC
}
func (class *Class) IsFinal() bool {
	return 0 != class.accessFlags&ACC_FINAL
}
func (class *Class) IsSuper() bool {
	return 0 != class.accessFlags&ACC_SUPER
}
func (class *Class) IsInterface() bool {
	return 0 != class.accessFlags&ACC_INTERFACE
}
func (class *Class) IsAbstract() bool {
	return 0 != class.accessFlags&ACC_ABSTRACT
}
func (class *Class) IsSynthetic() bool {
	return 0 != class.accessFlags&ACC_SYNTHETIC
}
func (class *Class) IsAnnotation() bool {
	return 0 != class.accessFlags&ACC_ANNOTATION
}
func (class *Class) IsEnum() bool {
	return 0 != class.accessFlags&ACC_ENUM
}

func (class *Class) isAccessibleTo(d *Class) bool {
	return class.IsPublic() || class.GetPackageName() == d.GetPackageName()
}

func (class *Class) isJlObject() bool {
	return class.name == "java/lang/Object"
}

func (class *Class) isJlCloneable() bool {
	return class.name == "java/lang/Cloneable"
}
func (class *Class) isJioSerializable() bool {
	return class.name == "java/io/Serializable"
}

func (class *Class) GetPackageName() string {
	if i := strings.LastIndex(class.name, "/"); i >= 0 {
		return class.name[:i]
	}
	return ""
}
func (class *Class) GetMainMethod() *Method {
	return class.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (class *Class) GetClinitMethod() *Method {
	return class.getStaticMethod("<clinit>", "()V")
}

func (class *Class) JavaName() string {
	return strings.Replace(class.name, "/", ".", -1)
}

func (class *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range class.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}
func (class *Class) NewObject() *Object {
	return newObject(class)
}

// getters
func (class *Class) Name() string {
	return class.name
}
func (class *Class) JClass() *Object {
	return class.jClass
}
func (class *Class) Loader() *ClassLoader {
	return class.loader
}
func (class *Class) ConstantPool() *ConstantPool {
	return class.constantPool
}
func (class *Class) Fields() []*Field {
	return class.fields
}
func (class *Class) Methods() []*Method {
	return class.methods
}
func (class *Class) SuperClass() *Class {
	return class.superClass
}
func (class *Class) StaticVars() Slots {
	return class.staticVars
}
func (class *Class) InitStarted() bool {
	return class.initStarted
}
func (class *Class) StartInit() {
	class.initStarted = true
}
func (class *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(class.name)
	return class.loader.LoadClass(arrayClassName)
}
func (class *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := class; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic &&
				field.name == name &&
				field.descriptor == descriptor {

				return field
			}
		}
	}
	return nil
}

func (class *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[class.name]
	return ok
}
func (class *Class) GetRefVar(fieldName, fieldDescriptor string) *Object {
	field := class.getField(fieldName, fieldDescriptor, true)
	return class.staticVars.GetRef(field.slotId)
}
func (class *Class) SetRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := class.getField(fieldName, fieldDescriptor, true)
	class.staticVars.SetRef(field.slotId, ref)
}

func (class *Class) GetInstanceMethod(name, descriptor string) *Method {
	return class.getMethod(name, descriptor, false)
}

func (class *Class) getMethod(name string, descriptor string, isStatic bool) *Method {
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.IsStatic() == isStatic &&
				method.name == name &&
				method.descriptor == descriptor {

				return method
			}
		}
	}
	return nil
}
