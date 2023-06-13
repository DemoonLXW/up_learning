package service

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/database/ent/permission"
	"github.com/DemoonLXW/up_learning/database/ent/role"
)

type ManagementService struct {
	DB *ent.Client
}

func rollback(tx *ent.Tx, description string, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		return fmt.Errorf("%s rollback failed %w: %v", description, err, rerr)
	}
	return fmt.Errorf("%s failed: %w", description, err)
}

func (serv *ManagementService) CreatePermission(toCreates []*ent.Permission) error {
	ctx := context.Background()

	num, err := serv.DB.Permission.Query().Aggregate(ent.Count()).Int(ctx)
	if err != nil {
		return fmt.Errorf("create permissions count failed: %w", err)
	}

	actions := make([]string, len(toCreates))
	bulk := make([]*ent.PermissionCreate, len(toCreates))
	for i, v := range toCreates {
		actions[i] = *v.Action
		bulk[i] = serv.DB.Permission.Create().SetID(uint16(num + i + 1)).SetAction(*v.Action).SetDescription(*v.Description)
	}

	checkRepeatAction, err := serv.DB.Permission.Query().Where(permission.ActionIn(actions...)).Exist(ctx)
	if err != nil {
		return fmt.Errorf("create permissions check repeat action failed: %w", err)
	}
	if checkRepeatAction {
		return fmt.Errorf("repeat action")
	}

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("create permission start a transaction failed: %w", err)
	}

	err = tx.Permission.CreateBulk(bulk...).Exec(ctx)
	if err != nil {
		return rollback(tx, "create permissions", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("create permissions transaction commit failed: %w", err)
	}

	return nil

}

func (serv *ManagementService) UpdatePermission(toUpdate *ent.Permission) error {
	ctx := context.Background()

	originalPermission, err := serv.DB.Permission.Query().Where(permission.ID(toUpdate.ID)).WithRoles().First(ctx)
	if err != nil {
		return fmt.Errorf("update permission find original failed: %w", err)
	}
	if !originalPermission.DeletedTime.Equal(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)) {
		return fmt.Errorf("permission is already deleted")
	}
	if len(originalPermission.Edges.Roles) != 0 {
		return fmt.Errorf("permission is already used")
	}

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("update permission start a transaction failed: %w", err)
	}

	err = tx.Permission.Create().
		SetID(toUpdate.ID).
		SetAction(*toUpdate.Action).
		SetDescription(*toUpdate.Description).
		SetModifiedTime(time.Now()).
		OnConflict().
		UpdateAction().
		UpdateDescription().
		UpdateModifiedTime().
		Exec(ctx)
	if err != nil {
		return rollback(tx, "update permission", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("update permission transaction commit failed: %w", err)
	}

	return nil
}

func (serv *ManagementService) RetrievePermission(current, pageSize int, like, sort, order string) ([]*ent.Permission, error) {
	ctx := context.Background()

	offset := (current - 1) * pageSize

	permissions, err := serv.DB.Permission.Query().
		Where(
			permission.Or(
				permission.ActionContains(like),
				permission.DescriptionContains(like),
			),
		).
		Limit(pageSize).
		Offset(offset).
		Order(func(s *sql.Selector) {
			if order == "desc" {
				s.OrderBy(sql.Desc(sort))
			} else {
				s.OrderBy(sql.Asc(sort))
			}
		}).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("retrieve permission query failed: %w", err)
	}

	return permissions, nil

}

func (serv *ManagementService) DeletePermission(toDeleteIDs []uint16) error {
	ctx := context.Background()

	originalPermissions, err := serv.DB.Permission.Query().Where(permission.IDIn(toDeleteIDs...)).WithRoles().All(ctx)
	if err != nil {
		return fmt.Errorf("delete permissions find original failed: %w", err)
	}
	if len(originalPermissions) == 0 {
		return fmt.Errorf("delete permissions are not found")
	}
	for _, v := range originalPermissions {
		if !v.DeletedTime.Equal(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)) {
			return fmt.Errorf("permission (id:%d) is already deleted", v.ID)
		}
		if len(v.Edges.Roles) != 0 {
			return fmt.Errorf("permission (id:%d) is already used", v.ID)
		}
	}

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("delete permissions start a transaction failed: %w", err)
	}

	err = tx.Permission.Update().
		SetDeletedTime(time.Now()).
		Where(permission.IDIn(toDeleteIDs...)).
		Exec(ctx)
	if err != nil {
		return rollback(tx, "delete permissions", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("delete permissions transaction commit failed: %w", err)
	}

	return nil
}

func (serv *ManagementService) CreateRole(toCreates []*ent.Role) error {
	ctx := context.Background()

	num, err := serv.DB.Role.Query().Aggregate(ent.Count()).Int(ctx)
	if err != nil {
		return fmt.Errorf("create roles count failed: %w", err)
	}

	names := make([]string, len(toCreates))
	bulk := make([]*ent.RoleCreate, len(toCreates))
	for i, v := range toCreates {
		names[i] = *v.Name
		bulk[i] = serv.DB.Role.Create().SetID(uint8(num + i + 1)).SetName(*v.Name).SetDescription(*v.Description)
	}

	checkRepeatName, err := serv.DB.Role.Query().Where(role.NameIn(names...)).Exist(ctx)
	if err != nil {
		return fmt.Errorf("create roles check repeat name failed: %w", err)
	}
	if checkRepeatName {
		return fmt.Errorf("repeat role")
	}

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("create role start a transaction failed: %w", err)
	}

	err = tx.Role.CreateBulk(bulk...).Exec(ctx)
	if err != nil {
		return rollback(tx, "create roles", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("create roles transaction commit failed: %w", err)
	}

	return nil
}

func (serv *ManagementService) UpdateRole(toUpdate *ent.Role) error {
	ctx := context.Background()

	originalRole, err := serv.DB.Role.Query().Where(role.ID(toUpdate.ID)).WithPermissions().WithUsers().First(ctx)
	if err != nil {
		return fmt.Errorf("update role find original failed: %w", err)
	}
	if !originalRole.DeletedTime.Equal(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)) {
		return fmt.Errorf("role is already deleted")
	}
	if len(originalRole.Edges.Permissions) != 0 || len(originalRole.Edges.Users) != 0 {
		return fmt.Errorf("role is already used")
	}

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("update role start a transaction failed: %w", err)
	}

	err = tx.Role.Create().
		SetID(toUpdate.ID).
		SetName(*toUpdate.Name).
		SetDescription(*toUpdate.Description).
		SetModifiedTime(time.Now()).
		OnConflict().
		UpdateName().
		UpdateDescription().
		UpdateModifiedTime().
		Exec(ctx)
	if err != nil {
		return rollback(tx, "update role", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("update role transaction commit failed: %w", err)
	}

	return nil
}

func (serv *ManagementService) RetrieveRole(current, pageSize int, like, sort, order string) ([]*ent.Role, error) {
	ctx := context.Background()

	offset := (current - 1) * pageSize

	roles, err := serv.DB.Role.Query().
		Where(
			role.Or(
				role.NameContains(like),
				role.DescriptionContains(like),
			),
		).
		Limit(pageSize).
		Offset(offset).
		Order(func(s *sql.Selector) {
			if order == "desc" {
				s.OrderBy(sql.Desc(sort))
			} else {
				s.OrderBy(sql.Asc(sort))
			}
		}).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("retrieve role query failed: %w", err)
	}

	return roles, nil

}
