package main

import (
	"fmt"

	"github.com/qmuntal/stateless"
)

// Define possible states
type OrderState string

const (
	OrderCreated  OrderState = "OrderCreated"
	OrderPaid     OrderState = "OrderPaid"
	OrderShipped  OrderState = "OrderShipped"
	OrderComplete OrderState = "OrderComplete"
)

// Define possible triggers
type Trigger string

const (
	TriggerPay     Trigger = "Pay"
	TriggerShip    Trigger = "Ship"
	TriggerDeliver Trigger = "Deliver"
)

func main() {
	// Initialize state machine with an initial state
	sm := stateless.NewStateMachine(OrderCreated)

	// Configure state transitions
	sm.Configure(OrderCreated).
		Permit(TriggerPay, OrderPaid)

	sm.Configure(OrderPaid).
		Permit(TriggerShip, OrderShipped)

	sm.Configure(OrderShipped).
		Permit(TriggerDeliver, OrderComplete)

	// Test transitions
	fmt.Println("Current State:", sm.MustState()) // Output: OrderCreated

	// Move to 'OrderPaid'
	sm.Fire(TriggerPay)
	fmt.Println("State after paying:", sm.MustState()) // Output: OrderPaid

	// Move to 'OrderShipped'
	sm.Fire(TriggerShip)
	fmt.Println("State after shipping:", sm.MustState()) // Output: OrderShipped

	// Move to 'OrderComplete'
	sm.Fire(TriggerDeliver)
	fmt.Println("State after delivery:", sm.MustState()) // Output: OrderComplete
}
