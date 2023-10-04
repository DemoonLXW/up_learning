package service_test

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/injection"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/stretchr/testify/assert"
)

func CreateTestManagementService() (*service.ManagementService, error) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	if err != nil {
		return nil, err
	}
	serv.DB = db
	return serv, nil
}

func TestCreatePermission(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	action1 := "action00"
	description1 := "i description"
	permission1 := entity.ToAddPermission{
		Action:      &action1,
		Description: &description1,
	}

	action2 := "action11"
	description2 := "12 description"
	permission2 := entity.ToAddPermission{
		Action:      &action2,
		Description: &description2,
	}
	// action3 := "action13"
	// description3 := "action13 description"
	// permission3 := entity.ToAddPermission{
	// 	Action:      &action3,
	// 	Description: &description3,
	// }
	// action4 := "action14"
	// description4 := "action14 description"
	// permission4 := entity.ToAddPermission{
	// 	Action:      &action4,
	// 	Description: &description4,
	// }
	// err = serv.CreatePermission([]*entity.ToAddPermission{&permission1, &permission2, &permission3, &permission4})
	err = serv.CreatePermission([]*entity.ToAddPermission{&permission1, &permission2})
	assert.Nil(t, err)

}

func TestUpdatePermission(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	// action := "test update permssion in service hhhhhh"
	description := ""
	// disabled := true
	permission := entity.ToModifyPermission{
		ID: 8,
		// Action:      &action,
		Description: &description,
		// IsDisabled:  &disabled,
	}
	err = serv.UpdatePermission(&permission)
	assert.Nil(t, err)

}

func TestRetrievePermission(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	// current := 2
	// pageSize := 3
	order := true
	// disabled := false
	permissons, err := serv.RetrievePermission(nil, nil, "", "id", &order, nil)
	assert.Nil(t, err)
	// for _, v := range permissons {

	// 	fmt.Println(v)
	// }
	assert.Len(t, permissons, 14)
}

func TestDeletePermission(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	IDs := []uint16{15}
	err = serv.DeletePermission(IDs)
	assert.Nil(t, err)

}

func TestCreateRole(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	name1 := "new role2"
	description1 := "new role1 description"
	role1 := entity.ToAddRole{
		Name:        &name1,
		Description: &description1,
	}

	name2 := "new role1"
	description2 := "new role2 description"
	role2 := entity.ToAddRole{
		Name:        &name2,
		Description: &description2,
	}

	name3 := "new role3"
	description3 := "new role3 description"
	role3 := entity.ToAddRole{
		Name:        &name3,
		Description: &description3,
	}

	name4 := "new role4"
	description4 := "new role4 description"
	role4 := entity.ToAddRole{
		Name:        &name4,
		Description: &description4,
	}
	err = serv.CreateRole([]*entity.ToAddRole{&role1, &role2, &role3, &role4})
	assert.Nil(t, err)

}

func TestUpdateRole(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	name := "test change role name"
	description := "test change role service description go go"
	isDisabled := true
	role := entity.ToModifyRole{
		ID:          5,
		Name:        &name,
		Description: &description,
		IsDisabled:  &isDisabled,
	}
	err = serv.UpdateRole(&role)
	assert.Nil(t, err)

}

func TestRetrieveRole(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	dao := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	dao.DB = db

	current := 1
	pageSize := 5
	order := true
	// disabled := false
	roles, err := dao.RetrieveRole(&current, &pageSize, "", "id", &order, nil)
	assert.Nil(t, err)
	for _, v := range roles {
		fmt.Println(v)
	}
	assert.Len(t, roles, 5)
}

func TestGetTotalRetrievedRoles(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	// delete := false
	total, err := serv.GetTotalRetrievedRoles("role", nil)
	assert.Nil(t, err)
	assert.Equal(t, 6, total)
}

func TestDeleteRole(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	IDs := []uint8{10, 12}
	err = serv.DeleteRole(IDs)
	assert.Nil(t, err)

}

func TestFindOneRoleById(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	role, err := serv.FindOneRoleById(5)
	assert.Nil(t, err)
	fmt.Println(role)

}

func TestGetTotalRetrievedPermissions(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	// delete := false
	total, err := serv.GetTotalRetrievedPermissions("", nil)
	assert.Nil(t, err)
	assert.Equal(t, 14, total)
}

func TestFindOnePermissionById(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	p, err := serv.FindOnePermissionById(14)
	assert.Nil(t, err)
	fmt.Println(p)

}

func TestFindPermissionsByRoleIds(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	ps, err := serv.FindPermissionsByRoleIds([]uint8{2, 99})
	assert.Nil(t, err)
	for _, v := range ps {
		fmt.Println(v)
	}
}

func TestUpdatePermissionForRole(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	rids := []uint8{2}
	pids := []uint16{99}

	err = serv.UpdatePermissionForRole(rids, pids)
	assert.Nil(t, err)
}

