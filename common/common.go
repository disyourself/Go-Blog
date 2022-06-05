package common

import (
	"Go-Blog/config"
	"Go-Blog/models"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	Template = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
}
