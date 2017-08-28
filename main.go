package main

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
)

const (
	link = ""
	requests = 1000
	sleep = 0
)


func MakeRequest(url string, ch chan<-string, num int) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Sprintf("%d requests were made before error", num)
	}
	if resp != nil && resp.Body != nil{
		body, _ := ioutil.ReadAll(resp.Body)
		secs := time.Since(start).Seconds()
		ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
	} 
	
}

func main() {
	start := time.Now()
	ch := make(chan string)
	for i := 0; i < requests; i++ {
		go MakeRequest(link, ch, i)
		if sleep > 0 {
			time.Sleep(sleep * time.Millisecond)
		}
		if i > 1000 && i % 10000 == 0 {
			elapsed := time.Since(start).Seconds()
			speed := float64(i) / elapsed
			fmt.Printf("%.2fs elapsed\nSpeed is %.2f\n", elapsed, speed)
		}
		fmt.Println(i)
	}
	// for i := 0; i < requests; i++ {
	// 	fmt.Println(<-ch, i)
	// }
	elapsed := time.Since(start).Seconds()
	speed := requests / elapsed
	fmt.Printf("%.2fs elapsed\nSpeed is %.2f/s\n", elapsed, speed)
}
