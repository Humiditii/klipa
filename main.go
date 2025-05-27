package main

import (
	"fmt"
	"os"

	"github.com/Humiditii/klipa/cmd"
)


var version = "dev"

func main(){
	cmd.Excecute()

	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Println("klipa version:", version)
		return
	}
}


