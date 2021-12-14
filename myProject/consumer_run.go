package main

import (
	"flag"
	"myProject/types"
	"myProject/consumers"
	"net/rpc"

)

/*
	This is the main file to run the consumer.
	You can run this file on your machine and set your parameters (at least Payload and key)
*/

func main() {
	var reply types.Message
	var consumer *rpc.Client

	srcPtr := flag.String("Source", "producer", "the producer")
	destPtr := flag.String("Dest", "consumer", "the consumer")
	paylPtr := flag.String("Payload", "", "the payload of the message")
	labPtr := flag.String("Label", "string", "the type of the payload")
	statePtr := flag.Int("State", 0, "the state of the message")
	lockPtr := flag.Int("Lock", 0, "the lock for the message")

	keyPtr := flag.String("key", "", "the name of the queue")

	flag.Parse()

	var a = types.Message{*srcPtr, *destPtr, *paylPtr, *labPtr, *statePtr, *lockPtr}

	consumer = consumers.ConsumerConnection()
	consumers.GetAllCall(consumer, *keyPtr)
	consumers.GetOneCall(consumer, a, reply, *keyPtr)

}
