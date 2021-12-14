package producers

import (
	"fmt"
	"log"
	"myProject/types"
	"net/rpc"
	"time"
)

var P_queue types.Queue

/*
	Connecting the producer with the server
 */
func ProducerConnection() *rpc.Client{
	prod, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("connection error", err)
	}
	return prod
}

func AddCall(prod *rpc.Client, item types.Message, reply types.Message, key string) {
	i := 0 //index mantaining the number of sent requests (at-least-once semantics)
	for {
		divCall1 := prod.Go("API.AddItem", types.Args2{item, key}, &reply, nil)
		i += 1
		replyCall1 := <-divCall1.Done
		if divCall1 == replyCall1 {
			fmt.Println("Received message: ", reply.Payload)
		} else {
			time.Sleep(P_queue.ResTimeOut) //retransmitting timeout
			continue
		}
		if i >= 1 {
			break
		}
	}
}

func GetQueueCall(cons *rpc.Client, queue types.Queue) {
	i := 0 //index mantaining the number of sent requests (at-least-once semantics)
	for {
		deliver := cons.Go("API.GetAll", "", &queue, nil)
		i += 1
		deliverCall := <-deliver.Done
		if deliverCall != deliver {
			P_queue = queue
			time.Sleep(queue.ResTimeOut) //retransmitting timeout
			continue
		}
		if i >= 1 {
			break
		}
	}
}