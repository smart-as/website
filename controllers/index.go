package controllers

type IndexController struct {
	BaseController
}

func (ic *IndexController) Get() {
	ic.Data["Title"] = "SmartAs website"
	ic.Data["Type"] = "index"
}
