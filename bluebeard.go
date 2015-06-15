package main

import "fmt"

//Instructions
const (
	NOP    = iota //NOP
	PUSH          //PUSH STACK
	POP           //POP stack
	LOAD          //LOAD register: LOAD a, 12
	ADD           //ADD stack
	ADDR          //ADD register
	SUB           //SUB stack
	SUBR          //SUB register
	PRINT         //PRINT stack
	PRINTR        //PRINT Registers
	JMP           //Jump
	CJNE          //COMPARE and JUMP if NOT EQUAL
	MOV           //MOVE: MOV, 1,
	HALT          //HALT
)

//Registers
const (
	a = iota
	b
	c
	d
)

type BluebeardVM struct {
	sp, ip    int
	running   bool
	stack     []int
	program   []int
	registers map[int]int
}

//PUSH stack: PUSH, data
func (vm *BluebeardVM) Push() {
	vm.ip++
	vm.stack[vm.sp] = vm.program[vm.ip]
	vm.sp++
}

//POP stack
func (vm *BluebeardVM) Pop() {
	vm.sp--
}

//ADD Stack
func (vm *BluebeardVM) Add() {
	vm.sp--
	var a = vm.stack[vm.sp]
	vm.sp--
	var b = vm.stack[vm.sp]
	vm.stack[vm.sp] = a + b
	vm.sp++
}

//SUB stack
func (vm *BluebeardVM) Sub() {
	vm.sp--
	var a = vm.stack[vm.sp]
	vm.sp--
	var b = vm.stack[vm.sp]
	vm.stack[vm.sp] = a - b
	vm.sp++
}

//ADDRegister: ADDR, registerSave, registerA, registerB
func (vm *BluebeardVM) AddR() {
	regSave := vm.program[vm.ip+1]
	regA := vm.program[vm.ip+2]
	regB := vm.program[vm.ip+3]
	vm.registers[regSave] = vm.registers[regA] + vm.registers[regB]
	vm.ip += 3
}

//Load: LOAD register, data
func (vm *BluebeardVM) Load() {
	reg := vm.program[vm.ip+1]
	load := vm.ip + 2
	vm.registers[reg] = vm.program[load]
	vm.ip = vm.ip + 2
}

//MOVE: MOVE destination, source
func (vm *BluebeardVM) Move() {
	regD := vm.program[vm.ip+1]
	regS := vm.program[vm.ip+2]
	vm.registers[regD] = vm.registers[regS]
	vm.ip += 2
}

//JMP to Instruction: JMP instructionNumber
func (vm *BluebeardVM) Jmp() {
	vm.ip = vm.program[vm.ip+1]
	vm.ip--
}

//Compare and Jump if Not Equal: CJNE registerA, registerB, JMP
func (vm *BluebeardVM) Cjne() {
	regA := vm.program[vm.ip+1]
	regB := vm.program[vm.ip+2]
	if vm.registers[regA] != vm.registers[regB] {
		vm.ip = vm.ip + 2
		vm.Jmp()
		return
	} else {
		vm.ip = vm.ip + 3
	}
}

//PRINTR: PRINTR, register
func (vm *BluebeardVM) PrintR() {
	load := vm.program[vm.ip+1]
	fmt.Printf("Register Value: %d \n", vm.registers[load])
	vm.ip++
}

//Execute Programm
func (vm *BluebeardVM) execute() {
	for vm.running {
		switch vm.program[vm.ip] {
		case NOP:
			//fmt.Printf("NOP\n")
		case PUSH:
			//fmt.Printf("PUSH\n")
			vm.Push()
		case POP:
			//fmt.Printf("POP\n")
			vm.Pop()
		case LOAD:
			//fmt.Printf("LOAD\n")
			vm.Load()
		case ADD:
			//fmt.Printf("ADD\n")
			vm.Add()
		case ADDR:
			//fmt.Printf("ADDR\n")
			vm.AddR()
		case MOV:
			//fmt.Printf("MOVE\n")
			vm.Move()
		case JMP:
			//fmt.Printf("JMP\n")
			vm.Jmp()
		case CJNE:
			//fmt.Printf("CJNE\n")
			vm.Cjne()
		case PRINTR:
			vm.PrintR()
		case HALT:
			//fmt.Printf("HALT\n")
			vm.running = false
		}
		vm.ip++
	}
}
