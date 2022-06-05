package views

import (
	"Go-Blog/common"
	"Go-Blog/config"
	"Go-Blog/models"
	"net/http"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {

	index := common.Template.Index

	var categorys = []models.Category{
		{
			Cid:  1,
			Nmae: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "Go Blog",
			Content:      "content",
			CategoryName: "go",
			CategoryId:   1,
			ViewCount:    123,
			Type:         0,
			CreateAt:     "2022-04-01",
		},
	}
	var hr = &models.HomeResponse{
		Viewer:   config.Cfg.Viewer,
		Category: categorys,
		Posts:    posts,
		Total:    1,
		Page:     1,
		Pages:    []int{1},
		PageEnd:  true,
	}

	index.WriteData(w, hr)
}
