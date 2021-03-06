/* Zaki - 6 April 2018 - Initial creation to demonstrate DLL call in Golang */
package main

import (
    "fmt"
    "log"
    "syscall"
)

func main() {
    h, e := syscall.LoadLibrary("dll.dll")   //Make sure this DLL follows Golang machine bit architecture (64-bit in my case)
    if e != nil {
        log.Fatal(e)
    }
    defer syscall.FreeLibrary(h)
    proc, e := syscall.GetProcAddress(h, "PrintBye") //One of the functions in the DLL
    if e != nil {
        log.Fatal(e)
    }
    
    n, _, _ := syscall.Syscall9(uintptr(proc), 0, 2, 2, 2, 2, 0, 0, 0, 0, 0)  //Pay attention to the positioning of the parameter
    fmt.Printf("Hello dll function returns %d\n", n)
}

////////
// package main

// import "C"


// //export PrintBye
// func PrintBye() string {
//    return ("From DLL: Bye!")
// }

// func main() {
//     // Need a main function to make CGO compile package as C shared library
// }
