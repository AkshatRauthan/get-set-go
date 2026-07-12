package main

import (
	"fmt"
	"sync"
)

/*
	Mutex: Mutual Exclusion Lock
	Mutex is a type of locking mechanism that helps us to lock memory resources that are shared between
	multiple goroutines when someone is doing a read/write operation on them.

	This is done to ensure that unwanted things like race conditions, dirty reads/writes do no occur.
	We see these things happening in our ../1-goroutines/roroutines-race-conditions.go
	Lets solve that problem using mutex now.
*/

/*
	Here, the mutex.Lock() function is used: This function places a Write on the data.
	It means no other variable can place a RLock (Read) or Lock (Write Lock).
	RLock -> Read Lock -> Data can be Read by other but no written. Therefore, Other RLock() proceeds but Lock() gets blocked.
	Lock -> Write Lock -> No one can read or write data. Therefore, both RLock() and Lock() will get blocked.

	RLock => Is implemented by another module RWMutex. This lock instread of a single state [(0,1) in case of Lock], have to
	store a counter for no. of RLock() currently applied. [Like a semaphore].
	Due to which it has an overhead of concurrently updating the counter.

	Therefore, this kind of lock is exclusively used in places where;  NUM(write) <<< NUM(reads).
	[Then we will be saving Read operations form being blocked by other Read operations due to usage of RLock()]

	But in normal scenario, where NUM(writes) is still high, if we use RLock(), our Reads will still be getting blocked by many
	Write Operations, and we would still be getting the overhead of concurrently updating RLock internal counter.

	Therefore, USE RLOCK ONLY WHEN NUM(Writes) << NUM(Reads)
*/

/*
	Here, we are using a Mutex to protect single primitive value [int].
	But, its kinda overkill and is here only for demonstration purposes.
	In real life scenarios, use atomic.Int64, atomic.bool etc for these purposes....
	As they use CPU level atomic clocks and always guarantee concurrency even under high load.

	Keep Mutex only for complex data-structures like maps, structs etc.
*/

func incrementCounter(count *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	mutex.Lock()
	defer mutex.Unlock()

	*count++
}

func MutexBasics() {
	count := 0
	countMutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 1; i <= 100; i++ {
		go incrementCounter(&count, &wg, &countMutex)
	}
	wg.Wait()
	fmt.Println("Increment Counter finished running 100 times")
	fmt.Println("Final Value Of Count:", count) // Here no need to use mutex as all goroutines have already been executed.
}
