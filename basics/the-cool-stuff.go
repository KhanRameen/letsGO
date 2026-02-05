package main

import (
	"fmt"
	"math/rand"
	// "math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{} //like counters. they wait till they go back to zero before finishing the program
var m = sync.Mutex{}      //provides lock unlock function
var rw = sync.RWMutex{}   // provides read write locks

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func goroutines() { //concurrency
	t0 := time.Now()
	for i := range len(dbData) {
		wg.Add(1)    //adds to the counter when the functions calls
		go dbCall(i) //go keywords makes the program not wait for the exection to complete for the task but moves to next execution. the program however can end without waiting for them to complete that is where we use wait groups from sync package
	}

	wg.Wait()
	fmt.Printf("Total execution time for db call: %v\n", time.Since(t0))
	fmt.Printf("The results are: %v\n", results)
}

func dbCall(i int) {
	//simulate db call delay
	// var delay float32 = rand.Float32() * 2000 //random float number upto 2 sec
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Microsecond)
	// fmt.Printf("The result from database after delay of %vs is: %v\n", delay, dbData[i])
	fmt.Printf("The result from database is: %v\n", dbData[i])

	// m.Lock() //locks so no two processes try to update the db at the same time
	// results = append(results, dbData[i])
	// m.Unlock() //unlock for other process ones done

	save(dbData[i])
	log()
	wg.Done() //decrements from the counter when the functions end
}

//results: execution time without goroutines:  8.2768ms
//results: execution time with goroutines:   2.6786ms

func log() {
	rw.RLock() //read only locks: lock processes from writing when you are reading the data to prevent potential changes while reading
	fmt.Printf("\nThe Current Results are: %v\n", results)
	rw.RUnlock()
}

func save(result string) {
	rw.Lock() //full locks
	results = append(results, result)
	rw.Unlock()
}

func channels() {
	//goroutines memory sharing pipes (mutext locks but better).
	// var c = make(chan int) //chan type, no.of vlues.

	// c <- 1                 // set 1 to c. here 1 is stored as an array value, using it outside the go process causes dealocks "channels waits till some process asks for the content in their memory and without a goroutine the code gets stuck on this line.. no asynchronous stuff"

	// var i = <-c //pops value out from channel c to i

	// go process(c)
	// for i := range c {
	// 	fmt.Println(i)
	// }

	var bufferChannel = make(chan int, 5) //this lets the process function do the job and keeps running the main (where it is called) in the backgroun
	go process(bufferChannel)
	for i := range bufferChannel {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}


	chickenNTofu()
}

func process(c chan int) {
	defer close(c) //close the channel when no in use
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Exiting function")

}

func chickenNTofu() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)

	var websites = []string{"x.com", "y.com", "z.com"}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkTofuPrices(websites string, c chan string) {
	var maxTofuPrice float32 = 5

	for {
		time.Sleep(time.Second * 1)
		var tofu_price = rand.Float32() * 20
		if tofu_price < maxTofuPrice {
			c <- websites
			break
		}
	}
}
func checkChickenPrices(websites string, c chan string) {
	var maxChickenPrice float32 = 15

	for {
		time.Sleep(time.Second * 1)
		var tofu_price = rand.Float32() * 20
		if tofu_price < maxChickenPrice {
			c <- websites
			break
		}
	}
}


func sendMessage(chickenChannel chan string, tofuChannel chan string){
	select{
	case websites := <-chickenChannel
		fmt.Println("Deal in chicken at website", websites)
	
	case websites := <- tofuChannel
		fmt.Println("Deal in tofu at website", websites)
	}
}