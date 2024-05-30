// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"github.com/huangsihao7/scooter-WSVA/user/rpc/internal/logic"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"
)

type UseSrvServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUseSrvServer
}

func NewUseSrvServer(svcCtx *svc.ServiceContext) *UseSrvServer {
	return &UseSrvServer{
		svcCtx: svcCtx,
	}
}

func (s *UseSrvServer) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UseSrvServer) Register(ctx context.Context, in *user.RegisterRequest) (*user.RegisterResponse, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UseSrvServer) UserInfo(ctx context.Context, in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

func (s *UseSrvServer) Update(ctx context.Context, in *user.UpdateReq) (*user.UpdateResp, error) {
	l := logic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}