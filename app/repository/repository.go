package repository

import (
	"go_scaffold/app/model"
	"gorm.io/gorm"
)

//	@method page
//	@description: 分页
//	@param page int 页码数
//	@param limit int 每页数量
//	@return func(db *gorm.DB) *gorm.DB
func page(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		if limit < 1 {
			limit = 10
		}
		// 计算偏移量
		offset := (page - 1) * limit
		// 返回组装结果
		return db.Offset(offset).Limit(page)
	}
}

// base repository method interface
type repository interface {
	modelFilter(param map[string]interface{}) *gorm.DB
	Add(model *model.Model) error
	Update(params map[string]interface{}, model *model.Model) error
	Delete(params map[string]interface{}, model *model.Model) error
	List(params map[string]interface{}) map[string]interface{}
	All(params map[string]interface{}) map[string]interface{}
}
