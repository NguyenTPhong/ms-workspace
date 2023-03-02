package grpc

import (
	"context"
	"ms-workspace/ms-notification/internal/v1/entity"
	"ms-workspace/package/proto/ms-notification/v1/go-proto"
)

func (s *Server) SendActiveEmail(ctx context.Context, req *go_proto.SendActiveEmailRequest) (*go_proto.SendActiveEmailResponse, error) {
	email, err := s.emailUseCase.SendActiveEmail(ctx, &entity.SendActiveEmailRequest{
		UserId: req.UserId,
		Name:   req.Name,
		Email:  req.Email,
		Code:   req.Code,
		Url:    req.Url,
	})

	if err != nil {
		return nil, err
	}
	return &go_proto.SendActiveEmailResponse{
		MessageId: email.MessageId,
	}, nil
}
