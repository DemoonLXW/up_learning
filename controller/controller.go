package controller

import "github.com/google/wire"

var ControllerProvider = wire.NewSet(
	wire.Struct(new(Controllers), "*"),
	wire.Struct(new(UserController), "*"),
)

type Controllers struct {
	User *UserController
}
