package models

// 从数据库中获取的
type Category struct {
	Cid      int
	Name     string
	CreateAt string
	UpdateAt string
}

type GetegoryResponse struct {
	*HomeResponse
	CategoryName string
}
