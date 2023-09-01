package todo_list

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/todo_list"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    todo_listReq "github.com/flipped-aurora/gin-vue-admin/server/model/todo_list/request"
    "gorm.io/gorm"
)

type CustomTodoListDBService struct {
}

// CreateCustomTodoListDB 创建CustomTodoListDB记录
// Author [piexlmax](https://github.com/piexlmax)
func (TodoListService *CustomTodoListDBService) CreateCustomTodoListDB(TodoList *todo_list.CustomTodoListDB) (err error) {
	err = global.GVA_DB.Create(TodoList).Error
	return err
}

// DeleteCustomTodoListDB 删除CustomTodoListDB记录
// Author [piexlmax](https://github.com/piexlmax)
func (TodoListService *CustomTodoListDBService)DeleteCustomTodoListDB(TodoList todo_list.CustomTodoListDB) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&todo_list.CustomTodoListDB{}).Where("id = ?", TodoList.ID).Update("deleted_by", TodoList.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&TodoList).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteCustomTodoListDBByIds 批量删除CustomTodoListDB记录
// Author [piexlmax](https://github.com/piexlmax)
func (TodoListService *CustomTodoListDBService)DeleteCustomTodoListDBByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&todo_list.CustomTodoListDB{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&todo_list.CustomTodoListDB{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateCustomTodoListDB 更新CustomTodoListDB记录
// Author [piexlmax](https://github.com/piexlmax)
func (TodoListService *CustomTodoListDBService)UpdateCustomTodoListDB(TodoList todo_list.CustomTodoListDB) (err error) {
	err = global.GVA_DB.Save(&TodoList).Error
	return err
}

// GetCustomTodoListDB 根据id获取CustomTodoListDB记录
// Author [piexlmax](https://github.com/piexlmax)
func (TodoListService *CustomTodoListDBService)GetCustomTodoListDB(id uint) (TodoList todo_list.CustomTodoListDB, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&TodoList).Error
	return
}

// GetCustomTodoListDBInfoList 分页获取CustomTodoListDB记录
// Author [piexlmax](https://github.com/piexlmax)
func (TodoListService *CustomTodoListDBService)GetCustomTodoListDBInfoList(info todo_listReq.CustomTodoListDBSearch) (list []todo_list.CustomTodoListDB, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&todo_list.CustomTodoListDB{})
    var TodoLists []todo_list.CustomTodoListDB
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Title != "" {
        db = db.Where("title LIKE ?","%"+ info.Title+"%")
    }
    if info.Content != "" {
        db = db.Where("content LIKE ?","%"+ info.Content+"%")
    }
    if info.Status != nil {
        db = db.Where("status = ?",info.Status)
    }
        if info.StartExpire_time != nil && info.EndExpire_time != nil {
            db = db.Where("expire_time BETWEEN ? AND ? ",info.StartExpire_time,info.EndExpire_time)
        }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["status"] = true
         	orderMap["expire_time"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&TodoLists).Error
	return  TodoLists, total, err
}
