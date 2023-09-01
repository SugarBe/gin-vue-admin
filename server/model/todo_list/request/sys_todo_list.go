package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/todo_list"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CustomTodoListDBSearch struct{
    todo_list.CustomTodoListDB
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    StartExpire_time  *time.Time  `json:"startExpire_time" form:"startExpire_time"`
    EndExpire_time  *time.Time  `json:"endExpire_time" form:"endExpire_time"`
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
