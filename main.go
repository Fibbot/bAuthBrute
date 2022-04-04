package main

import (
	conc "bAuthBrute/pkg/handlers"
	fileHandler "bAuthBrute/pkg/handlers"
	"flag"
	"fmt"
)

func main() {
	//====================	Handle flags
	fmt.Println("basic auth bruting")
	userList := flag.String("users", "./sampleUsers", "user list")
	passwordList := flag.String("pass", "./small", "password list")
	//====================	Threading TBD - probably need to figure out a way to slow this down
	// threads := flag.Int("t", 1, "threads")
	urlInput := flag.String("u", "http://localhost:8080/test", "target url")
	method := flag.String("m", "POST", "request method")
	//================TODO	probably need a cookie(s) value here too

	flag.Parse()
	// fmt.Println(*threads)
	//====================	Handle Files
	users := fileHandler.ImportFile(*userList)
	pass := fileHandler.ImportFile(*passwordList)
	//====================	Hit 'em up
	conc.SyncScanWg(users, *urlInput, pass, *method)

}
