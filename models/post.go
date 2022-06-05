package models

import (
	"Go-Blog/config"
	"html/template"
	"time"
)

type Post struct {
	Pid        int       `json:"pid"` //文章ID
	Title      string    `json:"Title"`
	Slug       string    `json:"slug"`       //自定义文章path
	Content    string    `json:"content"`    //文章的html
	Markdown   string    `json:"markdown"`   //文章的markdown
	CategoryId int       `json:"categoryId"` //分类id
	UserId     int       `json:"userId"`
	ViewCount  int       `json:"viewCount"`
	Type       int       `json:"type"` //文章类型 0 普通，1 自定义文章
	CreateAt   time.Time `json:"creatAt"`
	UpdateAt   time.Time `json:"updateAt"`
}

type PostMore struct {
	Pid        int           `json:"pid"` //文章ID
	Title      string        `json:"Title"`
	Slug       string        `json:"slug"`       //自定义文章path
	Content    template.HTML `json:"content"`    //文章的html
	Markdown   string        `json:"markdown"`   //文章的markdown
	CategoryId int           `json:"categoryId"` //分类id
	UserId     int           `json:"userId"`
	UserName   string        `json:"userName"`
	ViewCount  int           `json:"viewCount"`
	Type       int           `json:"type"` //文章类型 0 普通，1 自定义文章
	CreateAt   time.Time     `json:"creatAt"`
	UpdateAt   time.Time     `json:"updateAt"`
}

type PostReq struct {
	Pid        int    `json:"pid"` //文章ID
	Title      string `json:"Title"`
	Slug       string `json:"slug"`       //自定义文章path
	Content    string `json:"content"`    //文章的html
	Markdown   string `json:"markdown"`   //文章的markdown
	CategoryId int    `json:"categoryId"` //分类id
	UserId     int    `json:"userId"`
	Type       int    `json:"type"` //文章类型 0 普通，1 自定义文章
}

type SearchRes struct {
	Pid   int    `orm:"pid" json:"pid"`
	Title string `orm:"title" json:"title"`
}

type PostRes struct {
	config.Viewer
	config.SystemConfig
	Article PostMore
}
