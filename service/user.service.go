package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"time"

	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/database/ent/user"
	"github.com/redis/go-redis/v9"
)

type UserService struct {
	DB    *ent.Client
	Redis *redis.Client
}

func (serv *UserService) Login(account, password string) (string, error) {
	ctx := context.Background()

	u, err := serv.DB.User.Query().Where(
		user.And(
			user.Or(
				user.Account(account),
				user.Email(account),
				user.Phone(account),
			),
			user.Password(password),
			user.DeletedTimeEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
		),
	).First(ctx)
	if err != nil {
		return "", fmt.Errorf("user login check credential failed: %w", err)
	}

	expire_date := time.Now().In(time.Local)
	expire_date_byte, err := expire_date.MarshalText()
	if err != nil {
		return "", fmt.Errorf("user login convert time to string failed: %w", err)
	}
	key := fmt.Sprintf("user:%d", u.ID)
	cookie := fmt.Sprintf("%x", md5.Sum([]byte(key+expire_date.String())))
	user_cookies, err := serv.Redis.HGetAll(ctx, key).Result()
	switch {
	case err == redis.Nil, err == nil && len(user_cookies) < 5:
		{
			_, err = serv.Redis.HSet(ctx, key, cookie, string(expire_date_byte)).Result()
			if err != nil {
				return "", fmt.Errorf("user login set cookie failed: %w", err)
			}
		}
	case err != nil:
		{
			return "", fmt.Errorf("user login get cookies failed: %w", err)
		}
	case len(user_cookies) >= 5:
		{
			to_delete_key := ""
			to_delete_time := time.Now()
			for k, v := range user_cookies {
				expire_time, err := time.Parse(time.RFC3339, v)
				if err != nil {
					return "", fmt.Errorf("user login parse string to time failed: %w", err)
				}
				if expire_time.Before(to_delete_time) {
					to_delete_key = k
					to_delete_time = expire_time
				}
			}

			_, err := serv.Redis.HDel(ctx, key, to_delete_key).Result()
			if err != nil {
				return "", fmt.Errorf("user login delete cookie failed: %w", err)
			}

			_, err = serv.Redis.HSet(ctx, key, cookie, string(expire_date_byte)).Result()
			if err != nil {
				return "", fmt.Errorf("user login set cookie failed: %w", err)
			}
		}
	}

	return cookie, nil
}
