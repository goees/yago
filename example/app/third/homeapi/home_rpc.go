package homeapi

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/goees/yago"

	"google.golang.org/grpc"

	"github.com/goees/yago/base/basethird"
	pb "github.com/goees/yago/example/app/modules/home/homerpc/homepb"
)

type homeRpcApi struct {
	basethird.RpcThird
}

// Usage: InsRpc().Hello()
func InsRpc() *homeRpcApi {
	name := "home_rpc_api"
	v := yago.Component.Ins(name, func() interface{} {
		api := new(homeRpcApi)

		// rpc 配置
		err := api.InitConfig(name)
		if err != nil {
			log.Fatal("init rpc api config error")
		}

		// 添加业务自己的拦截器
		api.AddUnaryClientInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			fmt.Println("method:", method, "client before")
			err := invoker(ctx, method, req, reply, cc, opts...)
			fmt.Println("method:", method, "client after")
			return err
		})

		api.AddStreamClientInterceptor(func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (stream grpc.ClientStream, e error) {
			fmt.Println("method:", method, "stream before")
			clientStream, err := streamer(ctx, desc, cc, method, opts...)
			return clientStream, err
		})

		return api
	})
	return v.(*homeRpcApi)
}

// eg. homeapi.InsRpc().Hello()
func (a *homeRpcApi) Hello() {

	var name = "zhangsan"

	req := &pb.HelloRequest{Name: name}

	conn, _ := a.GetConn()
	ctx, cancel := a.GetCtx()
	defer cancel()

	c := pb.NewHomeClient(conn)
	resp, err := c.Hello(ctx, req)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ok:", resp.Data)
}

func (a *homeRpcApi) HelloStream() {

	var name = "zhangsan"
	req := &pb.HelloRequest{Name: name}

	conn, _ := a.GetConn()
	ctx, cancel := a.GetCtx()
	defer cancel()

	c := pb.NewHomeClient(conn)
	stream, err := c.HelloStream(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println(reply)
	}
}
