package main

import (
	"engine"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"scheduler"
)

// Args: WebFetch.exe TargetUrl Re
// Return: TargetInformation

func main()  {

	if len(os.Args) != 3{
		fmt.Printf("usage: WebFetch.exe -u TargetUrl -r Re\n" +
			"example: WebFetch.exe -u https://greysec.net/ -r `<span class=\" subject_new\" id=\".*\"><a href=\"([^\"]+)\">`")
	}

	var TargetUrl string
	var Re string

	flag.StringVar(&TargetUrl, "u", "", "TargetUrl")
	flag.StringVar(&Re, "r", "", "Re")
	flag.Parse()

	//CheckArgs(TargetUrl)	This function maybe in other go file

	e := engine.ConcurrentEngine{
		Scheduler: 		&scheduler.QueuedScheduler{},
		WorkerCount: 	300,
		ItemChan: 		nil,
	}
	e.Run(engine.Request{
		Url: 				nil,
		ParserFunc: 		nil,
	})
}

func CheckArgs(TargetUrl string) bool {	// Maybe not use
	res, _ := http.Get(TargetUrl)
	defer res.Body.Close()
	if(res.StatusCode != 200){
		log.Fatal("Expect TargetUrl StatusCode, But Got: " + string(res.StatusCode))
		return false
	}
	return true
}