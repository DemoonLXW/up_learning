package facade

import "github.com/google/wire"

var FacadeProvider = wire.NewSet(
	wire.Struct(new(Facades), "*"),
	wire.Struct(new(TeacherFacade), "*"),
)

type Facades struct {
	Teacher *TeacherFacade
}
