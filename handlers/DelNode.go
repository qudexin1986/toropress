package handlers

import (
	"toropress/libs"
	"toropress/models"
)

type NodeDeleteHandler struct {
	libs.RootAuthHandler
}

func (self *NodeDeleteHandler) Get() {
	nid, _ := self.GetInt64(":nid")
	models.DelNode(nid)
	self.Ctx.Redirect(302, "/")
}
