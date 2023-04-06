package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}

func (ci *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	v := reader.readUint32()
	ci.val = int32(v)
}

func (ci *ConstantIntegerInfo) Value() int32 {
	return ci.val
}

type ConstantFloatInfo struct {
	val float32
}

func (ci *ConstantFloatInfo) readInfo(reader *ClassReader) {
	v := reader.readUint32()
	ci.val = math.Float32frombits(v)
}
func (ci *ConstantFloatInfo) Value() float32 {
	return ci.val
}

type ConstantLongInfo struct {
	val int64
}

func (ci *ConstantLongInfo) readInfo(reader *ClassReader) {
	v := reader.readUint64()
	ci.val = int64(v)
}
func (ci *ConstantLongInfo) Value() int64 {
	return ci.val
}

type ConstantDoubleInfo struct {
	val float64
}

func (ci *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	v := reader.readUint64()
	ci.val = math.Float64frombits(v)
}
func (ci *ConstantDoubleInfo) Value() float64 {
	return ci.val
}
