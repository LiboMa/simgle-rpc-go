package main

import (
	//"bufio"
	"bytes"
	"encoding/gob"
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

func EncodeToBytes(p interface{}) []byte {

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("uncompressed size (bytes): ", len(buf.Bytes()))
	return buf.Bytes()
}

func DecodeToPerson(s []byte) Message {

	msg := Message{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&msg)
	if err != nil {
		log.Fatal(err)
	}
	return msg
}

func main() {

	// set client flags

	// func Int(name string, value int, usage string) *int
	clientID := flag.Int("id", 1, "the ID of RPC client.")
	flag.Parse()

	fmt.Printf("TCP RPC client..with id %d \n", *clientID)

	client, err := rpc.Dial("tcp", "localhost:6789")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected for the RPC server")

	// fmt.Println("ticker start...")
	msg := Message{Id: *clientID, Msg: "PING"}

	for {
		line := EncodeToBytes(msg)
		time.Sleep(time.Second * 1)

		if err != nil {
			log.Fatal(err)
		}

		log.Println("Sending message <- ", msg)
		// var response Response
		var response interface{}

		err = client.Call("Listener.Getline", line, &response)

		if err != nil {
			fmt.Println("ERR")
			log.Fatal(err)
		}

		var reply Response
		err = json.Unmarshal(response.([]uint8), &reply)

		if err != nil {
			fmt.Println("ERR")
			log.Fatal(err)
		}
		// log.Printf("Replied %v \n", response.([]uint8))
		// log.Printf("Replied %s \n", response)
		// log.Printf("Timestame: %s, Original message: %v, Response message: %v, \nreponse Type: %T", reply.Ts, reply.Resp, reply.Mark, reply)
		log.Printf("Timestame: %s, Original message: %v, Response message: %v \n", reply.Ts, reply.Resp, reply.Mark)

	}

}
