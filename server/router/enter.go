package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/todo_list"
)

type RouterGroup struct {
	System    system.RouterGroup
	Example   example.RouterGroup
	Todo_list todo_list.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
