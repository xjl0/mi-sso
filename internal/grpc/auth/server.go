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
	validate *validation.App
}

func Register(gRPC *grpc.Server, validate *validation.App) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{validate: validate})
}

func (s *serverAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	if err := s.validate.Validate.Struct(struct {
		Email    string `validate:"required,email"`
		Pass     string `validate:"required,min=8,max=255"`
		ClientID int32  `validate:"required,min=1,max=255"`
	}{
		Email:    req.GetUsername(),
		Pass:     req.GetPassword(),
		ClientID: req.GetClientId(),
	}); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid credentials")
	}

	return &ssov1.LoginResponse{
		Token: "test",
	}, nil
}

func (s *serverAPI) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	panic("implement me")
}

func (s *serverAPI) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	panic("implement me")
}
