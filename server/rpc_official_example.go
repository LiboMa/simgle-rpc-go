package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Mathops int

func (t *Mathops) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Mathops) Devide(args *Args, quo *Quotient) error {

	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B

	return nil
}

func main() {

	ops := new(Mathops)
	rpc.Register(ops)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":6789")

	if e != nil {
		log.Fatal("listen error:", e)
	}

	go http.Serve(l, nil)

}
