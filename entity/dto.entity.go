package entity

import "fmt"

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
	Current   int    `form:"current" binding:"required" `
	PageSize  int    `form:"pagesize" binding:"required"`
	Like      string `form:"like"`
	Sort      string `form:"sort"`
	Order     *bool  `form:"order"`
	IsDeleted *bool  `form:"isdeleted"`
}

type RetrievedRoles struct {
	ID          uint8  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDeleted   bool   `json:"isDeleted"`
}

type RetrievedListData struct {
	Total      int         `json:"total"`
	IsNext     bool        `json:"isNext"`
	IsPrevious bool        `json:"isPrevious"`
	Record     interface{} `json:"record"`
}
