package controllers

import (
	"strings"

	"github.com/smart-as/website/models"
)

type VideoController struct {
	BaseController
}

func (vc *VideoController) Get() {
	vc.Data["Type"] = "video"
	vc.Data["Video"] = models.Videos[strings.ToLower(vc.Lang)]
}
