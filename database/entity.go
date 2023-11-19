package database

const (
	RoleAdministrator    = uint8(1)
	RoleTeacher          = uint8(2)
	RoleStudent          = uint8(3)
	RoleProjectApplicant = uint8(4)
	RoleOrdinaryUser     = uint8(5)
)

const (
	ProjectReviewStatusUnaudited = uint8(0)
	ProjectReviewStatusPending   = uint8(1)
	ProjectReviewStatusSucceeded = uint8(2)
	ProjectReviewStatusFailed    = uint8(3)
	ProjectReviewStatusTimeout   = uint8(4)
)

// const (
// 	ReviewProjectDetailPending = uint8(0)
// )
