package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"time"

	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/database/ent/menu"
	"github.com/DemoonLXW/up_learning/database/ent/user"
	"github.com/redis/go-redis/v9"
)

type UserService struct {
	DB    *ent.Client
	Redis *redis.Client
}

func (serv *UserService) CreateAndSaveToken(id uint32) (string, error) {
	ctx := context.Background()

	expire_date := time.Now().In(time.Local).Add(2 * time.Hour)
	expire_date_byte, err := expire_date.MarshalText()
	if err != nil {
		return "", fmt.Errorf("create and save token convert time to string failed: %w", err)
	}
	key := fmt.Sprintf("user:%d", id)
	token := fmt.Sprintf("%x", md5.Sum([]byte(key+expire_date.String())))
	user_tokens, err := serv.Redis.HGetAll(ctx, key).Result()
	switch {
	case err == redis.Nil, err == nil && len(user_tokens) < 5:
		{
			_, err = serv.Redis.HSet(ctx, key, token, string(expire_date_byte)).Result()
			if err != nil {
				return "", fmt.Errorf("create and save token set token failed: %w", err)
			}
		}
	case err != nil:
		{
			return "", fmt.Errorf("create and save token get tokens failed: %w", err)
		}
	case len(user_tokens) >= 5:
		{
			to_delete_key := ""
			to_delete_time := time.Now().Add(3 * time.Hour)
			for k, v := range user_tokens {
				expire_time, err := time.Parse(time.RFC3339, v)
				if err != nil {
					return "", fmt.Errorf("create and save token parse string to time failed: %w", err)
				}
				if expire_time.Before(to_delete_time) {
					to_delete_key = k
					to_delete_time = expire_time
				}
			}

			_, err := serv.Redis.HDel(ctx, key, to_delete_key).Result()
			if err != nil {
				return "", fmt.Errorf("create and save token delete token failed: %w", err)
			}

			_, err = serv.Redis.HSet(ctx, key, token, string(expire_date_byte)).Result()
			if err != nil {
				return "", fmt.Errorf("create and save token set token failed: %w", err)
			}
		}
	}

	return token, nil
}

func (serv *UserService) Login(account, password string) (*ent.User, string, error) {
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
	).WithRoles().First(ctx)
	if err != nil {
		return nil, "", fmt.Errorf("user login check credential failed: %w", err)
	}

	token, err := serv.CreateAndSaveToken(u.ID)
	if err != nil {
		return nil, "", fmt.Errorf("user login create and save token failed: %w", err)
	}

	return u, token, nil
}

func (serv *UserService) Logout(id, token string) error {
	ctx := context.Background()

	key := "user:" + id
	num, err := serv.Redis.HDel(ctx, key, token).Result()
	switch {
	case num == 0:
		{
			return fmt.Errorf("user logout token does not exist failed: %w", err)
		}
	case err != nil:
		{
			return fmt.Errorf("user logout delete token failed: %w", err)
		}
	}

	return nil
}

func (serv *UserService) CheckCredential(id uint32, token string) (string, error) {
	ctx := context.Background()

	key := fmt.Sprintf("user:%d", id)
	val, err := serv.Redis.HGet(ctx, key, token).Result()
	switch {
	case err == redis.Nil:
		{
			return "", fmt.Errorf("user check credential token does not exist failed: %w", err)
		}
	case err != nil:
		{
			return "", fmt.Errorf("user check credential get token failed: %w", err)
		}
	}
	expire_time, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return "", fmt.Errorf("user check credential parse string to time failed: %w", err)
	}
	if time.Now().After(expire_time) {
		num, err := serv.Redis.HDel(ctx, key, token).Result()
		switch {
		case num == 0:
			{
				return "", fmt.Errorf("user check credential token does not exist failed: %w", err)
			}
		case err != nil:
			{
				return "", fmt.Errorf("user check credential delete token failed: %w", err)
			}
		}
		return "", fmt.Errorf("user check credential credential is expired failed: %w", err)
	}
	if time.Now().Before(expire_time.Add(-30 * time.Minute)) {
		return token, nil
	}

	num, err := serv.Redis.HDel(ctx, key, token).Result()
	switch {
	case num == 0:
		{
			return "", fmt.Errorf("user check credential token does not exist failed: %w", err)
		}
	case err != nil:
		{
			return "", fmt.Errorf("user check credential delete token failed: %w", err)
		}
	}
	new_token, err := serv.CreateAndSaveToken(id)
	if err != nil {
		return "", fmt.Errorf("user check credential create and save token failed: %w", err)
	}

	return new_token, nil

}

func (serv *UserService) FindOneUserById(id uint32) (*ent.User, error) {
	ctx := context.Background()
	user, err := serv.DB.User.Query().Where(user.IDEQ(id)).WithRoles().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("user find one by id failed: %w", err)
	}
	return user, nil
}

func (serv *UserService) FindMenuByUserId(id uint32) ([]*ent.Menu, error) {
	ctx := context.Background()

	menu, err := serv.DB.User.Query().Where(user.IDEQ(id)).QueryRoles().QueryMenu().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("menus find by user id failed: %w", err)
	}

	return menu, nil
}

func (serv *UserService) FindMenuByRoleIds(ids []uint8) ([]*ent.Menu, error) {
	ctx := context.Background()

	menu, err := serv.DB.Menu.Query().Where(menu.RidIn(ids...)).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("menus find by role ids failed: %w", err)
	}

	return menu, nil
}

func (serv *UserService) FindRolesWithMenusByUserId(id uint32) ([]*ent.Role, error) {
	ctx := context.Background()

	roles, err := serv.DB.User.Query().Where(user.IDEQ(id)).QueryRoles().WithMenu().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("roles with menus find by user id failed: %w", err)
	}

	return roles, nil
}
