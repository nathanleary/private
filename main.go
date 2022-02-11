package main

import (
	"fmt"
	"github.com/letiantech/hotplugin"
)

func main() {
	options := hotplugin.ManagerOptions{
		Dir:    "./",
		Suffix: ".dll",
	}
	hotplugin.StartManager(options)
	result := hotplugin.Call("test", "Test", "my world")
	fmt.Println(result...)
}
