package repository

import (
	"encoding/json"
	"fmt"
	"ms-workspace/ms-user/global/config"

	"github.com/go-redis/redis"

	"ms-workspace/ms-user/internal/v1/repository/model"
)

const (
	RedisAuthTokenKey = "user_authenticate_token"
)

type AuthCacheRepository interface {
	CacheAuthSession(userId int64, session *model.Session) error
	GetAuthSessionFromCache(userId int64) (*model.Session, error)
	PublishSendActiveEmailJob(message string) error
}

type AuthCacheRepositoryImpl struct {
	redis *redis.Client
}

func NewAuthCacheRepository(redis *redis.Client) AuthCacheRepository {
	return &AuthCacheRepositoryImpl{redis: redis}
}

func (u *AuthCacheRepositoryImpl) CacheAuthSession(userId int64, session *model.Session) error {
	data, _ := json.Marshal(session)
	return u.redis.HSet(RedisAuthTokenKey, fmt.Sprint(userId), string(data)).Err()
}

func (u *AuthCacheRepositoryImpl) GetAuthSessionFromCache(userId int64) (*model.Session, error) {
	res, err := u.redis.HGet(RedisAuthTokenKey, fmt.Sprint(userId)).Result()
	if err != nil {
		return nil, err
	}

	var session model.Session
	if err = json.Unmarshal([]byte(res), &session); err != nil {
		return nil, err
	}
	return &session, nil
}

func (u *AuthCacheRepositoryImpl) PublishSendActiveEmailJob(message string) error {
	return u.redis.Publish(config.SendActiveEmailTopic, message).Err()
}
