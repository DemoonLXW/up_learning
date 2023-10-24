package entity

import (
	"fmt"
	"time"
)

type Result struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Login struct {
	Account  *string `json:"account" binding:"required"`
	Password *string `json:"password" binding:"required"`
}

type Menu struct {
	Title    string  `json:"title"`
	URL      string  `json:"url"`
	Icon     string  `json:"icon"`
	Children []*Menu `json:"children"`
}

func (m *Menu) String() string {
	return fmt.Sprintf("{title=%s, url=%s, icon=%s, children=%v}", m.Title, m.URL, m.Icon, m.Children)
}

type Search struct {
	Current    *int   `form:"current"`
	PageSize   *int   `form:"pagesize"`
	Like       string `form:"like"`
	Sort       string `form:"sort"`
	Order      *bool  `form:"order"`
	IsDisabled *bool  `form:"isdisabled"`
}

type RetrievedRole struct {
	ID          uint8  `json:"id" uri:"id" binding:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDisabled  bool   `json:"isDisabled"`
}

type RetrievedDetailedRole struct {
	ID           uint8      `json:"id" uri:"id" binding:"required"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	IsDisabled   bool       `json:"isDisabled"`
	CreatedTime  *time.Time `json:"createdTime"`
	ModifiedTime *time.Time `json:"modifiedTime"`
}

type RetrievedListData struct {
	Total      int         `json:"total"`
	IsNext     bool        `json:"isNext"`
	IsPrevious bool        `json:"isPrevious"`
	Record     interface{} `json:"record"`
}

type ToAddRole struct {
	Name        *string `json:"name" binding:"required"`
	Description *string `json:"description" binding:"required"`
}

type ToModifyRole struct {
	ID          uint8   `json:"id" binding:"required"`
	Name        *string `json:"name" binding:"required_without_all=Description IsDisabled"`
	Description *string `json:"description" binding:"required_without_all=Name IsDisabled"`
	IsDisabled  *bool   `json:"isDisabled" binding:"required_without_all=Name Description"`
}

type ToRemoveRoleIDs struct {
	IDs []uint8 `json:"ids" binding:"required,gte=1"`
}

type RetrievedPermission struct {
	ID          uint16 `json:"id" uri:"id" binding:"required"`
	Action      string `json:"action"`
	Description string `json:"description"`
	IsDisabled  bool   `json:"isDisabled"`
}

type RetrievedDetailedPermission struct {
	ID           uint16     `json:"id" uri:"id" binding:"required"`
	Action       string     `json:"action"`
	Description  string     `json:"description"`
	IsDisabled   bool       `json:"isDisabled"`
	CreatedTime  *time.Time `json:"createdTime"`
	ModifiedTime *time.Time `json:"modifiedTime"`
}

type ToAddPermission struct {
	Action      *string `json:"action" binding:"required"`
	Description *string `json:"description" binding:"required"`
}

type ToModifyPermission struct {
	ID          uint16  `json:"id" binding:"required"`
	Action      *string `json:"action" binding:"required_without_all=Description IsDisabled"`
	Description *string `json:"description" binding:"required_without_all=Action IsDisabled"`
	IsDisabled  *bool   `json:"isDisabled" binding:"required_without_all=Action Description"`
}

type ToRemovePermissionIDs struct {
	IDs []uint16 `json:"ids" binding:"required,gte=1"`
}

type ToModifyPermissionsOfRoles struct {
	RIDs []uint8  `json:"rids" binding:"required,gte=1"`
	PIDs []uint16 `json:"pids" binding:"required"`
	// IsDeleted *bool    `json:"isDeleted" binding:"required"`
}

