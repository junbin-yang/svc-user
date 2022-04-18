package services

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"svc-user/jwt"
	"svc-user/proto"
)

func (this *Svr) TokenCreate(ctx context.Context, in *proto.TokenCreateRequest) (*proto.Token, error) {
	token, err := jwt.Create(uint64(in.UserId), false, in.Issuer)
	if err != nil {
		return nil, grpc.Errorf(codes.DataLoss, err.Error())
	}
	return &proto.Token{Token: token}, nil
}

func (this *Svr) TokenVerify(ctx context.Context, in *proto.Token) (*proto.TokenVerifyReply, error) {
	appid := getMetaAppId(ctx)
	token, err := jwt.Parse(in.Token)
	if err != nil {
		return nil, grpc.Errorf(codes.DataLoss, err.Error())
	}
	return &proto.TokenVerifyReply{
		UserId:          uint32(token.UserId),
		ExpiresAt:       uint64(token.StandardClaims.ExpiresAt),
		Issuer:          token.StandardClaims.Issuer,
		MenuPermissions: getMenus(0, appid, uint(token.UserId)),
	}, nil
}

func (this *Svr) TokenRefresh(ctx context.Context, in *proto.Token) (*proto.Token, error) {
	token, err := jwt.Parse(in.Token)
	if err != nil {
		return nil, grpc.Errorf(codes.DataLoss, err.Error())
	}

	newToken, err := token.Refresh()
	if err != nil {
		return nil, grpc.Errorf(codes.DataLoss, err.Error())
	}

	return &proto.Token{Token: newToken}, nil
}
