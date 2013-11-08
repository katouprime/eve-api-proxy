package main

import (
	"fmt"
	"ieveapi/apicache"
	"log"
	"net/http"
	"os"
	"runtime"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	if conf.LogFile != "" {
		logfp, err := os.OpenFile(conf.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(fmt.Sprintf("Cannot Open Log File: %s", err))
		}
		log.SetOutput(logfp)
	}

	if conf.Threads == 0 {
		conf.Threads = runtime.NumCPU()
	}
	runtime.GOMAXPROCS(conf.Threads)
	log.Printf("EVEAPIProxy Starting Up with %d threads...", conf.Threads)

	apicache.NewClient(apicache.NilCache)
	startWorkers()

	var handler APIHandler
	log.Fatal(http.ListenAndServe(conf.Listen, &handler))
}
