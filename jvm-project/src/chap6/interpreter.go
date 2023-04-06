package main

import (
	"fmt"
	"jvm-project/src/chap6/instructions"
	"jvm-project/src/chap6/instructions/base"
	"jvm-project/src/chap6/rtda"
	"jvm-project/src/chap6/rtda/heap"
)

func interpret(method *heap.Method) {

	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(frame)

	loop(thread, method.Code())
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars: %v\n", frame.LocalVars())
		fmt.Printf("OperandStack: %v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, byteCode []byte) {
	frame := thread.PopFrame()
	reader := &base.ByteCodeReader{}

	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		//decode
		reader.Reset(byteCode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		//execute
		fmt.Printf("pc: %2d inst: %T %v \n", pc, inst, inst)
		inst.Execute(frame)

	}
}
