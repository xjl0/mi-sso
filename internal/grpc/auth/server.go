package auth

import (
	"context"
	ssov1 "github.com/xjl0/mi-protos/gen/go/sso"
	"github.com/xjl0/mi-sso/internal/app/validation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	validation *validation.App
}

func Register(gRPC *grpc.Server, validation *validation.App) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{validation: validation})
}

func (s *serverAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	if err := s.validation.Validate.Struct(struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required,min=8,max=255"`
		ClientId int32  `validate:"required,min=1,max=255"`
	}{
		Email:    req.GetUsername(),
		Password: req.GetPassword(),
		ClientId: req.GetClientId(),
	}); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &ssov1.LoginResponse{
		Token: "test",
	}, nil
}

func (s *serverAPI) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	if err := s.validation.Validate.Struct(struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required,min=8,max=255"`
	}{
		Email:    req.GetUsername(),
		Password: req.GetPassword(),
	}); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &ssov1.RegisterResponse{
		Uid: 0,
	}, nil
}

func (s *serverAPI) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	if err := s.validation.Validate.Struct(struct {
		Uid int32 `validate:"required,min=1"`
	}{
		Uid: req.GetUid(),
	}); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &ssov1.IsAdminResponse{
		IsAdmin: false,
	}, nil
}
