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
	Current   int    `json:"current" binding:"required" `
	PageSize  int    `json:"pageSize" binding:"required"`
	Like      string `json:"like"`
	Sort      string `json:"sort"`
	Order     *bool  `json:"order"`
	IsDeleted *bool  `json:"isDeleted"`
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
