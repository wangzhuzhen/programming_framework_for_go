package handler

import (
	"context"
	"fmt"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith int

// 记录当前 Server 的地址
var Addr *string

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	fmt.Printf("Processing on server [%s]: %d * %d = %d\n", *Addr, args.A, args.B, reply.C)
	return nil
}

func (t *Arith) Add(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	fmt.Printf("Processing on server [%s]: %d + %d = %d\n", *Addr, args.A, args.B, reply.C)
	return nil
}

func (t *Arith) Say(ctx context.Context, args *string, reply *string) error {
	*reply = "hello " + *args + " from server " + *Addr
	return nil
}

// 获取当前 Server 的地址
func SetServerAddr(addr *string){
	Addr = addr
}
