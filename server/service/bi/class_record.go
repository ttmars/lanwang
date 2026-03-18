package bi

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bi"
	biReq "github.com/flipped-aurora/gin-vue-admin/server/model/bi/request"
)

type ClassRecordService struct{}

// CreateClassRecord 创建课程记录记录
// Author [yourname](https://github.com/yourname)
func (classRecordService *ClassRecordService) CreateClassRecord(ctx context.Context, classRecord *bi.ClassRecord) (err error) {
	err = global.GVA_DB.Create(classRecord).Error
	return err
}

// DeleteClassRecord 删除课程记录记录
// Author [yourname](https://github.com/yourname)
func (classRecordService *ClassRecordService) DeleteClassRecord(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&bi.ClassRecord{}, "id = ?", ID).Error
	return err
}

// DeleteClassRecordByIds 批量删除课程记录记录
// Author [yourname](https://github.com/yourname)
func (classRecordService *ClassRecordService) DeleteClassRecordByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]bi.ClassRecord{}, "id in ?", IDs).Error
	return err
}

// UpdateClassRecord 更新课程记录记录
// Author [yourname](https://github.com/yourname)
func (classRecordService *ClassRecordService) UpdateClassRecord(ctx context.Context, classRecord bi.ClassRecord) (err error) {
	err = global.GVA_DB.Model(&bi.ClassRecord{}).Where("id = ?", classRecord.ID).Updates(&classRecord).Error
	return err
}

// GetClassRecord 根据ID获取课程记录记录
// Author [yourname](https://github.com/yourname)
func (classRecordService *ClassRecordService) GetClassRecord(ctx context.Context, ID string) (classRecord bi.ClassRecord, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&classRecord).Error
	return
}

// GetClassRecordInfoList 分页获取课程记录记录
// Author [yourname](https://github.com/yourname)
func (classRecordService *ClassRecordService) GetClassRecordInfoList(ctx context.Context, info biReq.ClassRecordSearch) (list []bi.ClassRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&bi.ClassRecord{})
	var classRecords []bi.ClassRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.StudentName != nil && *info.StudentName != "" {
		db = db.Where("student_name LIKE ?", "%"+*info.StudentName+"%")
	}
	if info.ClassName != nil && *info.ClassName != "" {
		db = db.Where("class_name LIKE ?", "%"+*info.ClassName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&classRecords).Error
	return classRecords, total, err
}
func (classRecordService *ClassRecordService) GetClassRecordPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
