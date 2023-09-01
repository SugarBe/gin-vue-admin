package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/todo_list"
)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	ExampleApiGroup   example.ApiGroup
	Todo_listApiGroup todo_list.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
