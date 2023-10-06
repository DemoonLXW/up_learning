package service

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"reflect"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/database/ent/class"
	"github.com/DemoonLXW/up_learning/database/ent/college"
	"github.com/DemoonLXW/up_learning/database/ent/file"
	"github.com/DemoonLXW/up_learning/database/ent/major"
	"github.com/DemoonLXW/up_learning/database/ent/permission"
	"github.com/DemoonLXW/up_learning/database/ent/role"
	"github.com/DemoonLXW/up_learning/database/ent/samplefile"
	"github.com/DemoonLXW/up_learning/database/ent/school"
	"github.com/DemoonLXW/up_learning/database/ent/student"
	"github.com/DemoonLXW/up_learning/database/ent/teacher"
	"github.com/DemoonLXW/up_learning/database/ent/user"
	"github.com/DemoonLXW/up_learning/entity"
	"github.com/xuri/excelize/v2"
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
		id, err := tx.Permission.Query().Where(func(s *sql.Selector) {
			s.Where(sql.NotNull(permission.FieldDeletedTime))
		}).FirstID(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return fmt.Errorf("create permission find a deleted permission id query failed: %w", err)
			}
			break
		}

		num, err := tx.Permission.Update().Where(permission.And(
			permission.IDEQ(id),
			func(s *sql.Selector) {
				s.Where(sql.NotNull(permission.FieldDeletedTime))
			},
		)).
			SetCreatedTime(time.Now()).
			ClearModifiedTime().
			ClearDeletedTime().
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
			bulk[i] = tx.Permission.Create().SetID(uint16(num + i + 1)).SetAction(*toCreates[current+i].Action).SetDescription(*toCreates[current+i].Description)
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

	updater := tx.Permission.Update().Where(
		permission.IDEQ(toUpdate.ID),
		func(s *sql.Selector) {
			s.Where(sql.IsNull(permission.FieldDeletedTime))
		},
	)
	mutation := updater.Mutation()
	// Action can not be modified, should be fixed
	// if toUpdate.Action != nil {
	// 	mutation.SetAction(*toUpdate.Action)
	// }
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

func (serv *ManagementService) RetrievePermission(current, pageSize *int, like, sort string, order, isDisabled *bool) ([]*ent.Permission, error) {
	ctx := context.Background()

	query := serv.DB.Permission.Query().
		Where(
			permission.And(
				permission.Or(
					permission.ActionContains(like),
					permission.DescriptionContains(like),
				),
				func(s *sql.Selector) {
					s.Where(sql.IsNull(permission.FieldDeletedTime))
					if isDisabled != nil {
						s.Where(sql.EQ(permission.FieldIsDisabled, *isDisabled))
					}
				},
			),
		).
		Order(func(s *sql.Selector) {
			isSorted := sort != "" && (sort == permission.FieldID || sort == permission.FieldAction || sort == permission.FieldDescription || sort == permission.FieldIsDisabled)
			if isSorted && order != nil {
				if *order {
					s.OrderBy(sql.Desc(sort))
				} else {
					s.OrderBy(sql.Asc(sort))
				}
			}
		})

	if pageSize != nil && current != nil {
		offset := (*current - 1) * (*pageSize)
		query.Limit(*pageSize).Offset(offset)
	}
	permissions, err := query.All(ctx)

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

	original, err := tx.Permission.Query().Where(permission.And(
		permission.IDIn(toDeleteIDs...),
		func(s *sql.Selector) {
			s.Where(sql.IsNull(permission.FieldDeletedTime))
		},
	)).All(ctx)
	if err != nil || len(original) != len(toDeleteIDs) {
		return fmt.Errorf("delete permissions not found permissions failed: %w", err)
	}

	for _, v := range original {
		num, err := tx.Permission.Update().
			Where(permission.And(
				permission.IDEQ(v.ID),
				func(s *sql.Selector) {
					s.Where(sql.IsNull(permission.FieldDeletedTime))
				},
			)).
			SetDeletedTime(time.Now()).
			ClearModifiedTime().
			SetAction("*" + v.Action).
			ClearRoles().
			Save(ctx)
		if err != nil {
			return rollback(tx, "delete permissions", err)
		}
		if num == 0 {
			return rollback(tx, "delete permissions", fmt.Errorf("delete permissions affect 0 row"))
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("delete permissions transaction commit failed: %w", err)
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
		id, err := tx.Role.Query().Where(func(s *sql.Selector) {
			s.Where(sql.NotNull(role.FieldDeletedTime))
		},
		).FirstID(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return fmt.Errorf("create role find a deleted role id query failed: %w", err)
			}
			break
		}
		num, err := tx.Role.Update().Where(role.And(
			role.IDEQ(id),
			func(s *sql.Selector) {
				s.Where(sql.NotNull(role.FieldDeletedTime))
			},
		)).
			SetCreatedTime(time.Now()).
			ClearModifiedTime().
			ClearDeletedTime().
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

	updater := tx.Role.Update().Where(
		role.IDEQ(toUpdate.ID),
		func(s *sql.Selector) {
			s.Where(sql.IsNull(role.FieldDeletedTime))
		},
	)
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

func (serv *ManagementService) RetrieveRole(current, pageSize *int, like, sort string, order, isDisabled *bool) ([]*ent.Role, error) {
	ctx := context.Background()

	query := serv.DB.Role.Query().
		Where(
			role.And(
				role.Or(
					role.NameContains(like),
					role.DescriptionContains(like),
				),
				func(s *sql.Selector) {
					s.Where(sql.IsNull(role.FieldDeletedTime))
					if isDisabled != nil {
						s.Where(sql.EQ(role.FieldIsDisabled, *isDisabled))
					}
				},
			),
		).
		Order(func(s *sql.Selector) {
			isSorted := sort != "" && (sort == role.FieldName || sort == role.FieldDescription || sort == role.FieldID)
			if isSorted && order != nil {
				if *order {
					s.OrderBy(sql.Desc(sort))
				} else {
					s.OrderBy(sql.Asc(sort))
				}
			}
		})

	if pageSize != nil && current != nil {
		offset := (*current - 1) * (*pageSize)
		query.Limit(*pageSize).Offset(offset)
	}
	roles, err := query.All(ctx)

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
					s.Where(sql.IsNull(role.FieldDeletedTime))
					if isDisabled != nil {
						s.Where(sql.EQ(role.FieldIsDisabled, *isDisabled))
					}
				},
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

	original, err := tx.Role.Query().Where(role.And(
		role.IDIn(toDeleteIDs...),
		func(s *sql.Selector) {
			s.Where(sql.IsNull(role.FieldDeletedTime))
		},
	)).All(ctx)
	if err != nil || len(original) != len(toDeleteIDs) {
		return fmt.Errorf("delete roles not found roles failed: %w", err)
	}

	for _, v := range original {
		num, err := tx.Role.Update().
			Where(role.And(
				role.IDEQ(v.ID),
				func(s *sql.Selector) {
					s.Where(sql.IsNull(role.FieldDeletedTime))
				},
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
	role, err := serv.DB.Role.Query().Where(
		role.IDEQ(id),
		func(s *sql.Selector) {
			s.Where(sql.IsNull(role.FieldDeletedTime))
		},
	).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("role find one by id failed: %w", err)
	}
	return role, nil
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
					s.Where(sql.IsNull(permission.FieldDeletedTime))
					if isDisabled != nil {
						s.Where(sql.EQ(permission.FieldIsDisabled, *isDisabled))
					}
				},
			),
		).Count(ctx)

	if err != nil {
		return -1, fmt.Errorf("get retrieved total permissions query failed: %w", err)
	}

	return total, nil
}

