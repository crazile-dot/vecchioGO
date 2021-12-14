package types

import "time"

type Message struct {
	Source string
	Dest string
	Payload string
	Label string
	State int // 0 if free, 1 if elaborated
	Lock int // 0 if free, 1 if locked
}

type Queue struct {
	Queue []Message
	ResTimeOut time.Duration //response timeout
	VisTimeOut time.Duration //visibility timeout
	RoutingKey string //id name of the queue
}

/*
	Two structures to make some parameters' passing easy
 */
type Args1 struct {
	ItemPayload string
	FifoKey string
}

type Args2 struct {
	Item Message
	FifoKey string
}
