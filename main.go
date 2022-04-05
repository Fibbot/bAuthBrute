package main

import (
	fileHandler "bAuthBrute/pkg/handlers"
	reqHandler "bAuthBrute/pkg/handlers"
	"flag"
	"fmt"
)

func main() {
	//====================	Handle flags
	fmt.Println("basic auth bruting")
	userList := flag.String("u", "./sampleUsers", "user list")
	passwordList := flag.String("p", "./small", "password list")
	//====================	Threading TBD - probably need to figure out a way to slow this down
	// threads := flag.Int("t", 1, "threads")
	urlInput := flag.String("url", "http://localhost:8080/test", "target url")
	method := flag.String("m", "POST", "request method")
	//================TODO	probably need a cookie(s) value here too
	flag.Parse()
	// fmt.Println(*threads)
	//====================	Handle Files
	users := fileHandler.ImportFile(*userList)
	pass := fileHandler.ImportFile(*passwordList)
	//====================	Hit 'em up
	results := reqHandler.MakeDict(users, *urlInput, pass, *method)
	for i := 0; i < len(results); i++ {
		fmt.Println(results[i])
	}

}
