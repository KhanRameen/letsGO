package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

 var maxRetries = 3

func worker(id int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {  // worker receives jobs: <-chan string, wokder sends results: chan<- string
	defer wg.Done() //decrement the wait group counter when the worker is done
	
	rateLimit := time.Tick(500*time.Millisecond) //1 req per 500ms


	for url := range jobs { //range over the jobs channel to receive jobs until it's closed{}
	retries := 0
	<-rateLimit //wait for the rate limit ticker before making the next request	
	for{
		fmt.Printf("Worker %d checking %s\n Attempt %d\n", id, url, retries) //%d:int , %s:string	
		
		resp , err := http.Get(url) //make an HTTP GET request to the URL
		if err != nil {
			retries++
			if retries >= maxRetries {
				results <- fmt.Sprintf("Worker %d: Error checking %s: %v", id, url, err) //send error message to results channel
				break
			}
			continue
			} 			
			resp.Body.Close() //close the response body to free up resources
			results <- fmt.Sprintf("Worker %d: %s is up! Status Code: %d", id, url, resp.StatusCode) //send success message to results channel
			break
		}
	}
}		