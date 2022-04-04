package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

func SyncScanWg(input []string, urlInput string, pass []string, method string) []string {
	var wg sync.WaitGroup
	var status = make([]string, 0)
	for i := 0; i < len(input); i++ { // change this to be less/gt concurrency?
		wg.Add(1) // Increases internal counter by 1
		go func(j int, uname string) {
			defer wg.Done() // Decrement counter by 1 until done with this stretch of code
			for k := 0; k < len(pass); k++ {
				v := url.Values{}
				v.Set("name", "valueOfName")                                                 // add to body
				client := &http.Client{}                                                     // create the request client
				req, err := http.NewRequest(method, urlInput, strings.NewReader(v.Encode())) // request with name:value
				req.SetBasicAuth(input[j], pass[k])                                          // add basic auth header
				resp, err := client.Do(req)                                                  // send req
				if err != nil {
					log.Fatal(err)
				}
				bodyText, err := ioutil.ReadAll(resp.Body) // resp. body requires this to print correctly
				b := string(bodyText)
				fmt.Println(b)
				fmt.Println(resp.Status)

				if resp.StatusCode == 200 {
					addToSuccess := "Success! " + strconv.Itoa(resp.StatusCode) + " with " + uname + ":" + pass[k]
					status = append(status, addToSuccess) //for now returning it this way in case someone wants to see if there are multiple winners
				}

			}

		}(i, input[i])
	}
	wg.Wait() // Block execution of goroutine until internal counter is 0

	return status
}

func makeRequests() {
	// need an int for threads
	var threads = 1
	startRequests := make(chan string)
	for i := 0; i < threads; i++ {

	}
}
