package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (la *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	la.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range la.lineNumberTable {
		la.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
func (la *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(la.lineNumberTable) - 1; i >= 0; i-- {
		entry := la.lineNumberTable[i]
		if pc >= int(entry.startPc) {
			return int(entry.lineNumber)
		}
	}
	return -1
}
