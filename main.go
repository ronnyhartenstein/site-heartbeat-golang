package main

import (
	//"fmt"
	// "log"
	"sync"
)

func main() {
	urls := ReadHosts()
	ch := make(chan bool)
	var wg sync.WaitGroup
	//fmt.Println("Start running")
	for _, site := range urls {
		//fmt.Printf("for %s\n", site.Url)
		wg.Add(1)
		go func(s Site) {
			//fmt.Printf("go %s\n", s.Url)
			defer wg.Done()
			ok := Check(s)
			ch <- ok
		}(site)
	}
	// closer
	go func() {
		wg.Wait()
		close(ch)
	}()
	// fetch chan results
	for range ch {
	}
}
