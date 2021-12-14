package server

import (
	"fmt"
	"myProject/types"
	"os"
	"time"
)

var fifoQ types.Queue

/*
	Function called at the initialization of the message queue.
	It sets fixed parameters for the queue.
 */
func setParams() {
	fifoQ.ResTimeOut = 2*time.Millisecond
	fifoQ.VisTimeOut = 30*time.Millisecond
	fifoQ.RoutingKey = "Fifo_12345"
}

/*
	Function returning all the messages in the queue with the state for each one
 */
func getAllMsg(key string) types.Queue{
	var ret types.Queue
	if key == fifoQ.RoutingKey { //check if the queue is the one I'm looking for
		ret = fifoQ
	} else {
		fmt.Println("Queue does not match!")
		os.Exit(0)
	}

	return ret
}

/*
	Adds new messages to the queue.
	They can be duplicated.
 */
func addMsg(item types.Message, key string) types.Message {
	if empty() {
		setParams()
	}
	if key == fifoQ.RoutingKey { //check if the queue is the one I'm looking for
		fifoQ.Queue = append(fifoQ.Queue, item)
	} else {
		fmt.Println("Queue does not match!")
		os.Exit(0)
	}

	return item
}

/*
	Gets the message from the queue, inputing a certain payload.
	It also manages the concurrency between consumers by a bit inside the message.
 */
func getByMsg(text string, key string) types.Message {
	var item types.Message
	if key == fifoQ.RoutingKey { //check if the queue is the one I'm looking for
		j := 0
		for { //thanks to this loop, if the Lock is not free, the consumer tries again until it's unlocked
			for i, val := range fifoQ.Queue {
				if val.Lock < 1 { //if Lock is free, the consumer can elaborate the message
					fifoQ.Queue[i].Lock += 1
					if val.Payload == text {
						fifoQ.Queue[i].State += 1
						item = fifoQ.Queue[i]
						j += 1
						time.Sleep(10 * time.Millisecond) //a certain elaboration of the message
						go timeoutManagement(i) //a goroutine manages the visibility timeout
						break
					} else {
						j += 1
						fifoQ.Queue[i].Lock -= 1
					}
				} else { //if message is loading, the other consumers waits
					continue
				}
			}
			if j > 0 {
				break
			}
		}
		if item.Lock > 0 {
			item.Lock -= 1
		}
	} else {
		fmt.Println("Queue does not match!")
		os.Exit(0)
	}
	return item
}

/*
	Deletes a message from the queue if it was elaborated
 */
func deleteMsg(item types.Message, key string) types.Message{
	var del types.Message
	if key == fifoQ.RoutingKey {
		for index, value := range fifoQ.Queue {
			if value.Payload == item.Payload {
				if value.Lock > 0 {
					fifoQ.Queue = append(fifoQ.Queue[:index], fifoQ.Queue[index+1:]...)
					del = item
				}
			}
		}
	} else {
		fmt.Println("Queue does not match!")
		os.Exit(0)
	}
	return del
}

/*
	It manages the visibility timeout
	Timeout-based semantics
*/
func timeoutManagement(index int) {
	time.Sleep(fifoQ.VisTimeOut)
	if !empty() {
		if exists(fifoQ.Queue[index].Payload) && fifoQ.Queue[index].Lock > 0 {
			fifoQ.Queue[index].Lock -= 1
		}
	}

}

/*
 	Checks if a message exists in the queue
 */
func exists(text string) bool {
	var res bool
	for _, value := range fifoQ.Queue {
		if value.Payload == text {
			res = true
			break
		} else {
			res = false
		}
	}
	return res
}

/*
	Checks if the queue is empty
 */
func empty() bool {
	if len(fifoQ.Queue) < 1 {
		return true
	} else {
		return false
	}
}