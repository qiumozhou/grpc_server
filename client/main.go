package main
import (
	"fmt"
	"log"
	"os"
	"context"
	"google.golang.org/grpc"
	pb "test/pb"
)


func main() {
	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Waiter服务的客户端
	t := pb.NewTesterClient(conn)
	// 模拟请求数据
	res := "test123"
	// os.Args[1] 为用户执行输入的参数 如：go run ***.go 123
	if len(os.Args) > 1 {
		res = os.Args[1]
	}
	// 调用gRPC接口
	wq,err :=t.MyTest(context.Background(), &pb.Request{JsonStr: res})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(wq)

}