func (serv *ManagementService) FindOnePermissionById(id uint16) (*ent.Permission, error) {
	ctx := context.Background()
	p, err := serv.DB.Permission.Query().Where(
		permission.IDEQ(id),
		func(s *sql.Selector) {
			s.Where(sql.IsNull(permission.FieldDeletedTime))
		},
	).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("permission find one by id failed: %w", err)
	}
	return p, nil
}

func (serv *ManagementService) FindPermissionsByRoleIds(ids []uint8) ([]*ent.Permission, error) {
	ctx := context.Background()
	rs, err := serv.DB.Role.Query().Where(role.And(
		role.IDIn(ids...),
		func(s *sql.Selector) {
			s.Where(sql.IsNull(role.FieldDeletedTime))
		},
	)).WithPermissions(func(q *ent.PermissionQuery) {
		q.Where(permission.And(
			func(s *sql.Selector) {
				s.Where(sql.IsNull(permission.FieldDeletedTime))
			},
		))
	}).All(ctx)

	if len(rs) != len(ids) {
		return nil, fmt.Errorf("permission find many by role id failed: %w", fmt.Errorf("not enough retrieve roles"))
	}
	if err != nil {
		return nil, fmt.Errorf("permission find many by role id failed: %w", err)
	}

	var ps []*ent.Permission
	for _, v := range rs {
		ps = append(ps, v.Edges.Permissions...)
	}

	return ps, nil
}

func (serv *ManagementService) UpdatePermissionForRole(rids []uint8, pids []uint16) error {
	ctx := context.Background()

	ps, err := serv.DB.Permission.Query().Where(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(permission.FieldDeletedTime))
		},
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

	// updater := tx.Role.Update().Where(
	// 	func(s *sql.Selector) {
	// s.Where(sql.IsNull(role.FieldDeletedTime))
	// },
	// 	role.IsDisabledEQ(false),
	// 	role.IDIn(rids...),
	// )
	// mutation := updater.Mutation()
	// if isDeleted {
	// 	mutation.RemovePermissionIDs(pids...)
	// } else {
	// 	mutation.AddPermissionIDs(pids...)
	// }

	// num, err := updater.Save(ctx)
	// if err != nil {
	// 	return rollback(tx, "update permission for role", err)
	// }
	// if num != len(rids) {
	// 	return rollback(tx, "update permission for role", fmt.Errorf("not enough updated roles"))
	// }

	num, err := tx.Role.Update().Where(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(role.FieldDeletedTime))
		},
		role.IsDisabledEQ(false),
		role.IDIn(rids...),
	).ClearPermissions().Save(ctx)
	if err != nil {
		return rollback(tx, "update permission for role -> clear permissions", err)
	}
	if num != len(rids) {
		return rollback(tx, "update permission for role -> clear permissions", fmt.Errorf("not enough updated roles"))
	}

	num, err = tx.Role.Update().Where(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(role.FieldDeletedTime))
		},
		role.IsDisabledEQ(false),
		role.IDIn(rids...),
	).AddPermissionIDs(pids...).Save(ctx)
	if err != nil {
		return rollback(tx, "update permission for role -> add permissions' ids", err)
	}
	if num != len(rids) {
		return rollback(tx, "update permission for role -> add permissions' ids", fmt.Errorf("not enough updated roles"))
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("update permissions for role transaction commit failed: %w", err)
	}

	return nil

}

func (serv *ManagementService) RetrieveUser(current, pageSize *int, like, sort string, order, isDisabled *bool) ([]*ent.User, error) {
	ctx := context.Background()

	query := serv.DB.User.Query().Where(user.And(
		user.Or(
			user.AccountContains(like),
			user.UsernameContains(like),
			user.EmailContains(like),
			user.PhoneContains(like),
		),
		func(s *sql.Selector) {
			s.Where(sql.IsNull(user.FieldDeletedTime))
			if isDisabled != nil {
				s.Where(sql.EQ(user.FieldIsDisabled, *isDisabled))
			}
		},
	)).
		Order(func(s *sql.Selector) {
			isSorted := sort != "" && (sort == user.FieldID || sort == user.FieldAccount || sort == user.FieldUsername || sort == user.FieldEmail || sort == user.FieldPhone || sort == user.FieldIsDisabled)
			if isSorted && order != nil {
				if *order {
					s.OrderBy(sql.Desc(sort))
				} else {
					s.OrderBy(sql.Asc(sort))
				}
			}
		})

	if pageSize != nil && current != nil {
		offset := (*current - 1) * (*pageSize)
		query.Limit(*pageSize).Offset(offset)
	}
	users, err := query.All(ctx)
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
			s.Where(sql.IsNull(user.FieldDeletedTime))
			if isDisabled != nil {
				s.Where(sql.EQ(user.FieldIsDisabled, *isDisabled))
			}
		},
	)).Count(ctx)

	if err != nil {
		return -1, fmt.Errorf("get retrieved total users query failed: %w", err)
	}

	return total, nil
}

func (serv *ManagementService) CreateUser(toCreates []*entity.ToAddUser, roleId uint8) error {
	ctx := context.Background()

	r, err := serv.DB.Role.Query().Where(role.And(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(role.FieldDeletedTime))
		},
		role.IsDisabledEQ(false),
		role.IDEQ(roleId),
	)).First(ctx)
	if err != nil {
		return fmt.Errorf("create user found role by id failed: %w", err)
	}

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("create user start a transaction failed: %w", err)
	}

	current := 0
	length := len(toCreates)

	for ; current < length; current++ {
		id, err := tx.User.Query().Where(func(s *sql.Selector) {
			s.Where(sql.NotNull(user.FieldDeletedTime))
		}).FirstID(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return fmt.Errorf("create user find a deleted user id query failed: %w", err)
			}
			break
		}

		num, err := tx.User.Update().Where(user.And(
			user.IDEQ(id),
			func(s *sql.Selector) {
				s.Where(sql.NotNull(user.FieldDeletedTime))
			},
		)).
			SetCreatedTime(time.Now()).
			ClearModifiedTime().
			ClearDeletedTime().
			SetAccount(*toCreates[current].Account).
			SetUsername(*toCreates[current].Username).
			SetIntroduction("").
			SetIsDisabled(false).
			SetPassword(*toCreates[current].Password).
			AddRoles(r).
			Save(ctx)

		if err != nil {
			return rollback(tx, "create user", err)
		}
		if num == 0 {
			return rollback(tx, "create user", fmt.Errorf("create user update deleted user affect 0 row"))
		}
	}
	if current < length {
		num, err := tx.User.Query().Aggregate(ent.Count()).Int(ctx)
		if err != nil {
			return fmt.Errorf("create users count failed: %w", err)
		}

		bulkLength := length - current
		// bulk := make([]*ent.UserCreate, bulkLength)
		// for i := 0; i < bulkLength; i++ {
		// 	bulk[i] = tx.User.Create().SetID(uint32(num + i + 1)).
		// 		SetAccount(*toCreates[current+i].Account).
		// 		SetUsername(*toCreates[current+i].Username).
		// 		SetPassword(*toCreates[current+i].Password).
		// 		SetIntroduction("").
		// 		AddRoles(r)
		// }
		for i := 0; i < bulkLength; i++ {
			err = tx.User.Create().SetID(uint32(num + i + 1)).
				SetAccount(*toCreates[current+i].Account).
				SetUsername(*toCreates[current+i].Username).
				SetPassword(*toCreates[current+i].Password).
				SetIntroduction("").
				AddRoles(r).
				Exec(ctx)
			if err != nil {
				return rollback(tx, "create users", err)
			}
		}

		// _, err = tx.User.CreateBulk(bulk...).Save(ctx)
		// if err != nil {
		// 	return rollback(tx, "create users", err)
		// }

	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("create users transaction commit failed: %w", err)
	}

	return nil
}

