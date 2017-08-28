package main

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
)

const link = ""

func MakeRequest(url string, ch chan<-string) {
  start := time.Now()
  resp, err := http.Get(url)
  if err != nil {
  	fmt.Println(err)
  }
  secs := time.Since(start).Seconds()
  body, _ := ioutil.ReadAll(resp.Body)
  ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
}

func main() {
  requests := 10000
  start := time.Now()
  ch := make(chan string)
  for i := 0; i < requests; i++ {
      go MakeRequest(link, ch)
      defer fmt.Println(i)
  }

  for i := 0; i < requests; i++ {
    fmt.Println(<-ch)
  }
  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
