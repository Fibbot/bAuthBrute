package main

import (
	fileHandler "bAuthBrute/pkg/handlers"
	reqHandler "bAuthBrute/pkg/handlers"
	"flag"
	"fmt"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func main() {
	//defer timeTrack(time.Now(), "timetest")
	//====================	Handle flags
	fmt.Println("basic auth bruting")
	userList := flag.String("u", "./sampleUsers", "user list")
	passwordList := flag.String("p", "./samplePasswords", "password list")
	urlInput := flag.String("url", "http://localhost:8080/example", "target url")
	method := flag.String("m", "POST", "request method")
	success := flag.Int("s", 200, "success on status code")
	//================TODO	probably need to add cookie value(s) here too if those are needed
	flag.Parse()
	//====================	Handle Files
	users := fileHandler.ImportFile(*userList)
	pass := fileHandler.ImportFile(*passwordList)
	//====================	Hit 'em up
	results := reqHandler.MakeDict(users, *urlInput, pass, *method, *success)
	for i := 0; i < len(results); i++ {
		fmt.Println(results[i])
	}

}
