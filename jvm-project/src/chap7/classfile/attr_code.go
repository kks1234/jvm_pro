package classfile

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (ca *CodeAttribute) readInfo(reader *ClassReader) {
	ca.maxStack = reader.readUint16()
	ca.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	ca.code = reader.readBytes(codeLength)
	ca.exceptionTable = readExceptionTable(reader)
	ca.attributes = readAttributes(reader, ca.cp)
}

func (ca *CodeAttribute) MaxStack() uint {
	return uint(ca.maxStack)
}
func (ca *CodeAttribute) MaxLocals() uint {
	return uint(ca.maxLocals)
}
func (ca *CodeAttribute) Code() []byte {
	return ca.code
}
func (ca *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return ca.exceptionTable
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}
func (ete *ExceptionTableEntry) StartPc() uint16 {
	return ete.startPc
}
func (ete *ExceptionTableEntry) EndPc() uint16 {
	return ete.endPc
}
func (ete *ExceptionTableEntry) HandlerPc() uint16 {
	return ete.handlerPc
}
func (ete *ExceptionTableEntry) CatchType() uint16 {
	return ete.catchType
}
