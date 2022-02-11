package main

import "C"

import (

"fmt"

"syscall"

)

func main() {

var mod = syscall.NewLazyDLL("dll.dll")

var proc = mod.NewProc("PrintBye")

ptr := proc.Addr()
r1, r2, err := syscall.Syscall(ptr, 0, 0, 0, 0)

fmt.Printf("Return: ", r1, r2, err)

}
