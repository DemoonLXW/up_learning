package service

import (
	"context"
	"fmt"

	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/database/ent/user"
)

func (serv *ManagementService) CreateUser(toCreates []*ent.User) error {
	ctx := context.Background()

	num, err := serv.DB.User.Query().Aggregate(ent.Count()).Int(ctx)
	if err != nil {
		return fmt.Errorf("create users count failed: %w", err)
	}

	accounts := make([]string, len(toCreates))
	bulk := make([]*ent.UserCreate, len(toCreates))
	for i, v := range toCreates {
		accounts[i] = *v.Account
		bulk[i] = serv.DB.User.Create().
			SetID(uint32(num + i + 1)).
			SetAccount(*v.Account).
			SetPassword(*v.Password)
	}

	checkRepeatAccount, err := serv.DB.User.Query().Where(user.AccountIn(accounts...)).Exist(ctx)
	if err != nil {
		return fmt.Errorf("create users check repeat user failed: %w", err)
	}
	if checkRepeatAccount {
		return fmt.Errorf("repeat user")
	}

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("create user start a transaction failed: %w", err)
	}

	err = tx.User.CreateBulk(bulk...).Exec(ctx)
	if err != nil {
		return rollback(tx, "create users", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("create users transaction commit failed: %w", err)
	}

	return nil

}
