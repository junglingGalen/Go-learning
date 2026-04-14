package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection" // 引入反射包

	pb "grpc-demo/hello" // 导入刚才生成的 proto 包
)

// server 结构体，用来实现 Greeter 服务
type server struct {
	pb.UnimplementedGreeterServer // 必须嵌入这个，保证向前兼容
}

// 实现 SayHello 方法
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("收到客户端请求，名字是: %v", in.GetName())
	return &pb.HelloReply{Message: "你好, " + in.GetName() + "! 欢迎来到 Go gRPC 世界。"}, nil
}

func main() {
	// 1. 监听本地 50051 端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("无法监听端口: %v", err)
	}

	// 2. 创建 gRPC 服务器实例
	s := grpc.NewServer()

	// 3. 将我们的服务注册到 gRPC 服务器上
	pb.RegisterGreeterServer(s, &server{})

	// 4. 【核心秘籍】开启服务端反射，方便 Postman 等工具调试！
	reflection.Register(s)

	log.Printf("gRPC 服务端已启动，监听端口 50051...")

	// 5. 启动服务并阻塞等待
	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务端启动失败: %v", err)
	}
}