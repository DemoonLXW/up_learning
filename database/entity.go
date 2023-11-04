package database

const (
	RoleAdministrator    = uint8(1)
	RoleTeacher          = uint8(2)
	RoleStudent          = uint8(3)
	RoleProjectApplicant = uint8(4)
	RoleOrdinaryUser     = uint8(5)
)

const (
	ReviewProjectCreated   = uint8(0)
	ReviewProjectStarted   = uint8(1)
	ReviewProjectCompleted = uint8(2)
	ReviewProjectFailed    = uint8(3)
)

const (
	ReviewProjectDetailPending = uint8(0)
)