func (serv *ManagementService) SaveImportedFile(file *multipart.FileHeader, dir, prefix string) (*os.File, error) {
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("save imported file open file failed: %w", err)
	}
	defer src.Close()

	if err = os.MkdirAll(dir, 0750); err != nil {
		return nil, fmt.Errorf("save imported file mkdir failed: %w", err)
	}

	out, err := os.CreateTemp(dir, prefix)
	if err != nil {
		return nil, fmt.Errorf("save imported file create temp failed: %w", err)
	}

	_, err = io.Copy(out, src)
	if err != nil {
		return nil, fmt.Errorf("save imported file copy failed: %w", err)
	}
	defer out.Close()

	return out, nil
}

func (serv *ManagementService) ReadSchoolsFromFile(f *os.File) ([]*entity.ToAddSchool, error) {
	book, err := excelize.OpenFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("read schools from file open failed: %w", err)
	}

	rows, err := book.GetRows(book.GetSheetName(0))
	if err != nil {
		return nil, fmt.Errorf("read schools from file get rows failed: %w", err)
	}

	length := len(rows)
	if length < 2 {
		return nil, fmt.Errorf("read schools from file less than two rows failed: %w", err)
	}

	schools := make([]*entity.ToAddSchool, length-1)
	columnMap := map[int]string{0: "Name", 1: "Code", 2: "CompetentDepartment", 3: "Location", 4: "EducationLevel", 5: "Remark"}
	fieldMap := make(map[string]int)
	v := reflect.ValueOf(&entity.ToAddSchool{}).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldMap[t.Field(i).Name] = i
	}

	for i := 1; i < length; i++ {
		var school entity.ToAddSchool
		schoolValue := reflect.ValueOf(&school).Elem()

		size := len(rows[i])
		if size < 5 {
			return nil, fmt.Errorf("read schools from file less than five columns failed: %w", err)
		}

		for j := 0; j < 5; j++ {
			col := strings.Trim(rows[i][j], " ")
			if col == "" {
				return nil, fmt.Errorf("read schools from file cell[%d][%d] %s empty failed: %w", i, j, columnMap[j], err)
			}
			switch j {
			case 4:
				{
					var level uint8
					switch col {
					case "本科":
						level = 0
					case "专科":
						level = 1
					default:
						return nil, fmt.Errorf("read schools from file cell[%d][%d] %s invalid failed: %w", i, j, columnMap[j], err)
					}
					schoolValue.Field(fieldMap[columnMap[j]]).Set(reflect.ValueOf(level))

				}
			default:
				schoolValue.Field(fieldMap[columnMap[j]]).Set(reflect.ValueOf(col))
			}

		}

		if size == 6 && strings.Trim(rows[i][5], " ") != "" {
			school.Remark = rows[i][5]
		}

		// note that i
		schools[i-1] = &school
	}

	return schools, nil
}

func (serv *ManagementService) CreateSchool(toCreates []*entity.ToAddSchool) error {
	ctx := context.Background()

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("create school start a transaction failed: %w", err)
	}

	current := 0
	length := len(toCreates)
	for ; current < length; current++ {
		id, err := tx.School.Query().Where(func(s *sql.Selector) {
			s.Where(sql.NotNull(school.FieldDeletedTime))
		}).FirstID(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return fmt.Errorf("create school find a deleted school id query failed: %w", err)
			}
			break
		}

		num, err := tx.School.Update().Where(school.And(
			school.IDEQ(id),
			func(s *sql.Selector) {
				s.Where(sql.NotNull(school.FieldDeletedTime))
			},
		)).
			SetCreatedTime(time.Now()).
			ClearModifiedTime().
			ClearDeletedTime().
			SetIsDisabled(false).
			SetCode(toCreates[current].Code).
			SetName(toCreates[current].Name).
			SetCompetentDepartment(toCreates[current].CompetentDepartment).
			SetLocation(toCreates[current].Location).
			SetEducationLevel(toCreates[current].EducationLevel).
			SetRemark(toCreates[current].Remark).
			Save(ctx)
		if err != nil {
			return rollback(tx, "create school", err)
		}
		if num == 0 {
			return rollback(tx, "create school", fmt.Errorf("create school update deleted school affect 0 row"))
		}
	}
	if current < length {
		num, err := tx.School.Query().Aggregate(ent.Count()).Int(ctx)
		if err != nil {
			return fmt.Errorf("create school count failed: %w", err)
		}

		bulkLength := length - current
		bulk := make([]*ent.SchoolCreate, bulkLength)
		for i := 0; i < bulkLength; i++ {
			bulk[i] = tx.School.Create().SetID(uint16(num + i + 1)).
				SetCode(toCreates[current+i].Code).
				SetName(toCreates[current+i].Name).
				SetCompetentDepartment(toCreates[current+i].CompetentDepartment).
				SetLocation(toCreates[current+i].Location).
				SetEducationLevel(toCreates[current+i].EducationLevel).
				SetRemark(toCreates[current+i].Remark)
		}

		err = tx.School.CreateBulk(bulk...).Exec(ctx)
		if err != nil {
			return rollback(tx, "create school", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("create school transaction commit failed: %w", err)
	}

	return nil
}

func (serv *ManagementService) RetrieveSchool(current, pageSize *int, like, sort string, order, isDisabled *bool) ([]*ent.School, error) {
	ctx := context.Background()

	query := serv.DB.School.Query().
		Where(school.And(
			school.Or(
				school.NameContains(like),
			),
			func(s *sql.Selector) {
				s.Where(sql.IsNull(school.FieldDeletedTime))
				if isDisabled != nil {
					s.Where(sql.EQ(school.FieldIsDisabled, *isDisabled))
				}
			},
		)).
		Order(func(s *sql.Selector) {
			isSorted := sort != "" && (sort == school.FieldID || sort == school.FieldCode || sort == school.FieldName ||
				sort == school.FieldLocation || sort == school.FieldCompetentDepartment || sort == school.FieldIsDisabled)
			if isSorted && order != nil {
				if *order {
					s.OrderBy(sql.Desc(sort))
				} else {
					s.OrderBy(sql.Asc(sort))
				}
			}
		})

	if pageSize != nil && current != nil {
		offset := (*current - 1) * (*pageSize)
		query.Limit(*pageSize).Offset(offset)
	}
	schools, err := query.All(ctx)

	if err != nil {
		return nil, fmt.Errorf("retrieve school query failed: %w", err)
	}

	return schools, nil
}

func (serv *ManagementService) GetTotalRetrievedSchools(like string, isDisabled *bool) (int, error) {
	ctx := context.Background()

	total, err := serv.DB.School.Query().
		Where(school.And(
			school.Or(
				school.NameContains(like),
			),
			func(s *sql.Selector) {
				s.Where(sql.IsNull(school.FieldDeletedTime))
				if isDisabled != nil {
					s.Where(sql.EQ(school.FieldIsDisabled, *isDisabled))
				}
			},
		)).Count(ctx)

	if err != nil {
		return -1, fmt.Errorf("get retrieved total schools query failed: %w", err)
	}

	return total, nil
}

func (serv *ManagementService) FindOneSampleFileByType(t string) (*ent.File, error) {
	ctx := context.Background()

	f, err := serv.DB.SampleFile.Query().Where(samplefile.And(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(samplefile.FieldDeletedTime))
		},
		samplefile.TypeEQ(t),
	)).
		QueryFile().Where(file.And(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(file.FieldDeletedTime))
		},
	)).
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf("sample file find one by type failed: %w", err)
	}

	return f, nil
}

