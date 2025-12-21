package main

import "fmt"

type payment struct {
	id     string
	method string
	amount float32
}

// concrete implementations of razorpay payment processors
type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay: ", amount)
}

// concrete implementations of stripe payment processors
type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe: ", amount)
}

// Open close principle violation example
func (p payment) makePayment(amount float32) {
	// depending on the payment method, use different processor
	switch p.method {
	case "stripe":
		stripeProcessor := stripe{}
		stripeProcessor.pay(amount)
	case "razorpay":
		razorpayProcessor := razorpay{}
		razorpayProcessor.pay(amount)
	default:
		fmt.Println("payment method not supported")
	}
}

type stripePaymentProcessot struct {
	id     string
	method stripe
	amount float32
}

// Open close principle adherence example
func (p stripePaymentProcessot) makeStripePayment(amount float32) {
	p.method.pay(amount)
}

// interface for payment processors
type paymenter interface {
	pay(amount float32)
}

type genericPayment struct {
	id     string
	method paymenter
	amount float32
}

func main() {

	// Open close principle violation example
	myPayment := payment{
		id:     "101",
		method: "razorpay",
		amount: 150.75,
	}
	myPayment1 := payment{
		id:     "102",
		method: "stripe",
		amount: 250.00,
	}

	myPayment.makePayment(myPayment.amount)
	myPayment1.makePayment(myPayment1.amount)

	// Open close principle adherence example
	myStripePayment := stripePaymentProcessot{
		id:     "201",
		amount: 300.50,
	}

	myStripePayment.makeStripePayment(myStripePayment.amount)

	// using interface to make payment
	myGenericPayment := genericPayment{
		id:     "301",
		method: razorpay{},
		amount: 400.00,
	}

	myGenericPayment.method.pay(myGenericPayment.amount)
}
