package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-demo/hello" // 导入生成的 proto 包
)

func main() {
	// 1. 连接到 server (因为没有 SSL 证书，所以使用 insecure 模式)
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("无法连接到服务端: %v", err)
	}
	defer conn.Close()

	// 2. 创建一个客户端存根
	c := pb.NewGreeterClient(conn)

	// 3. 设置 1 秒的超时上下文 (Context 在 Go 微服务中非常重要)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 4. 发起远程调用
	log.Println("正在呼叫服务端...")
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Golang 探索者"})
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}

	// 5. 打印返回结果
	log.Printf("收到服务端响应: %s", r.GetMessage())
}