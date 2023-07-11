package controller

import "github.com/google/wire"

var ControllerProvider = wire.NewSet(
	wire.Struct(new(Controllers), "*"),
	wire.Struct(new(UserController), "*"),
	wire.Struct(new(ManagementController), "*"),
)

type Controllers struct {
	User       *UserController
	Management *ManagementController
}
