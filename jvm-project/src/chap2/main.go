package main

import (
	"fmt"
	"strings"
)
import "jvm-project/src/chap2/classpath"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJvm(cmd)
	}

}

func startJvm(cmd *Cmd) {
	fmt.Println("start jvm...")
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("calsspath: %v ,class: %v, args: %v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)

	data, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}

	fmt.Printf("class data : %v\n", data)

}
