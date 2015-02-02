package handlers

import (
	"toropress/libs"
	"toropress/models"
)

type TopicDeleteHandler struct {
	libs.RootAuthHandler
}

func (self *TopicDeleteHandler) Get() {
	tid, _ := self.GetInt64(":tid")
	models.DelTopic(tid)
	self.Ctx.Redirect(302, "/")
}
