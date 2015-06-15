package main

import "fmt"

func main() {
	fmt.Printf("BLUEBEARD SIMPLE VIRTUAL MACHINE\n")
	var vm = BluebeardVM{
		sp:      0,
		ip:      0,
		running: true,
		program: programFib,
		stack:   make([]int, 256),
		registers: map[int]int{
			a: 0,
			b: 0,
			c: 0,
			d: 0,
		},
	}
	vm.execute()
}
