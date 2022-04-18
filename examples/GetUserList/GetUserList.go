package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	structpb "google.golang.org/protobuf/types/known/structpb"
	auth "svc-user/examples"
	"svc-user/proto"
)

func main() {
	// 连接服务器
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	// 指定自定义认证
	opts = append(opts, grpc.WithPerRPCCredentials(new(auth.CustomCredential)))
	conn, err := grpc.Dial(":8888", opts...)
	if err != nil {
		fmt.Printf("连接服务端失败: %s", err)
		return
	}
	defer conn.Close()
	// 新建一个客户端
	c := proto.NewUserClient(conn)
	// 调用服务端函数
	newListValue, _ := structpb.NewList([]interface{}{"account", "mobile", "email"})
	r, e := c.GetUserList(context.Background(), &proto.PaginateRequest{
		Page:   1,
		Limit:  10,
		Search: map[string]*structpb.ListValue{"150": newListValue},
	})
	if e != nil {
		fmt.Printf("调用服务端代码失败: %s", e)
		return
	}
	fmt.Printf("调用成功: %s\n", r.String())
}
