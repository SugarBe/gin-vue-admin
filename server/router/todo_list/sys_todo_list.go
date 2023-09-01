package todo_list

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CustomTodoListDBRouter struct {
}

// InitCustomTodoListDBRouter 初始化 CustomTodoListDB 路由信息
func (s *CustomTodoListDBRouter) InitCustomTodoListDBRouter(Router *gin.RouterGroup) {
	TodoListRouter := Router.Group("TodoList").Use(middleware.OperationRecord())
	TodoListRouterWithoutRecord := Router.Group("TodoList")
	var TodoListApi = v1.ApiGroupApp.Todo_listApiGroup.CustomTodoListDBApi
	{
		TodoListRouter.POST("createCustomTodoListDB", TodoListApi.CreateCustomTodoListDB)   // 新建CustomTodoListDB
		TodoListRouter.DELETE("deleteCustomTodoListDB", TodoListApi.DeleteCustomTodoListDB) // 删除CustomTodoListDB
		TodoListRouter.DELETE("deleteCustomTodoListDBByIds", TodoListApi.DeleteCustomTodoListDBByIds) // 批量删除CustomTodoListDB
		TodoListRouter.PUT("updateCustomTodoListDB", TodoListApi.UpdateCustomTodoListDB)    // 更新CustomTodoListDB
	}
	{
		TodoListRouterWithoutRecord.GET("findCustomTodoListDB", TodoListApi.FindCustomTodoListDB)        // 根据ID获取CustomTodoListDB
		TodoListRouterWithoutRecord.GET("getCustomTodoListDBList", TodoListApi.GetCustomTodoListDBList)  // 获取CustomTodoListDB列表
	}
}