func TestRetrieveUser(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	// current := 3
	// pageSize := 2
	order := false
	disabled := false
	users, err := serv.RetrieveUser(nil, nil, "", "id", &order, &disabled)
	assert.Nil(t, err)
	for _, v := range users {
		fmt.Println(v)
	}
	assert.Equal(t, len(users), 2)
}

func TestGetTotalRetrievedUsers(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	disabled := false
	total, err := serv.GetTotalRetrievedUsers("", &disabled)
	assert.Nil(t, err)
	assert.Equal(t, 2, total)
}

func TestCreateUser(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	password := "88888888"

	account1 := "test account1"
	username1 := "test username1"
	user1 := entity.ToAddUser{
		Account:  &account1,
		Username: &username1,
		Password: &password,
	}

	account2 := "test account2"
	username2 := "test username2"
	user2 := entity.ToAddUser{
		Account:  &account2,
		Username: &username2,
		Password: &password,
	}

	// account3 := "new account3"
	// username3 := "new username3"
	// user3 := entity.ToAddUser{
	// 	Account:  &account3,
	// 	Username: &username3,
	// 	Password: &password,
	// }

	// account4 := "new account6"
	// username4 := "new username6"
	// user4 := entity.ToAddUser{
	// 	Account:  &account4,
	// 	Username: &username4,
	// 	Password: &password,
	// }

	err = serv.CreateUser([]*entity.ToAddUser{&user1, &user2}, 3)
	assert.Nil(t, err)

}

func TestSaveImportedFile(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	f, err := os.Open("../domain.config.json")
	assert.Nil(t, err)
	defer f.Close()

	b := bytes.Buffer{}
	writer := multipart.NewWriter(&b)
	part, err := writer.CreateFormFile("upload", filepath.Base(f.Name()))
	assert.Nil(t, err)

	_, err = io.Copy(part, f)
	assert.Nil(t, err)
	err = writer.Close()
	assert.Nil(t, err)

	req, err := http.NewRequest(http.MethodPost, "localhost:8080", &b)
	assert.Nil(t, err)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	_, fh, err := req.FormFile("upload")
	assert.Nil(t, err)

	dir := "../temp/importSchool"
	prefix := "liu"

	o, err := serv.SaveImportedFile(fh, dir, prefix)
	assert.Nil(t, err)
	fmt.Println(o.Name())
	// o.Close()
	// os.Remove(o.Name())
}

func TestReadSchoolsFromFile(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	f, err := os.Open("../importSchools.xlsx")
	assert.Nil(t, err)

	schools, err := serv.ReadSchoolsFromFile(f)
	assert.Nil(t, err)

	for _, s := range schools {
		fmt.Println(s.Name, s.Code, s.CompetentDepartment, s.EducationLevel, s.Location, s.Remark)
	}
}

func TestCreateSchool(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	school1 := entity.ToAddSchool{
		Code:                "4161011397",
		Name:                "安康学院",
		CompetentDepartment: "陕西省",
		Location:            "安康市",
		EducationLevel:      0,
		Remark:              "",
	}

	school2 := entity.ToAddSchool{
		Code:                "4161011400",
		Name:                "西安培华学院",
		CompetentDepartment: "陕西省教育厅",
		Location:            "西安市",
		EducationLevel:      0,
		Remark:              "民办",
	}

	// school3 := entity.ToAddSchool{
	// 	Code:                "4161011560",
	// 	Name:                "西安财经大学",
	// 	CompetentDepartment: "陕西省",
	// 	Location:            "西安市",
	// 	EducationLevel:      0,
	// 	Remark:              "",
	// }

	school4 := entity.ToAddSchool{
		Code:                "4161013739",
		Name:                "西安健康工程职业学院",
		CompetentDepartment: "陕西省教育厅",
		Location:            "西安市",
		EducationLevel:      1,
		Remark:              "民办",
	}

	school5 := entity.ToAddSchool{
		Code:                "44161013945",
		Name:                "西安铁路职业技术学院",
		CompetentDepartment: "陕西省",
		Location:            "西安市",
		EducationLevel:      1,
		Remark:              "",
	}

	schools := []*entity.ToAddSchool{&school2, &school1, &school5, &school4}

	err = serv.CreateSchool(schools)
	assert.Nil(t, err)

}

func TestRetrieveSchool(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	// current := 1
	// pageSize := 6
	order := false
	// disabled := false
	schools, err := serv.RetrieveSchool(nil, nil, "", "id", &order, nil)
	assert.Nil(t, err)
	// for _, v := range schools {
	// 	fmt.Println(v)
	// }
	assert.Len(t, schools, 13)
}

