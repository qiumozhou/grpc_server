package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "test/pb"
	"google.golang.org/grpc/reflection"
)

type MyGrpcServer struct{}

func (myserver *MyGrpcServer) MyTest(ctx context.Context, rq *pb.Request) (*pb.Response, error) {
	fmt.Println(rq)
	return &pb.Response{BackJson:"111"}, nil
}




func main(){
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	s := grpc.NewServer() //创建gRPC服务

	/**注册接口服务
	 * 以定义proto时的service为单位注册，服务中可以有多个方法
	 * (proto编译时会为每个service生成Register***Server方法)
	 * 包.注册服务方法(gRpc服务实例，包含接口方法的结构体[指针])
	 */
	pb.RegisterTesterServer(s, &MyGrpcServer{})
	/**如果有可以注册多个接口服务,结构体要实现对应的接口方法
	 * user.RegisterLoginServer(s, &server{})
	 * minMovie.RegisterFbiServer(s, &server{})
	 */
	// 在gRPC服务器上注册反射服务
	reflection.Register(s)
	// 将监听交给gRPC服务处理
	err = s.Serve(lis)
	if  err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
