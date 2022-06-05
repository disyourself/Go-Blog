package models

import "Go-Blog/config"

type HomeResponse struct {
	config.Viewer
	Category []Category
	Posts    []PostMore
	Total    int
	Page     int
	Pages    []int
	PageEnd  bool
}
