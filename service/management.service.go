package service

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/database/ent/permission"
	"github.com/DemoonLXW/up_learning/database/ent/role"
	"github.com/DemoonLXW/up_learning/database/ent/user"
	"github.com/DemoonLXW/up_learning/entity"
)

type ManagementService struct {
	DB *ent.Client
}

func (serv *ManagementService) CreatePermission(toCreates []*entity.ToAddPermission) error {
	ctx := context.Background()

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("create permission start a transaction failed: %w", err)
	}

	current := 0
	length := len(toCreates)
	for ; current < length; current++ {
		id, err := tx.Permission.Query().Where(permission.DeletedTimeNEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local))).FirstID(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return fmt.Errorf("create permission find a deleted permission id query failed: %w", err)
			}
			break
		}

		num, err := tx.Permission.Update().Where(permission.And(
			permission.IDEQ(id),
			permission.DeletedTimeNEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
		)).
			SetCreatedTime(time.Now()).
			SetModifiedTime(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)).
			SetDeletedTime(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)).
			SetAction(*toCreates[current].Action).
			SetDescription(*toCreates[current].Description).
			SetIsDisabled(false).
			Save(ctx)
		if err != nil {
			return rollback(tx, "create permission", err)
		}
		if num == 0 {
			return rollback(tx, "create permission", fmt.Errorf("create permission update deleted permission affect 0 row"))
		}

	}
	if current < length {
		num, err := tx.Permission.Query().Aggregate(ent.Count()).Int(ctx)
		if err != nil {
			return fmt.Errorf("create permissions count failed: %w", err)
		}

		bulkLength := length - current
		bulk := make([]*ent.PermissionCreate, bulkLength)
		for i := 0; i < bulkLength; i++ {
			bulk[i] = serv.DB.Permission.Create().SetID(uint16(num + i + 1)).SetAction(*toCreates[current+i].Action).SetDescription(*toCreates[current+i].Description)
		}

		err = tx.Permission.CreateBulk(bulk...).Exec(ctx)
		if err != nil {
			return rollback(tx, "create permissions", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("create permissions transaction commit failed: %w", err)
	}

	return nil

}

func (serv *ManagementService) UpdatePermission(toUpdate *entity.ToModifyPermission) error {
	ctx := context.Background()

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("update permission start a transaction failed: %w", err)
	}

	updater := tx.Permission.Update().Where(permission.IDEQ(toUpdate.ID), permission.DeletedTimeEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)))
	mutation := updater.Mutation()
	if toUpdate.Action != nil {
		mutation.SetAction(*toUpdate.Action)
	}
	if toUpdate.Description != nil {
		mutation.SetDescription(*toUpdate.Description)
	}
	if toUpdate.IsDisabled != nil {
		mutation.SetIsDisabled(*toUpdate.IsDisabled)
	}
	mutation.SetModifiedTime(time.Now())

	num, err := updater.Save(ctx)
	if err != nil {
		return rollback(tx, "update permission", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("update permission transaction commit failed: %w", err)
	}
	if num == 0 {
		return fmt.Errorf("update permission affect 0 row")
	}

	return nil
}

