package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error){
	<-rateLimiter
	resp, err := http.Get(url)
	if err != nil{
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Printf("Error: status code: %d, Error site: %s\n", resp.StatusCode, url)
		return nil,err
	}
	return ioutil.ReadAll(resp.Body)
}
