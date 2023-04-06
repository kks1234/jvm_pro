package rtda

import "jvm-project/src/chap10/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
