package main

import (
	"fmt"
	"jvm-project/src/chap5/classfile"
	"jvm-project/src/chap5/instructions"
	"jvm-project/src/chap5/instructions/base"
	"jvm-project/src/chap5/rtda"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	byteCode := codeAttr.Code()

	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)

	defer catchErr(frame)

	loop(thread, byteCode)
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
