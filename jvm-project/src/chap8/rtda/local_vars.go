package rtda

import (
	"jvm-project/src/chap8/rtda/heap"
	"math"
)

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (lv LocalVars) SetInt(index uint, value int32) {
	lv[index].num = value
}

func (lv LocalVars) GetInt(index uint) int32 {
	return lv[index].num
}

func (lv LocalVars) SetFloat(index uint, value float32) {
	bits := math.Float32bits(value)
	lv[index].num = int32(bits)
}
func (lv LocalVars) GetFloat(index uint) float32 {
	bits := uint32(lv[index].num)
	return math.Float32frombits(bits)
}

func (lv LocalVars) SetLong(index uint, value int64) {
	lv[index].num = int32(value)
	lv[index+1].num = int32(value >> 32)
}
func (lv LocalVars) GetLong(index uint) int64 {
	low := uint32(lv[index].num)
	high := uint32(lv[index+1].num)
	return int64(high)<<32 | int64(low)
}

func (lv LocalVars) SetDouble(index uint, value float64) {
	bits := math.Float64bits(value)
	lv.SetLong(index, int64(bits))
}

func (lv LocalVars) GetDouble(index uint) float64 {
	bits := uint64(lv.GetLong(index))
	return math.Float64frombits(bits)
}

func (lv LocalVars) SetRef(index uint, ref *heap.Object) {
	lv[index].ref = ref
}

func (lv LocalVars) GetRef(index uint) *heap.Object {
	return lv[index].ref
}

func (lv LocalVars) SetSlot(index uint, slot Slot) {
	lv[index] = slot
}
