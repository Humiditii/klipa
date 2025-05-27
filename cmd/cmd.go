package cmd

import (
	"fmt"
	"os"

	"github.com/Humiditii/klipa/db"
	"github.com/Humiditii/klipa/util"
)

var commandMap = map[string]util.CommandProp{
	"save": {
		Description: "Save clipboard entry or a value from the CLI clipboard",
		Usage:       "To save values copied already, `klipa save`, however if you'd like to save from cli `klipa save -v something -k id`",
	},
	"list": {
		Description: "To get the list of saved items by their key names",
		Usage:       " `klipa list`",
	},
	"get": {
		Description: "Get an item from klipa",
		Usage:       " `klipa get keyName` ",
	},
	"flush": {
		Description: "Clear the all saved items",
		Usage:       " `klipa flush` ",
	},
	"clear": {
		Description: "Delete a particular item from the clipboard",
		Usage:       " `klipa clear keyName` ",
	},
	"copy": {
		Description: "To copy an item from klipa to the device clipboard's memory",
		Usage:       "klipa copy keyName",
	},
}

func Excecute(){

	db.InitDb()

	args := os.Args

	if len(args) == 1 {
		fmt.Println("Please pass in the right command")
		PrintDoc()
		return
	}

	util.CommandSwitcher(&args)

}


func PrintDoc() {
	docs := ``

	for key, prop := range commandMap {
		docs = docs + "\t" +key + "\n \t \t" + "Description:" +prop.Description + ": \t" + "\n \t \tusage: " +prop.Usage + "\n \n"
	}

	fmt.Println(docs)
}