func (serv *ManagementService) ReadStudentsFromFile(f *os.File) ([]*entity.ToAddStudent, error) {
	book, err := excelize.OpenFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("read students from file open failed: %w", err)
	}

	rows, err := book.GetRows(book.GetSheetName(0))
	if err != nil {
		return nil, fmt.Errorf("read students from file get rows failed: %w", err)
	}

	length := len(rows)
	if length < 2 {
		return nil, fmt.Errorf("read students from file less than two rows failed: %w", err)
	}

	columnCheck := []string{"学号", "姓名", "性别"}
	columnSize := len(columnCheck)
	if len(rows[0]) < columnSize {
		return nil, fmt.Errorf("read students from file header less than %d", columnSize)
	}
	for i, v := range columnCheck {
		if rows[0][i] != v {
			return nil, fmt.Errorf("read students from file header[%d] %s does not exist", i, v)
		}
	}

	students := make([]*entity.ToAddStudent, length-1)
	columnMap := map[int]string{0: "StudentID", 1: "Name", 2: "Gender"}
	fieldMap := make(map[string]int)
	v := reflect.ValueOf(&entity.ToAddStudent{}).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldMap[t.Field(i).Name] = i
	}

	for i := 1; i < length; i++ {
		var student entity.ToAddStudent
		studentValue := reflect.ValueOf(&student).Elem()

		if len(rows[i]) < columnSize {
			return nil, fmt.Errorf("read students from file less than %d columns", columnSize)
		}

		for j := 0; j < 3; j++ {
			col := strings.Trim(rows[i][j], "")
			if col == "" {
				return nil, fmt.Errorf("read students from file cell[%d][%d] %s empty", i, j, columnMap[j])
			}
			switch j {
			case 2:
				{
					var gender uint8
					switch col {
					case "男":
						gender = 0
					case "女":
						gender = 1
					default:
						return nil, fmt.Errorf("read students from file cell[%d][%d] %s invalid", i, j, columnMap[j])
					}
					studentValue.Field(fieldMap[columnMap[j]]).Set(reflect.ValueOf(gender))

				}
			default:
				studentValue.Field(fieldMap[columnMap[j]]).Set(reflect.ValueOf(col))
			}
		}

		students[i-1] = &student
	}

	return students, nil
}

func (serv *ManagementService) CreateStudentByClassIDAndCreateUser(toCreates []*entity.ToAddStudent, classID uint32) error {
	ctx := context.Background()

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("create student start a transaction failed: %w", err)
	}

	current := 0
	length := len(toCreates)
	for ; current < length; current++ {
		id, err := tx.Student.Query().Where(func(s *sql.Selector) {
			s.Where(sql.NotNull(student.FieldDeletedTime))
		}).FirstID(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return fmt.Errorf("create student find a deleted student id query failed: %w", err)
			}
			break
		}

		num, err := tx.Student.Update().Where(student.And(
			student.IDEQ(id),
			func(s *sql.Selector) {
				s.Where(sql.NotNull(student.FieldDeletedTime))
			},
		)).
			SetCreatedTime(time.Now()).
			ClearDeletedTime().
			ClearModifiedTime().
			SetIsDisabled(false).
			SetName(toCreates[current].Name).
			SetStudentID(toCreates[current].StudentID).
			SetGender(toCreates[current].Gender).
			SetCid(classID).
			Save(ctx)
		if err != nil {
			return rollback(tx, "create student", err)
		}
		if num == 0 {
			return rollback(tx, "create student", fmt.Errorf("create student update deleted student affect 0 row"))
		}

	}
	if current < length {
		num, err := tx.Student.Query().Aggregate(ent.Count()).Int(ctx)
		if err != nil {
			return fmt.Errorf("create student count failed: %w", err)
		}

		bulkLength := length - current
		bulk := make([]*ent.StudentCreate, bulkLength)
		for i := 0; i < bulkLength; i++ {
			bulk[i] = tx.Student.Create().SetID(uint32(num + i + 1)).
				SetName(toCreates[current+i].Name).
				SetStudentID(toCreates[current+i].StudentID).
				SetGender(toCreates[current+i].Gender).
				SetCid(classID)
		}

		err = tx.Student.CreateBulk(bulk...).Exec(ctx)
		if err != nil {
			return rollback(tx, "create student", err)
		}
	}

	// create user for student

	r, err := tx.Role.Query().Where(role.And(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(role.FieldDeletedTime))
		},
		role.NameEQ("student"),
	)).First(ctx)
	if err != nil {
		return rollback(tx, "create user", fmt.Errorf("create user found role by name failed: %w", err))
	}

	studentIDs := make([]string, length)
	for i, v := range toCreates {
		studentIDs[i] = v.StudentID
	}
	s, err := tx.Student.Query().Where(student.And(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(student.FieldDeletedTime))
		},
		student.StudentIDIn(studentIDs...),
	)).All(ctx)
	if err != nil {
		return rollback(tx, "create user", fmt.Errorf("create user found student by student ids failed: %w", err))
	}
	if len(s) != length {
		return rollback(tx, "create user", fmt.Errorf("create user found student by student ids not enough failed"))
	}

	original_pwd := "123456"
	md5_pwd := fmt.Sprintf("%x", md5.Sum([]byte(original_pwd)))
	sha256_pwd := fmt.Sprintf("%x", sha256.Sum256([]byte(md5_pwd)))

	current = 0
	for ; current < length; current++ {
		id, err := tx.User.Query().Where(func(s *sql.Selector) {
			s.Where(sql.NotNull(user.FieldDeletedTime))
		}).FirstID(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return fmt.Errorf("create user find a deleted user id query failed: %w", err)
			}
			break
		}

		num, err := tx.User.Update().Where(user.And(
			user.IDEQ(id),
			func(s *sql.Selector) {
				s.Where(sql.NotNull(user.FieldDeletedTime))
			},
		)).
			SetCreatedTime(time.Now()).
			ClearModifiedTime().
			ClearDeletedTime().
			SetAccount(s[current].StudentID).
			SetUsername(s[current].Name).
			SetIntroduction("").
			SetIsDisabled(false).
			SetPassword(sha256_pwd).
			AddRoles(r).
			SetStudentID(s[current].ID).
			Save(ctx)

		if err != nil {
			return rollback(tx, "create user", err)
		}
		if num == 0 {
			return rollback(tx, "create user", fmt.Errorf("create user update deleted user affect 0 row"))
		}
	}
	if current < length {
		num, err := tx.User.Query().Aggregate(ent.Count()).Int(ctx)
		if err != nil {
			return fmt.Errorf("create users count failed: %w", err)
		}

		bulkLength := length - current
		for i := 0; i < bulkLength; i++ {
			err = tx.User.Create().SetID(uint32(num + i + 1)).
				SetAccount(s[current+i].StudentID).
				SetUsername(s[current+i].Name).
				SetPassword(sha256_pwd).
				SetIntroduction("").
				AddRoles(r).
				SetStudentID(s[current+i].ID).
				Exec(ctx)
			if err != nil {
				return rollback(tx, "create users", err)
			}
		}

	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("create student transaction commit failed: %w", err)
	}

	return nil
}

