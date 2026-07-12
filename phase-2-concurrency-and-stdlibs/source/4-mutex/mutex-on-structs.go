package main

import (
	"fmt"
	"sync"
)

/*
	Below is an example of using mutex with structs.
	Here, we have a root OrderBook struct which will store all orders and will be shared amongst multiple goroutines.
*/

type OrderStatus string

const (
	Pending   OrderStatus = "PENDING"
	Completed OrderStatus = "COMPLETED"
	Cancelled OrderStatus = "CANCELLED"
)

type Order struct {
	id     int
	qty    int
	status OrderStatus
}
type OrderBook struct {
	mutex  sync.Mutex
	orders map[int]Order
}

func CreateOrderBook() *OrderBook {
	return &OrderBook{
		orders: make(map[int]Order),
	}
}

func (ob *OrderBook) Insert(id int, qty int, wg *sync.WaitGroup) {
	defer wg.Done()

	ob.mutex.Lock()
	defer ob.mutex.Unlock()

	// Insertion
	order := Order{id, qty, Pending}
	ob.orders[id] = order
	fmt.Println("Inserted Order With Id:", id)
}
func (ob *OrderBook) Delete(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	ob.mutex.Lock()
	defer ob.mutex.Unlock()

	// Removal
	delete(ob.orders, id)
	fmt.Println("Deleted Order With Id:", id)
}
func (ob *OrderBook) Get(id int, wg *sync.WaitGroup) Order {
	defer wg.Done()

	ob.mutex.Lock()
	defer ob.mutex.Unlock()

	if order, ok := ob.orders[id]; ok {
		return order
	}
	return Order{}
}
func (ob *OrderBook) UpdateStatus(id int, newStatus OrderStatus, wg *sync.WaitGroup) {
	defer wg.Done()

	ob.mutex.Lock()
	defer ob.mutex.Unlock()

	if order, ok := ob.orders[id]; ok {
		order.status = newStatus
		ob.orders[id] = order
		fmt.Println("Updated Status Of Order With Id:", id)
	}
}

func MutexOnStructs() {
	orderBook := CreateOrderBook()
	wg := sync.WaitGroup{}

	fmt.Println("Initialised OrderBook, Starting Our Concurrent Operations...")
	wg.Add(8)
	go orderBook.Insert(1, 10, &wg)
	go orderBook.Insert(2, 20, &wg)
	go orderBook.Insert(3, 30, &wg)
	go orderBook.Insert(4, 40, &wg)
	go orderBook.Insert(5, 50, &wg)
	go orderBook.UpdateStatus(1, Completed, &wg)
	go orderBook.Delete(2, &wg)
	go orderBook.Insert(6, 60, &wg)
	wg.Wait()

	wg.Add(4)
	go orderBook.UpdateStatus(3, Cancelled, &wg)
	go orderBook.Delete(4, &wg)
	go orderBook.Insert(7, 70, &wg)
	go orderBook.UpdateStatus(6, Completed, &wg)
	wg.Wait()

	wg.Add(1)
	order := orderBook.Get(3, &wg)
	wg.Wait()
	fmt.Println("Fetched Order:", order.id, order.qty, order.status)
}
