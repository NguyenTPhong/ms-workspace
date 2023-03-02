package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	go_proto "ms-workspace/package/proto/ms-notification/v1/go-proto"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"ms-workspace/ms-user/global"
	"ms-workspace/ms-user/global/config"
	_const "ms-workspace/ms-user/global/const"
	"ms-workspace/ms-user/internal/v1/repository"
	"ms-workspace/ms-user/internal/v1/repository/model"
	"ms-workspace/package/logger"
	"ms-workspace/package/sha"
)

type AuthUseCase interface {
	CreateUser(ctx context.Context, user *model.User) error
	Login(ctx context.Context, email, password string) (*model.Session, error)
	Authentication(ctx context.Context, token string) (session *model.Session, err error)
}

type authUseCase struct {
	authRepo      repository.AuthRepository
	authCacheRepo repository.AuthCacheRepository
}

func NewUserUseCase(userRepository repository.AuthRepository, authCacheRepo repository.AuthCacheRepository) AuthUseCase {
	return &authUseCase{
		authRepo:      userRepository,
		authCacheRepo: authCacheRepo,
	}
}

func (u *authUseCase) isDuplicateUser(ctx context.Context, user *model.User) error {
	// find user by username
	dbUser, err := u.authRepo.FindFirstUser(repository.UserFilter{
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	})

	// process error
	if err != nil && err != gorm.ErrRecordNotFound {
		global.Logger.Error("login error", zap.Error(err), logger.TraceID(ctx))
		return fmt.Errorf(_const.InternalServerErr)
	}

	// check user exist
	if dbUser != nil {
		if dbUser.Email == user.Email {
			return fmt.Errorf(_const.EmailAlreadyExist)
		}

		return fmt.Errorf(_const.PhoneAlreadyExist)
	}

	return nil
}

func (u *authUseCase) CreateUser(ctx context.Context, user *model.User) error {

	// validate field
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		global.Logger.Error("validate error", zap.Error(err), logger.TraceID(ctx))
		return fmt.Errorf(_const.InvalidFieldData)
	}

	// find user by username
	if err := u.isDuplicateUser(ctx, user); err != nil {
		return err
	}

	// create user
	// gen salt
	user.Status = model.StatusActive
	user.Salt = uuid.New().String()
	user.Password = sha.Decode(user.Password, user.Salt)

	// save db
	err = u.authRepo.CreateUser(user)
	if err != nil {
		global.Logger.Error("login error", zap.Error(err), logger.TraceID(ctx))
		return fmt.Errorf(_const.InternalServerErr)
	}

	activeCode := &model.ActiveCode{
		UserId: user.Id,
		Code:   strings.ToUpper(uuid.New().String()),
	}

	if err = u.authRepo.CreateActiveCode(activeCode); err != nil {
		return err
	}

	// publish active email
	payload := go_proto.SendActiveEmailRequest{
		Email:  user.Email,
		UserId: user.Id,
		Name:   user.FirstName + " " + user.LastName,
		Code:   activeCode.Code,
		Url:    "http://localhost/active-account",
	}

	// publish to queue
	messageByte, _ := json.Marshal(payload)
	if err = u.authCacheRepo.PublishSendActiveEmailJob(string(messageByte)); err != nil {
		return err
	}

	return nil
}

func (u *authUseCase) Login(ctx context.Context, email, password string) (*model.Session, error) {
	// validate field
	if email == "" || password == "" {
		return nil, fmt.Errorf(_const.MissingEmailOrPassword)
	}

	// find user by email
	user, err := u.authRepo.FindFirstUser(repository.UserFilter{
		Email: email,
	})

	// process error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf(_const.UserNotFound)
		}
		global.Logger.Error("login error", zap.Error(err), logger.TraceID(ctx))
		return nil, fmt.Errorf(_const.InternalServerErr)
	}

	// check password
	sendHashedPass := sha.Decode(password, user.Salt)
	if sendHashedPass != user.Password {
		return nil, fmt.Errorf(_const.UserWrongPassword)
	}

	expiredAt := time.Now().Add(time.Duration(config.TokenLifeTime) * time.Minute)
	session := &model.Session{
		Id:          user.Id,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Status:      user.Status,
		LoggedInAt:  time.Now(),
		ExpiredAt:   expiredAt,
	}

	// generate token
	token, err := u.generateAuthToken(session)
	if err != nil {
		global.Logger.Error("generate token error", zap.Error(err), logger.TraceID(ctx))
		return nil, fmt.Errorf(_const.InternalServerErr)
	}

	session.Token = token
	// save session to redis
	err = u.authCacheRepo.CacheAuthSession(user.Id, session)
	if err != nil {
		global.Logger.Error("save session to redis error", zap.Error(err), logger.TraceID(ctx))
		return nil, fmt.Errorf(_const.InternalServerErr)
	}

	return session, nil
}

func (u *authUseCase) generateAuthToken(session *model.Session) (string, error) {
	secretKey := []byte(config.JWTKey)

	// Define the claims for the token
	claims := jwt.MapClaims{
		"id":           session.Id,
		"email":        session.Email,
		"phone_number": session.PhoneNumber,
		"first_name":   session.FirstName,
		"last_name":    session.LastName,
		"status":       session.Status,
		"expired_at":   session.ExpiredAt.Format(time.RFC3339),
		"logged_in_at": session.LoggedInAt,
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (u *authUseCase) Authentication(ctx context.Context, tokenString string) (session *model.Session, err error) {
	secretKey := []byte(config.JWTKey)

	// remove bearer word
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// Parse and verify the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify that the signing method is HMAC-SHA256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key for verification
		return secretKey, nil
	})
	if err != nil {
		global.Logger.Error("parse token error", zap.Error(err))
		return nil, err
	}

	// Verify that the token is valid
	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		global.Logger.Error("token is invalid", zap.Any("claims", token.Claims))
		return nil, fmt.Errorf(_const.Unauthorized)
	}

	defer func() {
		if r := recover(); r != nil {
			global.Logger.Error("parse token error", zap.Any("error", r))
			session = nil
			err = fmt.Errorf(_const.Unauthorized)
		}
	}()

	// get session from claims
	claims := token.Claims.(jwt.MapClaims)
	session = &model.Session{}
	session.Email = claims["email"].(string)
	session.Id = int64(claims["id"].(float64))
	session.PhoneNumber = claims["phone_number"].(string)
	session.FirstName = claims["first_name"].(string)
	session.LastName = claims["last_name"].(string)
	session.Status = model.UserStatus(claims["status"].(string))
	session.ExpiredAt, _ = time.Parse(time.RFC3339, claims["expired_at"].(string))
	session.LoggedInAt, _ = time.Parse(time.RFC3339, claims["logged_in_at"].(string))

	// check the token is expired
	if session.ExpiredAt.Before(time.Now()) {
		global.Logger.Error("token is expired", zap.Any("claims", token.Claims))
		return nil, fmt.Errorf(_const.Unauthorized)
	}

	// check session in redis, in case that server force logout user
	ssInCache, err := u.authCacheRepo.GetAuthSessionFromCache(session.Id)
	if err != nil || ssInCache == nil {
		global.Logger.Error("get session from redis error", zap.Error(err))
		return nil, fmt.Errorf(_const.Unauthorized)
	}

	return session, nil
}
