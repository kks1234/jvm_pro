package classfile

/*
	Deprecated_attribute {
	    u2 attribute_name_index;
	    u4 attribute_length;
	}
*/

/*
	Synthetic_attribute {
	    u2 attribute_name_index;
	    u4 attribute_length;
	}
*/
type MarkerAttribute struct {
}

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

func (ma *MarkerAttribute) readInfo(reader *ClassReader) {
}
