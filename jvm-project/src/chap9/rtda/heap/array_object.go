package heap

func (object *Object) Bytes() []int8 {
	return object.data.([]int8)
}
func (object *Object) Shorts() []int16 {
	return object.data.([]int16)
}
func (object *Object) Ints() []int32 {
	return object.data.([]int32)
}
func (object *Object) Longs() []int64 {
	return object.data.([]int64)
}
func (object *Object) Chars() []uint16 {
	return object.data.([]uint16)
}
func (object *Object) Floats() []float32 {
	return object.data.([]float32)
}
func (object *Object) Doubles() []float64 {
	return object.data.([]float64)
}
func (object *Object) Refs() []*Object {
	return object.data.([]*Object)
}

func (object *Object) ArrayLength() int32 {
	switch object.data.(type) {
	case []int8:
		return int32(len(object.data.([]int8)))
	case []int16:
		return int32(len(object.data.([]int16)))
	case []int32:
		return int32(len(object.data.([]int32)))
	case []int64:
		return int32(len(object.data.([]int64)))
	case []uint16:
		return int32(len(object.data.([]uint16)))
	case []float32:
		return int32(len(object.data.([]float32)))
	case []float64:
		return int32(len(object.data.([]float64)))
	case []*Object:
		return int32(len(object.data.([]*Object)))
	default:
		panic("Not array!")
	}
}

func ArrayCopy(src, dst *Object, srcPos, dstPos, length int32) {
	switch src.data.(type) {
	case []int8:
		_src := src.data.([]int8)[srcPos : srcPos+length]
		_dst := dst.data.([]int8)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []int16:
		_src := src.data.([]int16)[srcPos : srcPos+length]
		_dst := dst.data.([]int16)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []int32:
		_src := src.data.([]int32)[srcPos : srcPos+length]
		_dst := dst.data.([]int32)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []int64:
		_src := src.data.([]int64)[srcPos : srcPos+length]
		_dst := dst.data.([]int64)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []uint16:
		_src := src.data.([]uint16)[srcPos : srcPos+length]
		_dst := dst.data.([]uint16)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []float32:
		_src := src.data.([]float32)[srcPos : srcPos+length]
		_dst := dst.data.([]float32)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []float64:
		_src := src.data.([]float64)[srcPos : srcPos+length]
		_dst := dst.data.([]float64)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []*Object:
		_src := src.data.([]*Object)[srcPos : srcPos+length]
		_dst := dst.data.([]*Object)[dstPos : dstPos+length]
		copy(_dst, _src)
	default:
		panic("Not array!")
	}
}
