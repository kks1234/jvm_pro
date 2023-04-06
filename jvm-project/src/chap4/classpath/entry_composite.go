package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := CompositeEntry{}

	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (ce CompositeEntry) readClass(clasName string) ([]byte, Entry, error) {
	for _, entry := range ce {
		data, et, err := entry.readClass(clasName)
		if err == nil {
			return data, et, err
		}
	}
	return nil, nil, errors.New("class not found : " + clasName)
}

func (ce CompositeEntry) String() string {
	strs := make([]string, len(ce))
	for i, entry := range ce {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