type RetrievedUser struct {
	ID         uint32 `json:"id" uri:"id" binding:"required"`
	Account    string `json:"account"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	IsDisabled bool   `json:"isDisabled"`
}

type ToAddUser struct {
	Account  *string `json:"account" binding:"required"`
	Username *string `json:"username" binding:"required"`
	Password *string `json:"password"`
}

type ToAddUserParam struct {
	Users  []ToAddUser `json:"users" binding:"required"`
	RoleID uint8       `json:"rid" binding:"required,gte=1"`
}

type ToAddSchool struct {
	Code                string `json:"code"`
	Name                string `json:"name"`
	Location            string `json:"location"`
	CompetentDepartment string `json:"competent_department"`
	EducationLevel      uint8  `json:"education_level"`
	Remark              string `json:"remark"`
}

type RetrievedSchool struct {
	ID                  uint16 `json:"id" uri:"id" binding:"required"`
	Code                string `json:"code"`
	Name                string `json:"name"`
	Location            string `json:"location"`
	CompetentDepartment string `json:"competent_department"`
	EducationLevel      string `json:"education_level"`
	Remark              string `json:"remark"`
	IsDisabled          bool   `json:"isDisabled"`
}

type ToAddStudent struct {
	StudentID string `json:"studentID"`
	Name      string `json:"name"`
	Gender    uint8  `json:"gender"`
}

type RetrievedStudent struct {
	ID        uint32 `json:"id"`
	StudentID string `json:"studentID"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	// Class      *RetrievedClass `json:"class"`
	// User       *RetrievedUser  `json:"user"`
	IsDisabled bool `json:"isDisabled"`
}

type ToAddTeacher struct {
	TeacherID string `json:"teacherID"`
	Name      string `json:"name"`
	Gender    uint8  `json:"gender"`
}

type RetrievedTeacher struct {
	ID         uint32 `json:"id"`
	TeacherID  string `json:"teacherID"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	IsDisabled bool   `json:"isDisabled"`
}

type ToAddCollege struct {
	Name string `json:"name"`
}

type RetrievedCollege struct {
	ID         uint8  `json:"id" uri:"id"`
	Name       string `json:"name"`
	IsDisabled bool   `json:"isDisabled"`
}

type ToAddMajor struct {
	Name string `json:"name"`
}

type RetrievedMajor struct {
	ID   uint16 `json:"id" uri:"id"`
	Name string `json:"name"`
	// College    *RetrievedCollege `json:"college"`
	IsDisabled bool `json:"isDisabled"`
}

type ToAddClass struct {
	Grade string `json:"grade"`
	Name  string `json:"name"`
}

type RetrievedClass struct {
	ID         uint32 `json:"id"`
	Grade      string `json:"grade"`
	Name       string `json:"name"`
	IsDisabled bool   `json:"isDisabled"`
}

type Reviewer struct {
	IDs []uint32 `json:"ids"`
}

type Executor struct {
	IDs []uint32 `json:"ids"`
}

type ToAddFile struct {
	UID  uint32 `json:"uid"`
	Name string `json:"name"`
	Path string `json:"path"`
	Size int64  `json:"size"`
}

type ToAddProject struct {
	UID                 uint32       `json:"uid" form:"uid"`
	Title               string       `json:"title" form:"title" binding:"required"`
	Goal                string       `json:"goal" form:"goal"`
	Principle           string       `json:"principle" form:"principle"`
	ProcessAndMethod    string       `json:"process_and_method" form:"process_and_method"`
	Step                string       `json:"step" form:"step"`
	ResultAndConclusion string       `json:"result_and_conclusion" form:"result_and_conclusion"`
	Requirement         string       `json:"requirement" form:"requirement"`
	Attachments         []*ToAddFile `json:"attachments"`
}

type ToModifyProject struct {
	ID                  uint32       `json:"id" binding:"required"`
	Title               *string      `json:"title"`
	Goal                *string      `json:"goal"`
	Principle           *string      `json:"principle"`
	ProcessAndMethod    *string      `json:"process_and_method"`
	Step                *string      `json:"step"`
	ResultAndConclusion *string      `json:"result_and_conclusion"`
	Requirement         *string      `json:"requirement"`
	AddFile             []*ToAddFile `json:"add_file"`
	DeleteFileIDs       []uint32     `json:"delete_file_ids"`
	ReviewStatus        *uint8       `json:"review_status"`
	IsDisabled          *bool        `json:"isDisabled"`
}

type SearchProject struct {
	Current      *int   `form:"current"`
	PageSize     *int   `form:"pagesize"`
	Like         string `form:"like"`
	Sort         string `form:"sort"`
	Order        *bool  `form:"order"`
	IsDisabled   *bool  `form:"isdisabled"`
	ReviewStatus *uint8 `form:"reviewstatus"`
}
