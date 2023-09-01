package todo_list

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/todo_list"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    todo_listReq "github.com/flipped-aurora/gin-vue-admin/server/model/todo_list/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type CustomTodoListDBApi struct {
}

var TodoListService = service.ServiceGroupApp.Todo_listServiceGroup.CustomTodoListDBService


// CreateCustomTodoListDB 创建CustomTodoListDB
// @Tags CustomTodoListDB
// @Summary 创建CustomTodoListDB
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body todo_list.CustomTodoListDB true "创建CustomTodoListDB"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TodoList/createCustomTodoListDB [post]
func (TodoListApi *CustomTodoListDBApi) CreateCustomTodoListDB(c *gin.Context) {
	var TodoList todo_list.CustomTodoListDB
	err := c.ShouldBindJSON(&TodoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    TodoList.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "Status":{utils.NotEmpty()},
        "Expire_time":{utils.NotEmpty()},
    }
	if err := utils.Verify(TodoList, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := TodoListService.CreateCustomTodoListDB(&TodoList); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCustomTodoListDB 删除CustomTodoListDB
// @Tags CustomTodoListDB
// @Summary 删除CustomTodoListDB
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body todo_list.CustomTodoListDB true "删除CustomTodoListDB"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TodoList/deleteCustomTodoListDB [delete]
func (TodoListApi *CustomTodoListDBApi) DeleteCustomTodoListDB(c *gin.Context) {
	var TodoList todo_list.CustomTodoListDB
	err := c.ShouldBindJSON(&TodoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    TodoList.DeletedBy = utils.GetUserID(c)
	if err := TodoListService.DeleteCustomTodoListDB(TodoList); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCustomTodoListDBByIds 批量删除CustomTodoListDB
// @Tags CustomTodoListDB
// @Summary 批量删除CustomTodoListDB
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CustomTodoListDB"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /TodoList/deleteCustomTodoListDBByIds [delete]
func (TodoListApi *CustomTodoListDBApi) DeleteCustomTodoListDBByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := TodoListService.DeleteCustomTodoListDBByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCustomTodoListDB 更新CustomTodoListDB
// @Tags CustomTodoListDB
// @Summary 更新CustomTodoListDB
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body todo_list.CustomTodoListDB true "更新CustomTodoListDB"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TodoList/updateCustomTodoListDB [put]
func (TodoListApi *CustomTodoListDBApi) UpdateCustomTodoListDB(c *gin.Context) {
	var TodoList todo_list.CustomTodoListDB
	err := c.ShouldBindJSON(&TodoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    TodoList.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "Status":{utils.NotEmpty()},
          "Expire_time":{utils.NotEmpty()},
      }
    if err := utils.Verify(TodoList, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := TodoListService.UpdateCustomTodoListDB(TodoList); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCustomTodoListDB 用id查询CustomTodoListDB
// @Tags CustomTodoListDB
// @Summary 用id查询CustomTodoListDB
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query todo_list.CustomTodoListDB true "用id查询CustomTodoListDB"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TodoList/findCustomTodoListDB [get]
func (TodoListApi *CustomTodoListDBApi) FindCustomTodoListDB(c *gin.Context) {
	var TodoList todo_list.CustomTodoListDB
	err := c.ShouldBindQuery(&TodoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reTodoList, err := TodoListService.GetCustomTodoListDB(TodoList.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reTodoList": reTodoList}, c)
	}
}

// GetCustomTodoListDBList 分页获取CustomTodoListDB列表
// @Tags CustomTodoListDB
// @Summary 分页获取CustomTodoListDB列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query todo_listReq.CustomTodoListDBSearch true "分页获取CustomTodoListDB列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TodoList/getCustomTodoListDBList [get]
func (TodoListApi *CustomTodoListDBApi) GetCustomTodoListDBList(c *gin.Context) {
	var pageInfo todo_listReq.CustomTodoListDBSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := TodoListService.GetCustomTodoListDBInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
