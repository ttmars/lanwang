package bi

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bi"
	biReq "github.com/flipped-aurora/gin-vue-admin/server/model/bi/request"
	"gorm.io/gorm"
)

type ClassService struct{}

// CreateClass 创建班级记录
// Author [yourname](https://github.com/yourname)
func (classService *ClassService) CreateClass(ctx context.Context, class *bi.Class) (err error) {
	return global.GVA_DB.Create(&class).Error
}

// DeleteClass 删除班级记录
// Author [yourname](https://github.com/yourname)
func (classService *ClassService) DeleteClass(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&bi.Class{}, "id = ?", ID).Error
	return err
}

// DeleteClassByIds 批量删除班级记录
// Author [yourname](https://github.com/yourname)
func (classService *ClassService) DeleteClassByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]bi.Class{}, "id in ?", IDs).Error
	return err
}

// UpdateClass 更新班级记录
// Author [yourname](https://github.com/yourname)
func (classService *ClassService) UpdateClass(ctx context.Context, class bi.Class) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Updates(&class).Error; err != nil {
			return err
		}

		if err = tx.Model(&class).Association("Students").Replace(class.Students); err != nil {
			return err
		}
		return nil
	})
}

// GetClass 根据ID获取班级记录
// Author [yourname](https://github.com/yourname)
func (classService *ClassService) GetClass(ctx context.Context, ID string) (class bi.Class, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&class).Error
	return
}

// GetClassInfoList 分页获取班级记录
// Author [yourname](https://github.com/yourname)
func (classService *ClassService) GetClassInfoList(ctx context.Context, info biReq.ClassSearch) (list []bi.Class, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&bi.Class{})
	var classs []bi.Class
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.Name != nil && *info.Name != "" {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
	}
	if info.Remark != nil && *info.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+*info.Remark+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Preload("Students").Find(&classs).Error
	return classs, total, err
}
func (classService *ClassService) GetClassPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
