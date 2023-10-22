package facade

import (
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/google/wire"
)

var FacadeProvider = wire.NewSet(
	wire.Struct(new(Facades), "*"),
	wire.Struct(new(TeacherFacade), "*"),
)

type Facades struct {
	DB      *ent.Client
	Teacher *TeacherFacade
}
