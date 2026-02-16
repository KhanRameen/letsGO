package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	jobs := make(chan string)    //this channel will be used to send jobs to the worker pool
	results := make(chan string) //this channel will be used to receive results from the worker pool

	var wg sync.WaitGroup //create a wait group to wait for all workers to finish

	workerCount := 3 //max number of workers in the pool

	//start the worker pool
	for i := 0 ; i<workerCount; i++ {
		wg.Add(1)//increment the wait group counter for each worker
		go worker(i, jobs, results, &wg) //start a worker goroutine
	}

	
	//send jobs to the workers
	go func ()  { //anonymous goroutine to send jobs to the workers, called in the background
		urlBatch1 := []string{
			"https://www.github.com/KhanRameen",
			"https://www.instagram.com/her.artsysights",
			"https://www.linkedin.com/in/khan-rameen",
		}		
		
		for _, url := range urlBatch1 {
			fmt.Println("BATCH1")
			jobs <- url //send the URL to the jobs channel
		}
		
		time.Sleep(time.Second * 2) //simulate a delay before sending the next batch of jobs (background work)
		
		urlBatch2 := []string{
			"https://www.github.com",
			"https://www.instagram.com",
			"https://www.linkedin.com",
			}		
			
			for _, url := range urlBatch2 {
			fmt.Println("BATCH2")
			jobs <- url //send the URL to the jobs channel
		}

		close(jobs) //close the jobs channel to signal that no more jobs are coming
	}()

	//collect results from the workers
	go func ()  {
		wg.Wait() //wait for all workers to finish
		close(results)
	}()

	for results := range results { //range over the results channel to receive results until it's closed
		fmt.Println(results)
	}

}