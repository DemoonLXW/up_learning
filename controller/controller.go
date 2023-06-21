package controller

import "github.com/google/wire"

var ControllerProvider = wire.NewSet(
	wire.Struct(new(Controller), "*"),
	wire.Struct(new(UserController), "*"),
)

type Controller struct {
	User *UserController
}
