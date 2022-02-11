package main

import "syscall"
import "unsafe"

var (
  kernel32DLL          = syscall.NewLazyDLL("dll.dll")
  procCreateJobObjectA = kernel32DLL.NewProc("PrintBye")
)

// CreateJobObject uses the CreateJobObjectA Windows API Call to create and return a Handle to a new JobObject
func CreateJobObject(attr *syscall.SecurityAttributes, name string) (syscall.Handle, error) {
  r1, _, err := procCreateJobObjectA.Call()
	
	if err != syscall.Errno(0) {
		return 0, err
	}
	return syscall.Handle(r1), nil
}
