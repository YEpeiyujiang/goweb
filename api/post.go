package api

import (
	"errors"
	"fmt"
	"goweb/common"
	"goweb/dao"
	"goweb/models"
	"goweb/service"
	"goweb/utils"
	"net/http"
	"strconv"
	"time"
)

func SavePost(post *models.Post) {
	dao.SavePost(post)

}

func (*Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("发布文章")
	//获取用户Id 判断用户是否登录
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		fmt.Println("登录已经过期")
		common.Error(w, errors.New("登录已经过期"))
		return
	}
	fmt.Println("1")
	//POST save操作
	uid := claim.Uid
	method := r.Method
	switch method {
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)
		post := &models.Post{
			Pid:        -1,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     uid,
			Type:       pType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
	}
	common.GetRequestJsonParam(r)

}
