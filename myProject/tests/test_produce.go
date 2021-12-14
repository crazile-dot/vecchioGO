package main

import (
	"myProject/producers"
	"myProject/types"
	"net/rpc"
)

/*
	This is a test for the producer.
 */

func main() {
	var reply types.Message
	var producer *rpc.Client

	var a = types.Message{"producer1", "consumer1", "Hello world!64", "string", 0, 0}
	var b = types.Message{"producer2", "consumer2", "Hello world!65", "string",0, 0}
	var c = types.Message{"producer3", "consumer3", "Hello world!66", "string",0, 0}

	producer = producers.ProducerConnection()

	producers.GetQueueCall(producer, producers.P_queue)
	producers.AddCall(producer, a, reply, "Fifo_12345")
	producers.AddCall(producer, b, reply, "Fifo_12345")
	producers.AddCall(producer, c, reply, "Fifo_12345")
}