func (serv *ManagementService) RetrievePermission(current, pageSize int, like, sort string, order, isDisabled *bool) ([]*ent.Permission, error) {
	ctx := context.Background()

	offset := (current - 1) * pageSize

	permissions, err := serv.DB.Permission.Query().
		Where(
			permission.And(
				permission.Or(
					permission.ActionContains(like),
					permission.DescriptionContains(like),
				),
				func(s *sql.Selector) {
					if isDisabled != nil {
						s.Where(sql.EQ(permission.FieldIsDisabled, *isDisabled))
					}
				},
				permission.DeletedTimeEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
			),
		).
		Limit(pageSize).
		Offset(offset).
		Order(func(s *sql.Selector) {
			isSorted := sort != "" && (sort == permission.FieldID || sort == permission.FieldAction || sort == permission.FieldDescription || sort == permission.FieldIsDisabled)
			if isSorted && order != nil {
				if *order {
					s.OrderBy(sql.Desc(sort))
				} else {
					s.OrderBy(sql.Asc(sort))
				}
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

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("delete permissions start a transaction failed: %w", err)
	}

	num, err := tx.Permission.Update().
		Where(permission.And(
			permission.IDIn(toDeleteIDs...),
			permission.DeletedTimeEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
		)).
		SetDeletedTime(time.Now()).
		ClearRoles().
		Save(ctx)
	if err != nil {
		return rollback(tx, "delete permissions", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("delete permissions transaction commit failed: %w", err)
	}
	if num == 0 {
		return fmt.Errorf("delete permissions affect 0 row")
	}

	return nil
}

func (serv *ManagementService) CreateRole(toCreates []*entity.ToAddRole) error {
	ctx := context.Background()

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("create role start a transaction failed: %w", err)
	}

	current := 0
	length := len(toCreates)
	for ; current < length; current++ {
		id, err := tx.Role.Query().Where(role.DeletedTimeNEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local))).FirstID(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return fmt.Errorf("create role find a deleted role id query failed: %w", err)
			}
			break
		}
		num, err := tx.Role.Update().Where(role.And(
			role.IDEQ(id),
			role.DeletedTimeNEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
		)).
			SetCreatedTime(time.Now()).
			SetModifiedTime(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)).
			SetDeletedTime(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)).
			SetName(*toCreates[current].Name).
			SetDescription(*toCreates[current].Description).
			SetIsDisabled(false).
			Save(ctx)
		if err != nil {
			return rollback(tx, "create role", err)
		}
		if num == 0 {
			return rollback(tx, "create role", fmt.Errorf("create role update deleted role affect 0 row"))
		}

	}
	if current < length {
		num, err := tx.Role.Query().Aggregate(ent.Count()).Int(ctx)
		if err != nil {
			return fmt.Errorf("create roles count failed: %w", err)
		}

		bulkLength := length - current
		bulk := make([]*ent.RoleCreate, bulkLength)
		for i := 0; i < bulkLength; i++ {
			bulk[i] = serv.DB.Role.Create().SetID(uint8(num + i + 1)).SetName(*toCreates[current+i].Name).SetDescription(*toCreates[current+i].Name)
		}

		err = tx.Role.CreateBulk(bulk...).Exec(ctx)
		if err != nil {
			return rollback(tx, "create roles", err)
		}
	}

	// checkRepeatName, err := serv.DB.Role.Query().Where(role.NameIn(names...)).Exist(ctx)
	// if err != nil {
	// 	return fmt.Errorf("create roles check repeat role failed: %w", err)
	// }
	// if checkRepeatName {
	// 	return fmt.Errorf("repeat role")
	// }

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("create roles transaction commit failed: %w", err)
	}

	return nil
}

func (serv *ManagementService) UpdateRole(toUpdate *entity.ToModifyRole) error {
	ctx := context.Background()

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("update role start a transaction failed: %w", err)
	}

	updater := tx.Role.Update().Where(role.IDEQ(toUpdate.ID), role.DeletedTimeEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)))
	mutation := updater.Mutation()
	if toUpdate.Name != nil {
		mutation.SetName(*toUpdate.Name)
	}
	if toUpdate.Description != nil {
		mutation.SetDescription(*toUpdate.Description)
	}
	if toUpdate.IsDisabled != nil {
		mutation.SetIsDisabled(*toUpdate.IsDisabled)
	}
	mutation.SetModifiedTime(time.Now())

	num, err := updater.Save(ctx)
	if err != nil {
		return rollback(tx, "update role", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("update role transaction commit failed: %w", err)
	}
	if num == 0 {
		return fmt.Errorf("update role affect 0 row")
	}

	return nil
}

func (serv *ManagementService) RetrieveRole(current, pageSize int, like, sort string, order, isDisabled *bool) ([]*ent.Role, error) {
	ctx := context.Background()

	offset := (current - 1) * pageSize

	roles, err := serv.DB.Role.Query().
		Where(
			role.And(
				role.Or(
					role.NameContains(like),
					role.DescriptionContains(like),
				),
				func(s *sql.Selector) {
					if isDisabled != nil {
						s.Where(sql.EQ(role.FieldIsDisabled, *isDisabled))
					}
				},
				role.DeletedTimeEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
			),
		).
		Limit(pageSize).
		Offset(offset).
		Order(func(s *sql.Selector) {
			isSorted := sort != "" && (sort == role.FieldName || sort == role.FieldDescription || sort == role.FieldID)
			if isSorted && order != nil {
				if *order {
					s.OrderBy(sql.Desc(sort))
				} else {
					s.OrderBy(sql.Asc(sort))
				}
			}
		}).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("retrieve role query failed: %w", err)
	}

	return roles, nil

}

