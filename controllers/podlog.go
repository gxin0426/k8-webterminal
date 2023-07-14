package controllers

import (
	"github.com/astaxie/beego"
)

type PodLogController struct {
	beego.Controller
}

func (self *PodLogController) Get() {
	self.Data["context"] = self.GetString("context")
	self.Data["namespace"] = self.GetString("namespace")
	self.Data["pod"] = self.GetString("pod")
	self.Data["container"] = self.GetString("container")
	self.TplName = "podlog3.html"
}