func TestGetTotalRetrievedSchools(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	// disabled := true
	total, err := serv.GetTotalRetrievedSchools("相思湖", nil)
	assert.Nil(t, err)
	assert.Equal(t, 0, total)
}

func TestFindOneSampleFileByType(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	f, err := serv.FindOneSampleFileByType("school")
	assert.Nil(t, err)
	fmt.Println(f)
}

func TestReadStudentsFromFile(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	f, err := os.Open("../importStudents2.xlsx")
	assert.Nil(t, err)

	students, err := serv.ReadStudentsFromFile(f)
	assert.Nil(t, err)

	for _, s := range students {
		fmt.Println(s.Name, s.StudentID, s.Gender)
	}
}

func TestCreateStudent(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	// student1 := entity.ToAddStudent{
	// 	StudentID: "20221050045",
	// 	Name:      "lxw",
	// 	Gender:    0,
	// }

	student2 := entity.ToAddStudent{
		StudentID: "20221050046",
		Name:      "lyj",
		Gender:    1,
	}

	student3 := entity.ToAddStudent{
		StudentID: "20221050047",
		Name:      "lyk",
		Gender:    0,
	}

	students := []*entity.ToAddStudent{&student3, &student2}

	err = serv.CreateStudent(students, 3)
	assert.Nil(t, err)

}

func TestRetrieveStudent(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	current := 2
	pageSize := 1
	order := false
	disabled := true
	students, err := serv.RetrieveStudentBySchoolID(&current, &pageSize, "ly", "id", &order, &disabled, 3)
	assert.Nil(t, err)
	for _, v := range students {
		fmt.Println(v)
	}
	assert.Len(t, students, 0)
}

func TestGetTotalRetrievedStudentsBySchoolID(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	// disabled := true
	total, err := serv.GetTotalRetrievedStudentsBySchoolID("ly", nil, 3)
	assert.Nil(t, err)
	assert.Equal(t, 2, total)
}

func TestReadCollegesFromFile(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	f, err := os.Open("../importColleges.xlsx")
	assert.Nil(t, err)

	colleges, err := serv.ReadCollegesFromFile(f)
	assert.Nil(t, err)

	for _, c := range colleges {
		fmt.Println(c.Name)
	}
}

func TestCreateCollege(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	// student1 := entity.ToAddStudent{
	// 	StudentID: "20221050045",
	// 	Name:      "lxw",
	// 	Gender:    0,
	// }

	college1 := entity.ToAddCollege{
		Name: "college3",
	}

	college2 := entity.ToAddCollege{
		Name: "college2",
	}

	colleges := []*entity.ToAddCollege{&college1, &college2}

	err = serv.CreateCollege(colleges)
	assert.Nil(t, err)

}

func TestRetrieveCollege(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	// current := 1
	pageSize := 3
	order := true
	// disabled := true
	colleges, err := serv.RetrieveCollege(nil, &pageSize, "", "id", &order, nil)
	assert.Nil(t, err)
	for _, v := range colleges {
		fmt.Println(v)
	}
	assert.Len(t, colleges, 3)
}

func TestGetTotalRetrievedColleges(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	disabled := true
	total, err := serv.GetTotalRetrievedColleges("", &disabled)
	assert.Nil(t, err)
	assert.Equal(t, 1, total)
}

func TestReadMajorsFromFile(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	f, err := os.Open("../import.xlsx")
	assert.Nil(t, err)

	majors, err := serv.ReadMajorsFromFile(f)
	assert.Nil(t, err)

	for _, m := range majors {
		fmt.Println(m.Name)
	}
}

func TestCCreateMajorByCollegeID(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	major1 := entity.ToAddMajor{
		Name: "major1",
	}

	major2 := entity.ToAddMajor{
		Name: "major2",
	}

	major3 := entity.ToAddMajor{
		Name: "major3",
	}

	majors := []*entity.ToAddMajor{&major2, &major1, &major3}

	err = serv.CreateMajorByCollegeID(majors, 3)
	assert.Nil(t, err)

}

func TestRetrieveMajorWithCollege(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	current := 1
	pageSize := 3
	order := false
	// disabled := true
	majors, err := serv.RetrieveMajorWithCollege(&current, &pageSize, "2", "id", &order, nil)
	assert.Nil(t, err)
	for _, v := range majors {
		fmt.Println(v, v.Edges.College.Name)
	}
	assert.Len(t, majors, 2)
}

func TestGetTotalRetrievedMajors(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	disabled := false
	total, err := serv.GetTotalRetrievedMajors("3", &disabled)
	assert.Nil(t, err)
	assert.Equal(t, 1, total)
}

func TestReadClassesFromFile(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	f, err := os.Open("../import.xlsx")
	assert.Nil(t, err)

	classes, err := serv.ReadClassesFromFile(f)
	assert.Nil(t, err)

	for _, c := range classes {
		fmt.Println(c.Name)
	}
}
