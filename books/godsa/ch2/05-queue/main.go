package main

import "fmt"

// Order
type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

// Queue-Array
type Queue []*Order

// New
func (o *Order) New(priority int, quantity int, product string, customerName string) {
	o.priority = priority
	o.quantity = quantity
	o.product = product
	o.customerName = customerName
}

// Add
func (q *Queue) Add(order *Order) {
	if len(*q) == 0 {
		*q = append(*q, order)
	} else {
		appended := false
		for i, addedOrder := range *q {
			if order.priority > addedOrder.priority {
				*q = append((*q)[:i], append(Queue{order}, (*q)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*q = append(*q, order)
		}
	}
}

func main() {
	queue := make(Queue, 0)
	order1 := &Order{}
	priority1 := 2
	quantity1 := 20
	product1 := "Computer"
	customerName1 := "Greg White"
	order1.New(priority1, quantity1, product1, customerName1)
	order2 := &Order{}
	priority2 := 1
	quantity2 := 10
	product2 := "Monitor"
	customerName2 := "Jhon Smith"
	order2.New(priority2, quantity2, product2, customerName2)
	queue.Add(order1)
	queue.Add(order2)
	for i := range queue {
		fmt.Println(queue[i])
	}
}
