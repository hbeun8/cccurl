package main

import (
	//"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	//"strings"
	//"io"
	"encoding/json"
	"maps"
	"net/http"
	"slices"
	"time"
)


func main() {

	flagv := flag.Bool("v", false, "verbose")
	method := flag.String("X", "GET", "HTTP methods")
	postputd := flag.String("d", "", "{'key': 'value'}")
	ct := flag.String("H", "", "Content-Type: application/json")
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.DisableKeepAlives = true
	c := &http.Client{Transport: t}
	flag.Parse()
	url := flag.CommandLine.Args()
	//s := strings.NewReader("HTTP/1.1 200 ok") 
	req, err := http.NewRequest(*method, url[0], nil)
	if err != nil {
			fmt.Println("New Request Error" , err)
	}
	if err!=nil {
		panic(err)
	}
	if !*flagv {
		if *method == "POST" || *method == "PUT" {
			byt1 := []byte(*postputd)
			var dat1 map[string]interface{}
			if err := json.Unmarshal(byt1, &dat1); err != nil {
				panic(err)
			}
			req.Header.Add(slices.Sorted(maps.Keys(dat1))[0], dat1["key"].(string))
			req.Header.Add("data", dat1["key"].(string))
			if len(*ct) > 0 {
				req.Header.Add("Content-Type", "application/json")
			}
		} else {
			if *method == "POST" || *method == "PUT" {
				byt1 := []byte(*postputd)
				var dat1 map[string]interface{}
				if err := json.Unmarshal(byt1, &dat1); err != nil {
					panic(err)
				}
				req.Header.Add("data", dat1["key"].(string))
			} 
			req.Header.Write(os.Stdout)
			req.Header.Add("Host", url[0])
			req.Header.Add("Date", time.Now().String())
			req.Header.Add("Accept", `*/*`)
			req.Header.Add("Connection", "close")
			req.Header.Add("Application", "JSON")	
		}
		fmt.Printf("Sending request %v /get HTTP/1.1\n", *method)
		fmt.Println("Host: ", url[0])
		fmt.Println("Accept: */*")
		fmt.Println("Connection: close")
		fmt.Println()
	}		
	res, err := c.Do(req)
	if err != nil {
		fmt.Println("Client policy Error ", err)
	}
	/*
	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan(){
		fmt.Println(scanner.Text()) 
	}*/

	body, _ := io.ReadAll(res.Body)
	fmt.Printf("%s", body)
}
