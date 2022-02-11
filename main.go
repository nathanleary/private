package main

import "C"

import (

"fmt"

"syscall"

)

func main() {

var mod = syscall.NewLazyDLL("dll.dll")

var proc = mod.NewProc("PrintBye")

ret, _, _ := proc.Call()

fmt.Printf("Return: %d\n", ret)

}
