package classpath

import (
	"os"
	"path/filepath"
)

type ClassPath struct {
	bootClassPath Entry
	extClassPath  Entry
	appClassPath  Entry
}

func Parse(jreOption, cpOption string) *ClassPath {
	cp := &ClassPath{}
	cp.parseBootAndExtClassPath(jreOption)
	cp.parseUserClassPath(cpOption)
	return cp
}

func (csp *ClassPath) parseBootAndExtClassPath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	csp.bootClassPath = newWildcardEntry(jreLibPath)
	jreExtpath := filepath.Join(jreDir, "lib", "ext", "*")
	csp.extClassPath = newWildcardEntry(jreExtpath)

}

func (csp *ClassPath) parseUserClassPath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	csp.appClassPath = newEntry(cpOption)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not found jre folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (csp *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := csp.bootClassPath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := csp.extClassPath.readClass(className); err == nil {
		return data, entry, err
	}
	return csp.appClassPath.readClass(className)
}

func (csp *ClassPath) String() string {
	return csp.appClassPath.String()
}
