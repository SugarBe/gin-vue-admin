// 自动生成模板CustomTodoListDB
package todo_list

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// CustomTodoListDB 结构体
type CustomTodoListDB struct {
      global.GVA_MODEL
      Title  string `json:"title" form:"title" gorm:"column:title;comment:标题;"`
      Content  string `json:"content" form:"content" gorm:"column:content;comment:内容;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态;"`
      Expire_time  *time.Time `json:"expire_time" form:"expire_time" gorm:"column:expire_time;comment:过期时间;"`
      Executer  string `json:"executer" form:"executer" gorm:"column:executer;comment:执行人;"`
      Operator  string `json:"operator" form:"operator" gorm:"column:operator;comment:操作人;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName CustomTodoListDB 表名
func (CustomTodoListDB) TableName() string {
  return "custom_todo_list_d_b"
}

