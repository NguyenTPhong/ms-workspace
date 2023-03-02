package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"ms-workspace/ms-user/internal/v1/repository/model"
	"ms-workspace/package/proto/ms-user/v1/go-proto"
)

func (s *Server) CreateUser(ctx context.Context, req *go_proto.CreateUserRequest) (*go_proto.CreateUserResponse, error) {

	user := &model.User{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
	}

	err := s.authUseCase.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &go_proto.CreateUserResponse{
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		CreatedAt:   timestamppb.New(user.CreatedAt),
		Id:          user.Id,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *go_proto.LoginRequest) (*go_proto.LoginResponse, error) {
	session, err := s.authUseCase.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &go_proto.LoginResponse{
		Token:     session.Token,
		ExpiredAt: timestamppb.New(session.ExpiredAt),
	}, nil
}

func (s *Server) Authentication(ctx context.Context, request *go_proto.AuthenticationRequest) (*go_proto.AuthenticationResponse, error) {

	session, err := s.authUseCase.Authentication(ctx, request.Token)
	if err != nil {
		return nil, err
	}

	return &go_proto.AuthenticationResponse{
		Id:          session.Id,
		Email:       session.Email,
		PhoneNumber: session.PhoneNumber,
		FirstName:   session.FirstName,
		LastName:    session.LastName,
		Status:      string(session.Status),
		LoggedInAt:  timestamppb.New(session.LoggedInAt),
		ExpiredAt:   timestamppb.New(session.ExpiredAt),
	}, nil

}
