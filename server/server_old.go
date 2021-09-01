package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

type Listener int

// RPC method
// func (t *T) MethidName(argType T1, replyType *T2) error

//func (l *Listener) Getline(line []byte, ack *bool) error {

type Message struct {
	Id  int
	Msg string
}

type Response struct {
	Ts   time.Time
	Mark string
	Resp Message
}

func DecodeMessage(s []byte) Message {

	msg := Message{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&msg)
	if err != nil {
		log.Fatal(err)
	}
	return msg
}

//func (l *Listener) Getline(line []byte, resp *string) error {
func (l *Listener) Getline(line []byte, resp *interface{}) error {
	//func (l *Listener) Getline(msg *Message, resp *interface{}) error {
	//fmt.Println(string(line))
	msg := DecodeMessage(line)

	//log.Printf("line type -> %T, msg type -> %T\n", line, msg)

	fmt.Printf("handle message %v, from client: %d\n", msg.Msg, msg.Id)
	res := Response{Ts: time.Now(), Mark: "PONG", Resp: msg}
	resp_json, err := json.Marshal(res)

	if err != nil {
		log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	}

	*resp = resp_json
	return nil
}

func main() {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:6789")

	if err != nil {
		log.Fatal("RPC call failure", err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("TCP RPC Server listening on port 0.0.0.0.6789 ...")
	listener := new(Listener)
	rpc.Register(listener)
	rpc.Accept(inbound)

}
