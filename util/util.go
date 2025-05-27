package util

import (
	"fmt"

	"github.com/Humiditii/klipa/db"
	"github.com/atotto/clipboard"
)

type CommandProp struct {
	Description string `json:"description"`
	Usage string `json:"usage"`
}

type CommandT int

const  (
	Save CommandT = iota
	List
	Get
	Flush
	Clear
	Copy
)

func StrToCmdT(s string) CommandT {
	switch s {
	case "save":
		return Save
	case "list":
		return List
	case "get":
		return Get
	case "flush":
		return Flush
	case "clear":
		return Clear
	case "copy":
		return Copy
	default:
		return -1
	}
}



func GetFlagValue(args []string, flag string) (string, bool) {
	for i := 0; i < len(args); i++ {
		if args[i] == "-"+flag && i+1 < len(args) {
			return args[i+1], true
		}
	}
	return "", false
}

func CommandSwitcher(args *[]string){

	invertedCmd := StrToCmdT((*args)[1])
	var err error

	switch invertedCmd {
		case Save:
			v, ok := GetFlagValue(*args, "v")

			if !ok {
				fmt.Println("Will save from clipboard")
				
				v, err = clipboard.ReadAll()

				if err != nil {
					panic(err)
				}

				dto := &db.KlipaModel{
					Value: v,
					Key: db.GetLastIdFromDb(),
				}

				dto.Save()

				return
			}

			k, ok := GetFlagValue(*args, "k")

			if !ok {
				k = db.GetLastIdFromDb()
			}

			dto := &db.KlipaModel{
				Value: v,
				Key: k,
			}
			dto.Save()
			return

		case List:
			keys := db.GetKeys()

			for i :=0; i < len(*keys); i++ {
				fmt.Println((*keys)[i])
			}

			return
		
		case Get:

			fmt.Println( db.GetValueByKey((*args)[2]))

			return

		case Flush:
			db.FlushMem();
			fmt.Println("klipa memory flushed...")
			return
		
		case Clear:
			k := (*args)[2]
			db.DeleteOne(k)
			fmt.Printf("%v deleted from klipa \n", k)
			return
		
		case Copy:
			k := (*args)[2]
			db.CopyFromKlipaToMem(k)
			fmt.Printf(" %v value will be copied to device clipboard \n", k)
			return
	
	}
}