func (serv *ManagementService) RetrieveStudentWithClassAndUser(current, pageSize *int, like, sort string, order, isDisabled *bool) ([]*ent.Student, error) {
	ctx := context.Background()

	query := serv.DB.Student.Query().
		Where(student.And(
			student.Or(
				student.NameContains(like),
				student.StudentIDContains(like),
			),
			func(s *sql.Selector) {
				s.Where(sql.IsNull(student.FieldDeletedTime))
				if isDisabled != nil {
					s.Where(sql.EQ(student.FieldIsDisabled, *isDisabled))
				}
			},
		)).
		Order(func(s *sql.Selector) {
			isSorted := sort != "" && (sort == student.FieldID || sort == student.FieldStudentID || sort == student.FieldName ||
				sort == student.FieldGender || sort == student.FieldIsDisabled)
			if isSorted && order != nil {
				if *order {
					s.OrderBy(sql.Desc(sort))
				} else {
					s.OrderBy(sql.Asc(sort))
				}
			}
		}).
		WithClass(func(cq *ent.ClassQuery) {
			cq.Where(func(s *sql.Selector) {
				s.Where(sql.IsNull(class.FieldDeletedTime))
			})
		}).
		WithUser(func(uq *ent.UserQuery) {
			uq.Where(func(s *sql.Selector) {
				s.Where(sql.IsNull(user.FieldDeletedTime))
			})
		})

	if pageSize != nil && current != nil {
		offset := (*current - 1) * (*pageSize)
		query.Limit(*pageSize).Offset(offset)
	}
	students, err := query.All(ctx)

	if err != nil {
		return nil, fmt.Errorf("retrieve student query failed: %w", err)
	}

	return students, nil
}

func (serv *ManagementService) GetTotalRetrievedStudents(like string, isDisabled *bool) (int, error) {
	ctx := context.Background()

	total, err := serv.DB.Student.Query().
		Where(student.And(
			student.Or(
				student.NameContains(like),
				student.StudentIDContains(like),
			),
			func(s *sql.Selector) {
				s.Where(sql.IsNull(student.FieldDeletedTime))
				if isDisabled != nil {
					s.Where(sql.EQ(student.FieldIsDisabled, *isDisabled))
				}
			},
		)).Count(ctx)

	if err != nil {
		return -1, fmt.Errorf("get retrieved total students query failed: %w", err)
	}

	return total, nil
}

func (serv *ManagementService) ReadCollegesFromFile(f *os.File) ([]*entity.ToAddCollege, error) {
	book, err := excelize.OpenFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("read colleges from file open failed: %w", err)
	}

	rows, err := book.GetRows(book.GetSheetName(0))
	if err != nil {
		return nil, fmt.Errorf("read colleges from file get rows failed: %w", err)
	}

	length := len(rows)
	if length < 2 {
		return nil, fmt.Errorf("read colleges from file less than two rows failed: %w", err)
	}

	columnCheck := []string{"学院名称"}
	columnSize := len(columnCheck)
	if len(rows[0]) < columnSize {
		return nil, fmt.Errorf("read colleges from file header less than %d", columnSize)
	}
	for i, v := range columnCheck {
		if rows[0][i] != v {
			return nil, fmt.Errorf("read colleges from file header[%d] %s does not exist", i, v)
		}
	}

	colleges := make([]*entity.ToAddCollege, length-1)
	columnMap := map[int]string{0: "Name"}
	fieldMap := make(map[string]int)
	v := reflect.ValueOf(&entity.ToAddCollege{}).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldMap[t.Field(i).Name] = i
	}

	for i := 1; i < length; i++ {
		var college entity.ToAddCollege
		collegeValue := reflect.ValueOf(&college).Elem()

		if len(rows[i]) < columnSize {
			return nil, fmt.Errorf("read colleges from file less than %d column", columnSize)
		}

		for j := 0; j < 1; j++ {
			col := strings.Trim(rows[i][j], "")
			if col == "" {
				return nil, fmt.Errorf("read colleges from file cell[%d][%d] %s empty", i, j, columnMap[j])
			}
			collegeValue.Field(fieldMap[columnMap[j]]).Set(reflect.ValueOf(col))
		}

		colleges[i-1] = &college
	}

	return colleges, nil
}

