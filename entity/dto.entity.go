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
	Current    int    `form:"current" binding:"required,gte=1"`
	PageSize   int    `form:"pagesize" binding:"required,gte=1"`
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
	IDs []uint8 `json:"ids" binding:"required"`
}

type RetrievedPermission struct {
	ID          uint16 `json:"id" uri:"id" binding:"required"`
	Action      string `json:"action"`
	Description string `json:"description"`
	IsDisabled  bool   `json:"isDisabled"`
}

type ToAddPermission struct {
	Action      *string `json:"action" binding:"required"`
	Description *string `json:"description" binding:"required"`
}
