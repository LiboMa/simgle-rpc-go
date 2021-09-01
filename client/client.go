package main

import (
	//"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Message struct {
	Id  int
	Msg string
}

type Response struct {
	Ts   time.Time `json: "ts"`
	Mark string    `json:"Mark"`
	Resp Message   `json: Message`
}

func main() {

	// set client flags

	// func Int(name string, value int, usage string) *int
	clientID := flag.Int("id", 1, "the ID of RPC client.")
	flag.Parse()

	client, err := rpc.Dial("tcp", "localhost:6789")
	fmt.Printf("TCP RPC client..with id %d \n", *clientID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected for the RPC server")

	// fmt.Println("ticker start...")
	msg := Message{Id: *clientID, Msg: "PING"}

	for {
		//line := EncodeToBytes(msg)
		time.Sleep(time.Second * 1)

		if err != nil {
			log.Fatal(err)
		}

		log.Println("Sending message <- ", msg)
		// var response Response
		var response interface{}

		//err = client.Call("Listener.Getline", msg, &response)
		err = client.Call("Listener.Getline", msg, &response)

		if err != nil {
			fmt.Println("ERR")
			log.Fatal(err)
		}

		// Load message body to struct
		var reply Response
		err = json.Unmarshal(response.([]uint8), &reply)

		if err != nil {
			fmt.Println("ERR")
			log.Fatal(err)
		}
		log.Printf("%s\n", response)
		log.Printf("Timestame: %s, Original message: %v, Response message: %v \n", reply.Ts, reply.Resp, reply.Mark)

	}

}