func (serv *ManagementService) GetTotalRetrievedRoles(like string, isDisabled *bool) (int, error) {
	ctx := context.Background()

	total, err := serv.DB.Role.Query().
		Where(
			role.And(
				role.Or(
					role.NameContains(like),
					role.DescriptionContains(like),
				),
				func(s *sql.Selector) {
					if isDisabled != nil {
						s.Where(sql.EQ(role.FieldIsDisabled, *isDisabled))
					}
				},
				role.DeletedTimeEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
			),
		).Count(ctx)

	if err != nil {
		return -1, fmt.Errorf("get retrieved total roles query failed: %w", err)
	}

	return total, nil
}

func (serv *ManagementService) DeleteRole(toDeleteIDs []uint8) error {
	ctx := context.Background()

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("delete roles start a transaction failed: %w", err)
	}

	originalRoles, err := tx.Role.Query().Where(role.And(
		role.IDIn(toDeleteIDs...),
		role.DeletedTimeEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
	)).All(ctx)
	if err != nil || len(originalRoles) != len(toDeleteIDs) {
		return fmt.Errorf("delete roles not found roles failed: %w", err)
	}

	for _, v := range originalRoles {
		num, err := tx.Role.Update().
			Where(role.And(
				role.IDEQ(v.ID),
				role.DeletedTimeEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
			)).
			SetDeletedTime(time.Now()).
			SetName("*" + v.Name).
			ClearMenu().ClearUsers().ClearPermissions().
			Save(ctx)
		if err != nil {
			return rollback(tx, "delete roles", err)
		}
		if num == 0 {
			return rollback(tx, "delete roles", fmt.Errorf("delete roles affect 0 row"))
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("delete roles transaction commit failed: %w", err)
	}

	return nil
}

func (serv *ManagementService) FindOneRoleById(id uint8) (*ent.Role, error) {
	ctx := context.Background()
	role, err := serv.DB.Role.Query().Where(role.IDEQ(id), role.DeletedTimeEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local))).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("role find one by id failed: %w", err)
	}
	return role, nil
}

func (serv *ManagementService) UpdateDeletedRole(toUpdate *entity.ToModifyRole) error {
	ctx := context.Background()

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("update deleted role start a transaction failed: %w", err)
	}

	num, err := tx.Role.Update().Where(role.And(
		role.IDEQ(toUpdate.ID),
		role.DeletedTimeNEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
	)).
		SetCreatedTime(time.Now()).
		SetModifiedTime(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)).
		SetDeletedTime(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)).
		SetName(*toUpdate.Name).
		SetDescription(*toUpdate.Description).
		SetIsDisabled(*toUpdate.IsDisabled).
		Save(ctx)
	if err != nil {
		return rollback(tx, "update deleted role", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("update deleted role transaction commit failed: %w", err)
	}
	if num == 0 {
		return fmt.Errorf("update deleted role affect 0 row")
	}

	return nil
}

func (serv *ManagementService) FindADeletedRoleID() (uint8, error) {
	ctx := context.Background()

	id, err := serv.DB.Role.Query().Where(role.DeletedTimeNEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local))).FirstID(ctx)
	if err != nil {
		return 0, fmt.Errorf("find a deleted role id query failed: %w", err)
	}
	return id, nil
}

func (serv *ManagementService) GetTotalRetrievedPermissions(like string, isDisabled *bool) (int, error) {
	ctx := context.Background()

	total, err := serv.DB.Permission.Query().
		Where(
			permission.And(
				permission.Or(
					permission.ActionContains(like),
					permission.DescriptionContains(like),
				),
				func(s *sql.Selector) {
					if isDisabled != nil {
						s.Where(sql.EQ(permission.FieldIsDisabled, *isDisabled))
					}
				},
				permission.DeletedTimeEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
			),
		).Count(ctx)

	if err != nil {
		return -1, fmt.Errorf("get retrieved total permissions query failed: %w", err)
	}

	return total, nil
}

func (serv *ManagementService) UpdateDeletedPermission(toUpdate *entity.ToModifyPermission) error {
	ctx := context.Background()

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("update deleted permission start a transaction failed: %w", err)
	}

	num, err := tx.Permission.Update().Where(permission.And(
		permission.IDEQ(toUpdate.ID),
		permission.DeletedTimeNEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
	)).
		SetCreatedTime(time.Now()).
		SetModifiedTime(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)).
		SetDeletedTime(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)).
		SetAction(*toUpdate.Action).
		SetDescription(*toUpdate.Description).
		SetIsDisabled(*toUpdate.IsDisabled).
		Save(ctx)
	if err != nil {
		return rollback(tx, "update deleted permission", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("update deleted permission transaction commit failed: %w", err)
	}
	if num == 0 {
		return fmt.Errorf("update deleted permission affect 0 row")
	}

	return nil
}