func (serv *ManagementService) CreateCollege(toCreates []*entity.ToAddCollege) error {
	ctx := context.Background()

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("create college start a transaction failed: %w", err)
	}

	current := 0
	length := len(toCreates)
	for ; current < length; current++ {
		id, err := tx.College.Query().Where(func(s *sql.Selector) {
			s.Where(sql.NotNull(college.FieldDeletedTime))
		}).FirstID(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return fmt.Errorf("create college find a deleted college id query failed: %w", err)
			}
			break
		}

		num, err := tx.College.Update().Where(college.And(
			college.IDEQ(id),
			func(s *sql.Selector) {
				s.Where(sql.NotNull(college.FieldDeletedTime))
			},
		)).
			SetCreatedTime(time.Now()).
			ClearDeletedTime().
			ClearModifiedTime().
			SetIsDisabled(false).
			SetName(toCreates[current].Name).
			Save(ctx)
		if err != nil {
			return rollback(tx, "create college", err)
		}
		if num == 0 {
			return rollback(tx, "create college", fmt.Errorf("create college update deleted student affect 0 row"))
		}
	}
	if current < length {
		num, err := tx.College.Query().Aggregate(ent.Count()).Int(ctx)
		if err != nil {
			return fmt.Errorf("create college count failed: %w", err)
		}

		bulkLength := length - current
		bulk := make([]*ent.CollegeCreate, bulkLength)
		for i := 0; i < bulkLength; i++ {
			bulk[i] = tx.College.Create().SetID(uint8(num + i + 1)).
				SetName(toCreates[current+i].Name)
		}

		err = tx.College.CreateBulk(bulk...).Exec(ctx)
		if err != nil {
			return rollback(tx, "create college", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("create college transaction commit failed: %w", err)
	}

	return nil
}

func (serv *ManagementService) RetrieveCollege(current, pageSize *int, like, sort string, order, isDisabled *bool) ([]*ent.College, error) {
	ctx := context.Background()

	query := serv.DB.College.Query().
		Where(college.And(
			college.Or(
				college.NameContains(like),
			),
			func(s *sql.Selector) {
				s.Where(sql.IsNull(college.FieldDeletedTime))
				if isDisabled != nil {
					s.Where(sql.EQ(college.FieldIsDisabled, *isDisabled))
				}
			},
		)).
		Order(func(s *sql.Selector) {
			isSorted := sort != "" && (sort == college.FieldID || sort == college.FieldName || sort == college.FieldIsDisabled)
			if isSorted && order != nil {
				if *order {
					s.OrderBy(sql.Desc(sort))
				} else {
					s.OrderBy(sql.Asc(sort))
				}
			}
		})

	if pageSize != nil && current != nil {
		offset := (*current - 1) * (*pageSize)
		query.Limit(*pageSize).Offset(offset)
	}
	colleges, err := query.All(ctx)

	if err != nil {
		return nil, fmt.Errorf("retrieve college query failed: %w", err)
	}

	return colleges, nil
}

func (serv *ManagementService) GetTotalRetrievedColleges(like string, isDisabled *bool) (int, error) {
	ctx := context.Background()

	total, err := serv.DB.College.Query().
		Where(college.And(
			college.Or(
				college.NameContains(like),
			),
			func(s *sql.Selector) {
				s.Where(sql.IsNull(college.FieldDeletedTime))
				if isDisabled != nil {
					s.Where(sql.EQ(college.FieldIsDisabled, *isDisabled))
				}
			},
		)).
		Count(ctx)

	if err != nil {
		return -1, fmt.Errorf("get retrieved total colleges query failed: %w", err)
	}

	return total, nil
}

func (serv *ManagementService) ReadMajorsFromFile(f *os.File) ([]*entity.ToAddMajor, error) {
	book, err := excelize.OpenFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("read majors from file open failed: %w", err)
	}

	rows, err := book.GetRows(book.GetSheetName(0))
	if err != nil {
		return nil, fmt.Errorf("read majors from file get rows failed: %w", err)
	}

	length := len(rows)
	if length < 2 {
		return nil, fmt.Errorf("read majors from file less than two rows failed: %w", err)
	}

	columnCheck := []string{"专业名称"}
	columnSize := len(columnCheck)
	if len(rows[0]) < columnSize {
		return nil, fmt.Errorf("read majors from file header less than %d", columnSize)
	}
	for i, v := range columnCheck {
		if rows[0][i] != v {
			return nil, fmt.Errorf("read majors from file header[%d] %s does not exist", i, v)
		}
	}

	majors := make([]*entity.ToAddMajor, length-1)
	columnMap := map[int]string{0: "Name"}
	fieldMap := make(map[string]int)
	v := reflect.ValueOf(&entity.ToAddMajor{}).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldMap[t.Field(i).Name] = i
	}

	for i := 1; i < length; i++ {
		var major entity.ToAddMajor
		majorValue := reflect.ValueOf(&major).Elem()

		if len(rows[i]) < columnSize {
			return nil, fmt.Errorf("read majors from file less than %d column", columnSize)
		}

		for j := 0; j < 1; j++ {
			col := strings.Trim(rows[i][j], "")
			if col == "" {
				return nil, fmt.Errorf("read majors from file cell[%d][%d] %s empty", i, j, columnMap[j])
			}
			majorValue.Field(fieldMap[columnMap[j]]).Set(reflect.ValueOf(col))
		}

		majors[i-1] = &major
	}

	return majors, nil
}

func (serv *ManagementService) CreateMajorByCollegeID(toCreates []*entity.ToAddMajor, collegeID uint8) error {
	ctx := context.Background()

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("create major start a transaction failed: %w", err)
	}

	current := 0
	length := len(toCreates)
	for ; current < length; current++ {
		id, err := tx.Major.Query().Where(func(s *sql.Selector) {
			s.Where(sql.NotNull(major.FieldDeletedTime))
		}).FirstID(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return fmt.Errorf("create major find a deleted major id query failed: %w", err)
			}
			break
		}

		num, err := tx.Major.Update().Where(major.And(
			major.IDEQ(id),
			func(s *sql.Selector) {
				s.Where(sql.NotNull(major.FieldDeletedTime))
			},
		)).
			SetCreatedTime(time.Now()).
			ClearDeletedTime().
			ClearModifiedTime().
			SetIsDisabled(false).
			SetName(toCreates[current].Name).
			SetCid(collegeID).
			Save(ctx)
		if err != nil {
			return rollback(tx, "create major", err)
		}
		if num == 0 {
			return rollback(tx, "create major", fmt.Errorf("create major update deleted major affect 0 row"))
		}
	}
	if current < length {
		num, err := tx.Major.Query().Aggregate(ent.Count()).Int(ctx)
		if err != nil {
			return fmt.Errorf("create major count failed: %w", err)
		}

		bulkLength := length - current
		bulk := make([]*ent.MajorCreate, bulkLength)
		for i := 0; i < bulkLength; i++ {
			bulk[i] = tx.Major.Create().SetID(uint16(num + i + 1)).
				SetName(toCreates[current+i].Name).
				SetCid(collegeID)
		}

		err = tx.Major.CreateBulk(bulk...).Exec(ctx)
		if err != nil {
			return rollback(tx, "create major", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("create major transaction commit failed: %w", err)
	}

	return nil
}

func (serv *ManagementService) RetrieveMajorWithCollege(current, pageSize *int, like, sort string, order, isDisabled *bool) ([]*ent.Major, error) {
	ctx := context.Background()

	query := serv.DB.Major.Query().
		Where(major.And(
			major.Or(
				major.NameContains(like),
			),
			func(s *sql.Selector) {
				s.Where(sql.IsNull(major.FieldDeletedTime))
				if isDisabled != nil {
					s.Where(sql.EQ(major.FieldIsDisabled, *isDisabled))
				}
			},
		)).
		Order(func(s *sql.Selector) {
			isSorted := sort != "" && (sort == major.FieldID || sort == major.FieldName || sort == major.FieldIsDisabled)
			if isSorted && order != nil {
				if *order {
					s.OrderBy(sql.Desc(sort))
				} else {
					s.OrderBy(sql.Asc(sort))
				}
			}
		}).
		WithCollege(func(cq *ent.CollegeQuery) {
			cq.Where(func(s *sql.Selector) {
				s.Where(sql.IsNull(college.FieldDeletedTime))
			})
		})

	if pageSize != nil && current != nil {
		offset := (*current - 1) * (*pageSize)
		query.Limit(*pageSize).Offset(offset)
	}
	majors, err := query.All(ctx)

	if err != nil {
		return nil, fmt.Errorf("retrieve major query failed: %w", err)
	}

	return majors, nil
}

func (serv *ManagementService) GetTotalRetrievedMajors(like string, isDisabled *bool) (int, error) {
	ctx := context.Background()

	total, err := serv.DB.Major.Query().
		Where(major.And(
			major.Or(
				major.NameContains(like),
			),
			func(s *sql.Selector) {
				s.Where(sql.IsNull(major.FieldDeletedTime))
				if isDisabled != nil {
					s.Where(sql.EQ(major.FieldIsDisabled, *isDisabled))
				}
			},
		)).Count(ctx)

	if err != nil {
		return -1, fmt.Errorf("get retrieved total majors query failed: %w", err)
	}

	return total, nil
}

func (serv *ManagementService) ReadClassesFromFile(f *os.File) ([]*entity.ToAddClass, error) {
	book, err := excelize.OpenFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("read classes from file open failed: %w", err)
	}

	rows, err := book.GetRows(book.GetSheetName(0))
	if err != nil {
		return nil, fmt.Errorf("read classes from file get rows failed: %w", err)
	}

	length := len(rows)
	if length < 2 {
		return nil, fmt.Errorf("read classes from file less than two rows failed: %w", err)
	}

	columnCheck := []string{"年级", "班级名称"}
	columnSize := len(columnCheck)
	if len(rows[0]) < columnSize {
		return nil, fmt.Errorf("read classes from file header less than %d", columnSize)
	}
	for i, v := range columnCheck {
		if rows[0][i] != v {
			return nil, fmt.Errorf("read classes from file header[%d] %s does not exist", i, v)
		}
	}

	classes := make([]*entity.ToAddClass, length-1)
	columnMap := map[int]string{0: "Grade", 1: "Name"}
	fieldMap := make(map[string]int)
	v := reflect.ValueOf(&entity.ToAddClass{}).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldMap[t.Field(i).Name] = i
	}

	for i := 1; i < length; i++ {
		var class entity.ToAddClass
		classValue := reflect.ValueOf(&class).Elem()

		if len(rows[i]) < columnSize {
			return nil, fmt.Errorf("read classes from file less than %d columns", columnSize)
		}

		for j := 0; j < 2; j++ {
			col := strings.Trim(rows[i][j], "")
			if col == "" {
				return nil, fmt.Errorf("read classes from file cell[%d][%d] %s empty", i, j, columnMap[j])
			}
			classValue.Field(fieldMap[columnMap[j]]).Set(reflect.ValueOf(col))
		}

		classes[i-1] = &class
	}

	return classes, nil
}

