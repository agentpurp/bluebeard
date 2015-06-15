package main

//Result: 38
var program1 = []int{
	PUSH, 23,
	PUSH, 20,
	POP,
	PUSH, 15,
	ADD,
	HALT,
}

//Fibonacci series: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 144
var programFib = []int{
	LOAD, a, 0,
	LOAD, b, 1,
	LOAD, d, 233,
	MOV, c, a,
	ADDR, c, a, b,
	MOV, a, b,
	MOV, b, c,
	PRINTR, a,
	CJNE, c, d, 9,
	HALT,
}
