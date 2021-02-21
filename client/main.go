package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/mp-hl-2021/grpc-example/api"
)

func main() {
	var opts = []grpc.DialOption{grpc.WithInsecure()}

	conn, err := grpc.Dial("localhost:8081", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := api.NewEchoClient(conn)

	echo, err := client.Do(context.Background(), &api.EchoRequest{
		Line: "hello, world",
		Num: 3,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(echo.EchoLine)
}

