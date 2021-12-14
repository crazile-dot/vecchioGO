package consumers

import (
	"fmt"
	"log"
	"net/rpc"
	"myProject/types"
	"time"
)

var C_queue types.Queue

func ConsumerConnection() *rpc.Client{
	cons, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("connection error", err)
	}
	return cons
}

func GetOneCall(cons *rpc.Client, item types.Message, reply types.Message, key string) {
	i := 0 //index mantaining the number of sent requests (at-least-once semantics)
	for {
		g_deliver := cons.Go("API.GetItem", types.Args1{item.Payload, key}, &reply, nil)
		i += 1
		g_deliverCall := <-g_deliver.Done
		if g_deliverCall == g_deliver {
			fmt.Println("Delivered message: ", reply.Payload)
		} else {
			time.Sleep(C_queue.ResTimeOut) //retransmitting timeout
			continue
		}
		if i >= 1 {
			break
		}
	}

	j := 0 //index mantaining the number of sent requests (at-least-once semantics)
	for {
		d_deliver := cons.Go("API.DeleteItem", types.Args2{item, key}, &reply, nil)
		j += 1
		d_deliverCall := <-d_deliver.Done
		if d_deliverCall == d_deliver {
			fmt.Println("Deleted message: ", reply.Payload)
		} else {
			time.Sleep(C_queue.ResTimeOut) //retrasmitting timeout
			continue
		}
		if j >= 1 {
			break
		}
	}
}

func GetAllCall(cons *rpc.Client, key string) {
	var queue types.Queue
	i := 0 //index mantaining the number of sent requests (at-least-once semantics)
	for {
		deliver := cons.Go("API.GetAll", types.Args1{"", key}, &queue, nil)
		i += 1
		deliverCall := <-deliver.Done
		if deliverCall == deliver {
			C_queue = queue
			fmt.Println("Delivered message (C_queue): ", queue.Queue)
		} else {
			time.Sleep(C_queue.ResTimeOut) //retransmitting timeout
			continue
		}
		if i >= 1 {
			break
		}
	}

}


