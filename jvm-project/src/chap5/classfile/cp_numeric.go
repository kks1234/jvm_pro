package classfile

type ConstantIntegerInfo struct {
	val int32
}

func (ci *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	v := reader.readUint32()
	ci.val = int32(v)
}

type ConstantFloatInfo struct {
	val float32
}

func (ci *ConstantFloatInfo) readInfo(reader *ClassReader) {
	v := reader.readUint32()
	ci.val = float32(v)
}

type ConstantLongInfo struct {
	val int64
}

func (ci *ConstantLongInfo) readInfo(reader *ClassReader) {
	v := reader.readUint64()
	ci.val = int64(v)
}

type ConstantDoubleInfo struct {
	val float64
}

func (ci *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	v := reader.readUint64()
	ci.val = float64(v)
}
