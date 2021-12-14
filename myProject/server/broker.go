package server

import (
	"log"
	"myProject/types"
	"net"
	"net/http"
	"net/rpc"
)

type API int
var fifo types.Queue

/*
	Connecting the server to the 4040 port
 */
func ServerConnection() {
	var api = new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API ", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Listening error ", err)
	}

	log.Printf("serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("error serving ", err)
	}
}

//get the queue
func (a *API) GetAll(arg1 types.Args1, reply *types.Queue) error {
	*reply = getAllMsg(arg1.FifoKey)

	return nil
}

//consumer function to get an item from the queue
func (a *API) GetItem(arg1 types.Args1, reply *types.Message) error {
	var item types.Message
	item = getByMsg(arg1.ItemPayload, arg1.FifoKey)
	*reply = item

	//go DeleteItems()

	return nil
}

//producer function to add an item to the queue
func (a *API) AddItem(arg2 types.Args2, reply *types.Message) error {
	var item types.Message
	item = addMsg(arg2.Item, arg2.FifoKey)
	*reply = item

	return nil
}

//consumer function to delete an item from the queue
func (a *API) DeleteItem(arg2 types.Args2, reply *types.Message) error {
	var del types.Message
	del = deleteMsg(arg2.Item, arg2.FifoKey)
	*reply = del

	return nil
}