package main

import "C"

import (
	"github.com/sudachen/go-dl/dl"
	"runtime"
	"unsafe"
)

func init() {
    urlbase := "https://github.com/sudachen/go-dl/releases/download/initial/"
    if runtime.GOOS == "linux" && runtime.GOARCH == "amd64"{
        so := dl.Load(
            dl.Cache("dll.so"),
            dl.LzmaExternal(urlbase+"libfunction_lin64.lzma"))
    } else if runtime.GOOS == "windows" && runtime.GOARCH == "amd64" {
        so := dl.Load(
            dl.Cache("dll.dll"),
            dl.LzmaExternal(urlbase+"libfunction_win64.lzma"))
    }
    so.Bind("PrintBye",unsafe.Pointer(&C._godl_function))
}

func main() {
    C.function(0)
}
