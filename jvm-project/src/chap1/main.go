package main

import "fmt"

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
	fmt.Printf("calsspath: %s ,class: %s, args: %v", cmd.cpOption, cmd.class, cmd.args)
}
