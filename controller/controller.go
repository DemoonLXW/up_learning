package controller

import "github.com/google/wire"

var ControllerProvider = wire.NewSet(
	wire.Struct(new(Controllers), "*"),
	wire.Struct(new(AuthController), "*"),
	wire.Struct(new(ManagementController), "*"),
	wire.Struct(new(CommonController), "*"),
	wire.Struct(new(TeacherController), "*"),
)

type Controllers struct {
	Auth       *AuthController
	Management *ManagementController
	Common     *CommonController
	Teacher    *TeacherController
}
