package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nano second precision
}

type customer struct {
	name   string
	age    int
	email  string
	orders []order
}

func newOrder(id string, amount float32) *order {
	return &order{
		id:     id,
		amount: amount,
		status: "booked",
	}
}

// receiver function to print order details
func (o order) printOrder() {
	fmt.Println("Order ID:", o.id)
	fmt.Println("Amount:", o.amount)
	fmt.Println("Status:", o.status)
	fmt.Println("Created At:", o.createdAt)
}

func (o *order) updateStatus(newStatus string) {
	o.status = newStatus
}

func main() {
	myOrder := order{
		id:     "1",
		amount: 50.00,
		status: "booked",
	}

	myOrder.createdAt = time.Now()

	fmt.Println("myOrder: ", myOrder)
	fmt.Println("Status: ", myOrder.status)

	myOrder.updateStatus("shipped")
	fmt.Println("Updated Status: ", myOrder.status)

	myOrder.printOrder()

	myOrder2 := order{}   // all fields get default zero values
	myOrder2.printOrder() // default values of int, float, string, time.Time are printed

	myOrder3 := newOrder("3", 75.00) // returns pointer to order
	myOrder3.printOrder()

	userDetails := struct { // anonymous struct
		name  string
		email string
	}{
		name:  "Alice",
		email: "alice@example.com",
	}

	fmt.Println(userDetails)

	customer1 := customer{
		name:  "Bob",
		age:   30,
		email: "bob@gmail.com",
		orders: []order{
			*myOrder3,
		},
	}

	customer1.orders = append(customer1.orders, myOrder)
	customer1.orders[0].createdAt = time.Now()
	fmt.Println(customer1)
	customer1.orders[0].printOrder()
}
