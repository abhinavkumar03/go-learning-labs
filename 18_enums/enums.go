package main

import "fmt"

type orderStatus int

const (
	Recieved orderStatus = iota
	Confirmed
	Prepared
	Delivered
)

type cookingStatus string

func changeOrderStatus(status orderStatus) {
	fmt.Println("changing order status to: ", status)
}

const (
	NotStarted cookingStatus = "Not Started"
	InProgress               = "In Progress"
	Completed                = "Completed"
)

func changeCookingStatus(status cookingStatus) {
	fmt.Println("changing cooking status to: ", status)
}

func main() {
	changeOrderStatus(Recieved)
	changeCookingStatus(InProgress)
}
