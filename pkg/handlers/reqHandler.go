package handlers

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

func MakeDict(users []string, urlInput string, pass []string, method string, success int, threads int) []string {

	bigDict := make([]string, 0)
	results := make([]string, 0)

	for i := 0; i < len(users); i++ {
		for j := 0; j < len(pass); j++ {
			word := users[i] + ":" + pass[j]
			word = base64.StdEncoding.EncodeToString([]byte(word))
			bigDict = append(bigDict, word)
		}
	}

	semaphore := make(chan struct{}, 5)
	rate := make(chan struct{}, threads)
	ch := make(chan string, 10)
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		for range ticker.C {
			_, ok := <-rate
			// if this isn't going to run indefinitely, signal
			// this to return by closing the rate channel.
			if !ok {
				return
			}
		}
	}()

	time.Sleep(time.Second)
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < len(bigDict); i++ {
		wg.Add(1)
		go func(j int) {
			rate <- struct{}{}
			semaphore <- struct{}{}
			defer func() {
				<-semaphore
			}()

			makeRequest(bigDict[j], urlInput, method, success, ch, &wg)
		}(i)
	}

	wg.Wait()
	close(ch)

	dur := time.Since(start)
	fmt.Printf("sent %d requests in %s\n", len(bigDict), dur)
	for x := range ch {
		if x != "" { //uhhh i think i can remove this now with how there's no return value but idk, working on sth else rn
			results = append(results, x)
		}
	}
	return results
}

func makeRequest(userPass string, urlInput string, method string, success int, ch chan string, wg *sync.WaitGroup) {
	v := url.Values{}
	v.Set("name", "valueOfName")                                                 // add to body
	client := &http.Client{}                                                     // create the request client
	req, err := http.NewRequest(method, urlInput, strings.NewReader(v.Encode())) // request with name:value
	userPassHeader := "Basic " + userPass
	req.Header.Set("Authorization", userPassHeader) // add basic auth header
	resp, err := client.Do(req)                     // send req
	if err != nil {
		log.Fatal(err)
	}

	resp.Body.Close()
	if resp.StatusCode == success {
		decodeIt, err := base64.StdEncoding.DecodeString(userPass)
		if err != nil {
			log.Fatal(err)
		}
		addToSuccess := "\033[32m" + "Success! " + strconv.Itoa(resp.StatusCode) + " with " + string(decodeIt) + "\033[0m"
		ch <- addToSuccess
		wg.Done()
	} else {
		wg.Done()
	}
}
