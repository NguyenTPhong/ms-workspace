package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"go.uber.org/zap"
	"ms-workspace/ms-notification/global"
	"ms-workspace/ms-notification/internal/v1/repository"
	"time"

	"ms-workspace/ms-notification/internal/v1/entity"
	"ms-workspace/ms-notification/internal/v1/repository/model"
)

type EmailUseCase interface {
	SendActiveEmail(ctx context.Context, request *entity.SendActiveEmailRequest) (*model.Email, error)
}

type emailUseCase struct {
	emailRepo repository.EmailRepository

	APIKey      string
	SenderName  string
	SenderEmail string
}

func NewEmailUseCase(repo repository.EmailRepository, apiKey, senderName, senderEmail string) EmailUseCase {
	useCase := &emailUseCase{
		emailRepo:   repo,
		APIKey:      apiKey,
		SenderName:  senderName,
		SenderEmail: senderEmail,
	}
	return useCase
}

func (e *emailUseCase) SendActiveEmail(ctx context.Context, request *entity.SendActiveEmailRequest) (*model.Email, error) {

	// send email
	from := mail.NewEmail(e.SenderName, e.SenderEmail)

	subject := "Active your account on ms-workspace"

	to := mail.NewEmail(request.Name, request.Email)

	plainTextContent := fmt.Sprintf("Hi %s, please active your account on ms-workspace by click this link: %s", request.Name, request.Url+"?code="+request.Code)
	htmlContent := "<p>" + plainTextContent + "</p>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(e.APIKey)

	response, err := client.Send(message)

	if err != nil {
		return nil, err
	}

	messageID := uuid.New().String()
	headers := response.Headers
	if mID, has := headers["X-Message-Id"]; has && len(mID) > 0 {
		messageID = mID[0]
	}

	result := &model.Email{
		UserId:    request.UserId,
		Email:     request.Email,
		MessageId: messageID,
		CreatedAt: time.Now(),
		Events: []model.EmailEvent{
			{
				Event:     "ship to sendgrid",
				CreatedAt: time.Now(),
			},
		},
	}

	// save db
	if err = e.emailRepo.SaveEmail(ctx, result); err != nil {
		global.Logger.Error("save email to db error", zap.Error(err), zap.Any("result", result))
	}

	return result, nil
}