func (serv *ManagementService) FindADeletedPermissionID() (uint16, error) {
	ctx := context.Background()

	id, err := serv.DB.Permission.Query().Where(permission.DeletedTimeNEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local))).FirstID(ctx)
	if err != nil {
		return 0, fmt.Errorf("find a deleted permission id query failed: %w", err)
	}
	return id, nil
}

func (serv *ManagementService) FindOnePermissionById(id uint16) (*ent.Permission, error) {
	ctx := context.Background()
	p, err := serv.DB.Permission.Query().Where(
		permission.IDEQ(id),
		permission.DeletedTimeEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
	).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("permission find one by id failed: %w", err)
	}
	return p, nil
}

func (serv *ManagementService) FindPermissionsByRoleId(id uint8) ([]*ent.Permission, error) {
	ctx := context.Background()
	ps, err := serv.DB.Role.Query().Where(role.And(
		role.IDEQ(id),
		role.DeletedTimeEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
	)).QueryPermissions().Where(permission.And(
		permission.DeletedTimeEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
		permission.IsDisabledEQ(false),
	)).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("permission find many by role id failed: %w", err)
	}
	return ps, nil
}

func (serv *ManagementService) UpdatePermissionForRole(rids []uint8, pids []uint16, isDeleted bool) error {
	ctx := context.Background()

	ps, err := serv.DB.Permission.Query().Where(
		permission.DeletedTimeEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
		permission.IsDisabledEQ(false),
		permission.IDIn(pids...),
	).All(ctx)
	if err != nil {
		return fmt.Errorf("update permissions for role failed: %w", err)
	}
	if len(ps) != len(pids) {
		return fmt.Errorf("update permissions for role failed: not enough matching permissions")
	}

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("update permissions for role start a transaction failed: %w", err)
	}

	updater := tx.Role.Update().Where(
		role.DeletedTimeEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
		role.IsDisabledEQ(false),
		role.IDIn(rids...),
	)
	mutation := updater.Mutation()
	if isDeleted {
		mutation.RemovePermissionIDs(pids...)
	} else {
		mutation.AddPermissionIDs(pids...)
	}

	num, err := updater.Save(ctx)
	if err != nil {
		return rollback(tx, "update permission for role", err)
	}
	if num == 0 || num != len(rids) {
		return rollback(tx, "update permission for role", fmt.Errorf("not enough updated roles"))
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("update permissions for role transaction commit failed: %w", err)
	}

	return nil

}

func (serv *ManagementService) RetrieveUser(current, pageSize int, like, sort string, order, isDisabled *bool) ([]*ent.User, error) {
	ctx := context.Background()

	offset := (current - 1) * pageSize

	users, err := serv.DB.User.Query().Where(user.And(
		user.Or(
			user.AccountContains(like),
			user.UsernameContains(like),
			user.EmailContains(like),
			user.PhoneContains(like),
		),
		func(s *sql.Selector) {
			if isDisabled != nil {
				s.Where(sql.EQ(user.FieldIsDisabled, *isDisabled))
			}
		},
		user.DeletedTimeEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
	)).
		Limit(pageSize).
		Offset(offset).
		Order(func(s *sql.Selector) {
			isSorted := sort != "" && (sort == user.FieldID || sort == user.FieldAccount || sort == user.FieldUsername || sort == user.FieldEmail || sort == user.FieldPhone || sort == user.FieldIsDisabled)
			if isSorted && order != nil {
				if *order {
					s.OrderBy(sql.Desc(sort))
				} else {
					s.OrderBy(sql.Asc(sort))
				}
			}
		}).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("retrieve user query failed: %w", err)
	}
	return users, nil
}

func (serv *ManagementService) GetTotalRetrievedUsers(like string, isDisabled *bool) (int, error) {
	ctx := context.Background()

	total, err := serv.DB.User.Query().Where(user.And(
		user.Or(
			user.AccountContains(like),
			user.UsernameContains(like),
			user.EmailContains(like),
			user.PhoneContains(like),
		),
		func(s *sql.Selector) {
			if isDisabled != nil {
				s.Where(sql.EQ(user.FieldIsDisabled, *isDisabled))
			}
		},
		user.DeletedTimeEQ(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
	)).Count(ctx)

	if err != nil {
		return -1, fmt.Errorf("get retrieved total users query failed: %w", err)
	}

	return total, nil
}
