package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/smallnest/rpcx/client"
	"github.com/wangzhuzhen/programming_framework_for_go/rpcx_example/handler"
	"github.com/smallnest/rpcx/protocol"
)

var (
	zkAddr   = flag.String("zkAddr", "localhost:2181", "zookeeper address")
	basePath = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	flag.Parse()

	options := client.Option{
		Retries:        3,
		RPCPath:        *basePath,
		ConnectTimeout: 10 * time.Second,
		SerializeType:  protocol.MsgPack,
		CompressType:   protocol.None,
		BackupLatency:  10 * time.Millisecond,
	}

	d := client.NewZookeeperDiscovery(*basePath, "Arith", []string{*zkAddr}, nil)
	//xclient := client.NewXClient("Arith", client.Failover, client.RoundRobin, d, client.DefaultOption)
	xclient := client.NewXClient("Arith", client.Failover, client.RoundRobin, d, options)
	defer xclient.Close()

	args := &handler.Args{
		A: 10,
		B: 20,
	}

	for {

		reply := &handler.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(5 * time.Second)
	}

}