package entity

const (
	GetRoleList               = "get role list"
	AddRole                   = "add role"
	ModifyARole               = "modify a role"
	GetARoleById              = "get a role by id"
	RemoveRolesByIds          = "remove roles by ids"
	GetPermissionsByRoleId    = "get permission by role id"
	ModifyPermissionsForRoles = "modify permissions for roles"

	GetPermissionList = "get permission list"
	// AddPermission          = "add permission"
	ModifyAPermission  = "modify a permission"
	GetAPermissionById = "get a permission by id"
	// RemovePermissionsByIds = "remove permissions by ids"

	GetUserList = "get user list"

	ImportSchool  = "import school"
	GetSchoolList = "get school list"

	GetSampleOfImport = "get sample of import"

	GetTeacherList           = "get teacher list"
	ImportTeacherByCollegeID = "import teacher by college id"

	GetStudentList         = "get student list"
	ImportStudentByClassID = "import student by class id"

	ImportCollege        = "import college"
	GetCollegeList       = "get college list"
	GetColleges          = "get college"
	GetMajorsByCollegeID = "get major by college id"

	ImportMajorByCollegeID = "import major by college id"
	GetMajorList           = "get major list"
	GetClassesByMajorID    = "get class by major id"

	ImportClassByMajorID = "import class by major id"
	GetClassList         = "get class list"

	CommomDownloadFile = "common download file"

	ApplicantAddProject                        = "applicant add project"
	ApplicantRemoveProjectsByIds               = "applicant remove projects by ids"
	ApplicantModifyAProject                    = "applicant modify a project"
	ApplicantGetAProjectById                   = "applicant get a project by id"
	ApplicantGetProjectListByUserID            = "applicant get project list by user id"
	ApplicantUploadDocumentImage               = "applicant upload document image"
	ApplicantSubmitProjectForReview            = "applicant submit project for review"
	ApplicantGetReviewProjectRecordByProjectID = "applicant get review project record list by project id"
	ApplicantGetAReviewProjectRecordDetailByID = "applicant get a review project record detail by id"

	ProjectReviewerGetTaskListOfPlatformReviewer   = "project reviewer get task list of platform reviewer"
	ProjectReviewerGetATaskDetailByID              = "project reviewer get a task detail by id"
	ProjectReviewerReviewProjectByTaskID           = "project reviewer review project by task id"
	ProjectReviewerGetReviewedTaskListByReviewerID = "project reviewer get reviewed task list by reviewer id"
)
