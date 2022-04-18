package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
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
	r, e := c.CreateUser(context.Background(), &proto.CreateUserRequest{
		Password: "haishina",
		Account:  "junbin.yang2",
		NickName: "wtf",
		Mobile:   "15915862738",
		Email:    "182764407@qq.com",
		UserInfo: &proto.UserInfo{
			RealName: "张十三",
			Age:      18,
			Sex:      1,
			Nation:   "汉",
			Address:  "深圳",
			IdCard:   "440506000000000000",
		},
	})
	if e != nil {
		fmt.Printf("调用服务端代码失败: %s", e)
		return
	}
	fmt.Printf("调用成功: %s\n", r.String())
}