func (serv *ManagementService) CreateClassByMajorID(toCreates []*entity.ToAddClass, majorID uint16) error {
	ctx := context.Background()

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("create class start a transaction failed: %w", err)
	}

	current := 0
	length := len(toCreates)
	for ; current < length; current++ {
		id, err := tx.Class.Query().Where(func(s *sql.Selector) {
			s.Where(sql.NotNull(class.FieldDeletedTime))
		}).FirstID(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return fmt.Errorf("create class find a deleted class id query failed: %w", err)
			}
			break
		}

		num, err := tx.Class.Update().Where(class.And(
			class.IDEQ(id),
			func(s *sql.Selector) {
				s.Where(sql.NotNull(class.FieldDeletedTime))
			},
		)).
			SetCreatedTime(time.Now()).
			ClearDeletedTime().
			ClearModifiedTime().
			SetIsDisabled(false).
			SetName(toCreates[current].Name).
			SetGrade(toCreates[current].Grade).
			SetMid(majorID).
			Save(ctx)
		if err != nil {
			return rollback(tx, "create class", err)
		}
		if num == 0 {
			return rollback(tx, "create class", fmt.Errorf("create class update deleted class affect 0 row"))
		}
	}
	if current < length {
		num, err := tx.Class.Query().Aggregate(ent.Count()).Int(ctx)
		if err != nil {
			return fmt.Errorf("create class count failed: %w", err)
		}

		bulkLength := length - current
		bulk := make([]*ent.ClassCreate, bulkLength)
		for i := 0; i < bulkLength; i++ {
			bulk[i] = tx.Class.Create().SetID(uint32(num + i + 1)).
				SetName(toCreates[current+i].Name).
				SetGrade(toCreates[current+i].Grade).
				SetMid(majorID)
		}

		err = tx.Class.CreateBulk(bulk...).Exec(ctx)
		if err != nil {
			return rollback(tx, "create class", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("create class transaction commit failed: %w", err)
	}

	return nil
}

func (serv *ManagementService) RetrieveClassWithMajorAndCollege(current, pageSize *int, like, sort string, order, isDisabled *bool) ([]*ent.Class, error) {
	ctx := context.Background()

	query := serv.DB.Class.Query().
		Where(class.And(
			class.Or(
				class.GradeContains(like),
				class.NameContains(like),
			),
			func(s *sql.Selector) {
				s.Where(sql.IsNull(class.FieldDeletedTime))
				if isDisabled != nil {
					s.Where(sql.EQ(class.FieldIsDisabled, *isDisabled))
				}
			},
		)).
		Order(func(s *sql.Selector) {
			isSorted := sort != "" && (sort == class.FieldID || sort == class.FieldName || sort == class.FieldIsDisabled ||
				sort == class.FieldGrade)
			if isSorted && order != nil {
				if *order {
					s.OrderBy(sql.Desc(sort))
				} else {
					s.OrderBy(sql.Asc(sort))
				}
			}
		}).
		WithMajor(func(mq *ent.MajorQuery) {
			mq.Where(func(s *sql.Selector) {
				s.Where(sql.IsNull(major.FieldDeletedTime))
			}).
				WithCollege(func(cq *ent.CollegeQuery) {
					cq.Where(func(s *sql.Selector) {
						s.Where(sql.IsNull(college.FieldDeletedTime))
					})
				})
		})

	if pageSize != nil && current != nil {
		offset := (*current - 1) * (*pageSize)
		query.Limit(*pageSize).Offset(offset)
	}
	classes, err := query.All(ctx)

	if err != nil {
		return nil, fmt.Errorf("retrieve class query failed: %w", err)
	}

	return classes, nil
}

func (serv *ManagementService) GetTotalRetrievedClasses(like string, isDisabled *bool) (int, error) {
	ctx := context.Background()

	total, err := serv.DB.Class.Query().
		Where(class.And(
			class.Or(
				class.GradeContains(like),
				class.NameContains(like),
			),
			func(s *sql.Selector) {
				s.Where(sql.IsNull(class.FieldDeletedTime))
				if isDisabled != nil {
					s.Where(sql.EQ(class.FieldIsDisabled, *isDisabled))
				}
			},
		)).Count(ctx)

	if err != nil {
		return -1, fmt.Errorf("get retrieved total classes query failed: %w", err)
	}

	return total, nil
}

func (serv *ManagementService) ReadTeachersFromFile(f *os.File) ([]*entity.ToAddTeacher, error) {
	book, err := excelize.OpenFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("read teachers from file open failed: %w", err)
	}

	rows, err := book.GetRows(book.GetSheetName(0))
	if err != nil {
		return nil, fmt.Errorf("read teachers from file get rows failed: %w", err)
	}

	length := len(rows)
	if length < 2 {
		return nil, fmt.Errorf("read teachers from file less than two rows failed: %w", err)
	}

	columnCheck := []string{"工号", "姓名", "性别"}
	columnSize := len(columnCheck)
	if len(rows[0]) < columnSize {
		return nil, fmt.Errorf("read teachers from file header less than %d", columnSize)
	}
	for i, v := range columnCheck {
		if rows[0][i] != v {
			return nil, fmt.Errorf("read teachers from file header[%d] %s does not exist", i, v)
		}
	}

	teachers := make([]*entity.ToAddTeacher, length-1)
	columnMap := map[int]string{0: "TeacherID", 1: "Name", 2: "Gender"}
	fieldMap := make(map[string]int)
	v := reflect.ValueOf(&entity.ToAddTeacher{}).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldMap[t.Field(i).Name] = i
	}

	for i := 1; i < length; i++ {
		var teacher entity.ToAddTeacher
		teacherValue := reflect.ValueOf(&teacher).Elem()

		if len(rows[i]) < columnSize {
			return nil, fmt.Errorf("read teachers from file less than %d columns", columnSize)
		}

		for j := 0; j < columnSize; j++ {
			col := strings.Trim(rows[i][j], "")
			if col == "" {
				return nil, fmt.Errorf("read teachers from file cell[%d][%d] %s empty", i, j, columnMap[j])
			}
			switch j {
			case 2:
				{
					var gender uint8
					switch col {
					case "男":
						gender = 0
					case "女":
						gender = 1
					default:
						return nil, fmt.Errorf("read teachers from file cell[%d][%d] %s invalid", i, j, columnMap[j])
					}
					teacherValue.Field(fieldMap[columnMap[j]]).Set(reflect.ValueOf(gender))

				}
			default:
				teacherValue.Field(fieldMap[columnMap[j]]).Set(reflect.ValueOf(col))
			}
		}

		teachers[i-1] = &teacher
	}

	return teachers, nil
}

