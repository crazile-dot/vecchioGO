package main

import (
	"myProject/consumers"
	"myProject/types"
	"net/rpc"
)

/*
	This is a test for the consumer.
 */

func main() {
	var reply types.Message
	var consumer *rpc.Client

	var a = types.Message{"producer1", "consumer1", "Hello world!64", "string", 0, 0}
	var b = types.Message{"producer2", "consumer2", "Hello world!65", "string", 0, 0}

	consumer = consumers.ConsumerConnection()
	consumers.GetAllCall(consumer, "Fifo_12345")
	consumers.GetOneCall(consumer, a, reply, "Fifo_12345")
	consumers.GetOneCall(consumer, b, reply, "Fifo_12345")
	consumers.GetAllCall(consumer, "Fifo_12345")

}
