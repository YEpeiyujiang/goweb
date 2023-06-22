package service

import (
	"goweb/config"
	"goweb/dao"
	"goweb/models"
)

func GetAllIndexInfo() (*models.HomeResponse, error) {
	//数据库查询
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "叶培宇",
			ViewCount:    123,
			CreateAt:     "2023-6-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}
	return hr, nil
}