func (serv *ManagementService) CreateTeacherByCollegeIDAndCreateUser(toCreates []*entity.ToAddTeacher, collegeID uint8) error {
	ctx := context.Background()

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("create teacher start a transaction failed: %w", err)
	}

	current := 0
	length := len(toCreates)
	for ; current < length; current++ {
		id, err := tx.Teacher.Query().Where(func(s *sql.Selector) {
			s.Where(sql.NotNull(teacher.FieldDeletedTime))
		}).FirstID(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return fmt.Errorf("create teacher find a deleted student id query failed: %w", err)
			}
			break
		}

		num, err := tx.Teacher.Update().Where(teacher.And(
			teacher.IDEQ(id),
			func(s *sql.Selector) {
				s.Where(sql.NotNull(teacher.FieldDeletedTime))
			},
		)).
			SetCreatedTime(time.Now()).
			ClearDeletedTime().
			ClearModifiedTime().
			SetIsDisabled(false).
			SetName(toCreates[current].Name).
			SetTeacherID(toCreates[current].TeacherID).
			SetGender(toCreates[current].Gender).
			SetCid(collegeID).
			Save(ctx)
		if err != nil {
			return rollback(tx, "create teacher", err)
		}
		if num == 0 {
			return rollback(tx, "create teacher", fmt.Errorf("create teacher update deleted teacher affect 0 row"))
		}
	}
	if current < length {
		num, err := tx.Teacher.Query().Aggregate(ent.Count()).Int(ctx)
		if err != nil {
			return fmt.Errorf("create teacher count failed: %w", err)
		}

		bulkLength := length - current
		bulk := make([]*ent.TeacherCreate, bulkLength)
		for i := 0; i < bulkLength; i++ {
			bulk[i] = tx.Teacher.Create().SetID(uint32(num + i + 1)).
				SetName(toCreates[current+i].Name).
				SetTeacherID(toCreates[current+i].TeacherID).
				SetGender(toCreates[current+i].Gender).
				SetCid(collegeID)
		}

		err = tx.Teacher.CreateBulk(bulk...).Exec(ctx)
		if err != nil {
			return rollback(tx, "create teacher", err)
		}
	}

	r, err := tx.Role.Query().Where(role.And(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(role.FieldDeletedTime))
		},
		role.NameEQ("teacher"),
	)).First(ctx)
	if err != nil {
		return rollback(tx, "create user", fmt.Errorf("create user found role by name failed: %w", err))
	}

	teacherIDs := make([]string, length)
	for i, v := range toCreates {
		teacherIDs[i] = v.TeacherID
	}
	t, err := tx.Teacher.Query().Where(teacher.And(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(teacher.FieldDeletedTime))
		},
		teacher.TeacherIDIn(teacherIDs...),
	)).All(ctx)
	if err != nil {
		return rollback(tx, "create user", fmt.Errorf("create user found teacher by teacher ids failed: %w", err))
	}
	if len(t) != length {
		return rollback(tx, "create user", fmt.Errorf("create user found teacher by teacher ids not enough failed"))
	}

	original_pwd := "123456"
	md5_pwd := fmt.Sprintf("%x", md5.Sum([]byte(original_pwd)))
	sha256_pwd := fmt.Sprintf("%x", sha256.Sum256([]byte(md5_pwd)))

	current = 0
	for ; current < length; current++ {
		id, err := tx.User.Query().Where(func(s *sql.Selector) {
			s.Where(sql.NotNull(user.FieldDeletedTime))
		}).FirstID(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return fmt.Errorf("create user find a deleted user id query failed: %w", err)
			}
			break
		}

		num, err := tx.User.Update().Where(user.And(
			user.IDEQ(id),
			func(s *sql.Selector) {
				s.Where(sql.NotNull(user.FieldDeletedTime))
			},
		)).
			SetCreatedTime(time.Now()).
			ClearModifiedTime().
			ClearDeletedTime().
			SetAccount(t[current].TeacherID).
			SetUsername(t[current].Name).
			SetIntroduction("").
			SetIsDisabled(false).
			SetPassword(sha256_pwd).
			AddRoles(r).
			SetTeacherID(t[current].ID).
			Save(ctx)

		if err != nil {
			return rollback(tx, "create user", err)
		}
		if num == 0 {
			return rollback(tx, "create user", fmt.Errorf("create user update deleted user affect 0 row"))
		}
	}
	if current < length {
		num, err := tx.User.Query().Aggregate(ent.Count()).Int(ctx)
		if err != nil {
			return fmt.Errorf("create users count failed: %w", err)
		}

		bulkLength := length - current
		for i := 0; i < bulkLength; i++ {
			err = tx.User.Create().SetID(uint32(num + i + 1)).
				SetAccount(t[current+i].TeacherID).
				SetUsername(t[current+i].Name).
				SetPassword(sha256_pwd).
				SetIntroduction("").
				AddRoles(r).
				SetTeacherID(t[current].ID).
				Exec(ctx)
			if err != nil {
				return rollback(tx, "create users", err)
			}
		}

	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("create teacher transaction commit failed: %w", err)
	}

	return nil
}
