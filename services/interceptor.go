package services

/*
	gRpc中间件-访问验证
*/
import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"    // grpc 响应状态码
	"google.golang.org/grpc/metadata" // grpc metadata包
	"svc-user/model"
)

func getMetaAppId(ctx context.Context) string {
	md, _ := metadata.FromIncomingContext(ctx)
	var appid string
	if val, ok := md["appid"]; ok {
		appid = val[0]
	}
	return appid
}

func auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return grpc.Errorf(codes.Unauthenticated, "无认证信息")
	}

	var (
		appKey    string
		appSecret string
	)

	if val, ok := md["appid"]; ok {
		appKey = val[0]
	}

	if val, ok := md["appkey"]; ok {
		appSecret = val[0]
	}

	app := model.App{}
	result := model.M.Where("app_key = ? AND app_secret = ? ", appKey, appSecret).First(&app)

	if result.RowsAffected == 0 {
		return grpc.Errorf(codes.Unauthenticated, "认证信息无效: appKey=%s, appSecret=%s", appKey, appSecret)
	}

	if !app.Status {
		return grpc.Errorf(codes.Unauthenticated, "应用已停用，认证信息无效: appKey=%s, appSecret=%s", appKey, appSecret)
	}

	return nil
}

// interceptor 拦截器
func interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	err := auth(ctx)
	if err != nil {
		return nil, err
	}
	// 继续处理请求
	return handler(ctx, req)
}
