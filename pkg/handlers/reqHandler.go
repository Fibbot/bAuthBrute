package handlers

import (
	"encoding/base64"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func MakeDict(users []string, urlInput string, pass []string, method string, success int) []string {

	bigDict := make([]string, 0)
	results := make([]string, 0)

	for i := 0; i < len(users); i++ {
		for j := 0; j < len(pass); j++ {
			word := users[i] + ":" + pass[j]
			word = base64.StdEncoding.EncodeToString([]byte(word))
			bigDict = append(bigDict, word)
		}
	}

	for i := 0; i < len(bigDict); i++ {
		a := makeRequest(bigDict[i], urlInput, method, success)
		if a != "" {
			results = append(results, a)
		}
	}

	return results
}

func makeRequest(userPass string, urlInput string, method string, success int) string {
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
	/*
		bodyText, err := ioutil.ReadAll(resp.Body) // resp. body requires this to print correctly
		b := string(bodyText)
		fmt.Println(b)
		fmt.Println(resp.Status)
	*/

	if resp.StatusCode == success {
		decodeIt, err := base64.StdEncoding.DecodeString(userPass)
		if err != nil {
			log.Fatal(err)
		}
		addToSuccess := "\033[32m" + "Success! " + strconv.Itoa(resp.StatusCode) + " with " + string(decodeIt) + "\033[0m"
		return addToSuccess
	} else {
		return ""
	}

}
