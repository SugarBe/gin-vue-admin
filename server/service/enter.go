package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/todo_list"
)

type ServiceGroup struct {
	SystemServiceGroup    system.ServiceGroup
	ExampleServiceGroup   example.ServiceGroup
	Todo_listServiceGroup todo_list.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
