package main

import (
	"net"
	"net/rpc/jsonrpc"
	"learngo/rpc"
	"fmt"
)

func main() {
	dial, err := net.Dial("tcp", ":1234")
	if err != nil{
		panic(err)
	}

	client := jsonrpc.NewClient(dial)

	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{10, 3}, &result)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(result, err)
	}

	err = client.Call("DemoService.Div", rpcdemo.Args{10, 0}, &result)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(result, err)
	}


}
