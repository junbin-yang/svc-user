package main

import (
	"context"
	"fmt"
	"github.com/junbin-yang/golib/json"
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
	str := `
{
    "node": [
        {
            "code": "用户管理",
            "name": "用户管理",
            "sequence": 1,
            "uri": "/baidu",
            "remark": "",
            "actions": [
                "a",
                "b",
                "c"
            ],
            "children": [
                {
                    "code": "公众用户",
                    "name": "公众用户",
                    "remark": "",
                    "sequence": 1,
                    "uri": "/123",
                    "actions": [
                        "编辑",
                        "删除"
                    ],
                    "children": [
                        
                    ]
                },
                {
                    "code": "村委用户",
                    "name": "村委用户",
                    "remark": "",
                    "sequence": 2,
                    "uri": "/123/123",
                    "actions": [
                        "新建",
                        "查看"
                    ],
                    "children": [
                        
                    ]
                }
            ]
        },
        {
            "code": "角色管理",
            "name": "角色管理",
            "sequence": 2,
            "uri": "/role",
            "remark": "",
            "actions": [
                "a",
                "d"
            ],
            "children": [
                
            ]
        }
    ]
}`
	req := proto.PermissionsTplRequest{}
	json.JsonToObject(str, &req)
	r, e := c.PermissionsTpl(context.Background(), &req)
	if e != nil {
		fmt.Printf("调用服务端代码失败: %s", e)
		return
	}
	fmt.Printf("调用成功: %s\n", r.String())
}
