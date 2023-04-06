package rtda

import "jvm-project/src/chap9/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
