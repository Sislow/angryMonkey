package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

/*
average would have been a time.Duration, but it doesn't play well with json.marshal
Golang known issue since 2015
https://github.com/golang/go/issues/10275
*/
type Stat struct {
	count   int
	average int
}

type Config struct {
	done    bool
	kill    bool
	runTime time.Duration
	temp    string
}

var s Stat
var c Config

const form = "/Users/oliver/go/src/github.com/sislow/angryMonkey/public/static/form.html"

var wg = &sync.WaitGroup{}

func handleShutdown(w http.ResponseWriter, r *http.Request) {
	fmt.Println("shutdown called. finishing requests")
	// wait for wg to empty
	defer wg.Wait()
	// give a moment to ensure no other requests come through
	time.Sleep(3 * time.Second)
	// Go's version of a while loop
	for c.done {
		c.kill = true
		log.Println("Shutting down...")
		os.Exit(0)
	}
}

func handleStats(w http.ResponseWriter, r *http.Request) {
	s.average = int((c.runTime / time.Duration(s.count)) / time.Second)
	if r.Method == "GET" {
		statMap := map[string]int{"average": s.average, "count": s.count}
		out, err := json.Marshal(statMap)
		if err != nil {
			log.Println("JsonErr", err)
		}
		fmt.Fprintf(w, string(out))
		log.Println(string(out))
	} else {
		fmt.Fprintf(w, "Http "+r.Method+" Request not allowed")
		log.Println("unsupported method request: " + r.Method)
	}
}

func handleHash(w http.ResponseWriter, r *http.Request) {
	wg.Add(1)
	c.done = false
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, form)
	case "POST":
		if !c.kill {
			// track time and apply sleep
			start := time.Now()
			time.Sleep(5 * time.Second)
			// find request if empty check form
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				// didn't realize this is not used
				log.Println("Error: ", err)
			} else {
				c.temp = string(reqBody)
				c.temp = strings.Replace(c.temp, "password=", "", -1)
			}

			// do not allow blank entries
			if c.temp != "" {
				passEncoded := encryptPassword([]byte(c.temp))

				// To curl provider or frontend
				fmt.Fprintf(w, "Password inserted: "+c.temp)
				fmt.Fprintf(w, "\nPassword Encoded: "+string(passEncoded))

				// for server log
				log.Println("Password inserted: " + c.temp)
				log.Println("Password Encoded: " + string(passEncoded))
				s.count++
			}
			elapsed := time.Since(start)
			c.runTime += elapsed
			log.Println("Time Elapsed: ", elapsed)
			log.Println("Total Time: ", c.runTime)
		} else {
			log.Println("Rejected Request")
			return
		}
	default:
		fmt.Fprintf(w, "unsupported method type: "+r.Method)
		log.Println("unsuported method request : " + r.Method)
	}
	wg.Done()
	c.done = true
}

func Router() {
	// each of these is a go routine. Seperated process for each item
	http.HandleFunc("/shutdown", handleShutdown)
	http.HandleFunc("/stats", handleStats)
	http.HandleFunc("/hash", handleHash)
	http.HandleFunc("/", handleHash)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
