package bi

import (
	"context"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bi"
	biReq "github.com/flipped-aurora/gin-vue-admin/server/model/bi/request"
	"gorm.io/gorm"
)

type CourseRecordService struct{}

// CreateCourseRecord 创建课程记录记录
// Author [yourname](https://github.com/yourname)
func (courseRecordService *CourseRecordService) CreateCourseRecord(ctx context.Context, courseRecord *bi.CourseRecord) (err error) {
	courseRecord.SerialNumber = new(string)
	*courseRecord.SerialNumber = time.Now().Format("20060102150405")
	courseRecord.StudentId = new(int64)
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for _, id := range courseRecord.StudentIds {
			courseRecord.ID = 0
			*courseRecord.StudentId = id
			if err = tx.Create(courseRecord).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// DeleteCourseRecord 删除课程记录记录
// Author [yourname](https://github.com/yourname)
func (courseRecordService *CourseRecordService) DeleteCourseRecord(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&bi.CourseRecord{}, "id = ?", ID).Error
	return err
}

// DeleteCourseRecordByIds 批量删除课程记录记录
// Author [yourname](https://github.com/yourname)
func (courseRecordService *CourseRecordService) DeleteCourseRecordByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]bi.CourseRecord{}, "id in ?", IDs).Error
	return err
}

// UpdateCourseRecord 更新课程记录记录
// Author [yourname](https://github.com/yourname)
func (courseRecordService *CourseRecordService) UpdateCourseRecord(ctx context.Context, courseRecord bi.CourseRecord) (err error) {
	err = global.GVA_DB.Model(&bi.CourseRecord{}).Where("id = ?", courseRecord.ID).Updates(&courseRecord).Error
	return err
}

// GetCourseRecord 根据ID获取课程记录记录
// Author [yourname](https://github.com/yourname)
func (courseRecordService *CourseRecordService) GetCourseRecord(ctx context.Context, ID string) (courseRecord bi.CourseRecord, err error) {
	err = global.GVA_DB.Where("id = ?", ID).Preload("Student").First(&courseRecord).Error
	return
}

// GetCourseRecordInfoList 分页获取课程记录记录
// Author [yourname](https://github.com/yourname)
func (courseRecordService *CourseRecordService) GetCourseRecordInfoList(ctx context.Context, info biReq.CourseRecordSearch) (list []bi.CourseRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&bi.CourseRecord{})
	var courseRecords []bi.CourseRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.SerialNumber != nil && *info.SerialNumber != "" {
		db = db.Where("serial_number LIKE ?", "%"+*info.SerialNumber+"%")
	}
	if info.CourseName != nil && *info.CourseName != "" {
		db = db.Where("course_name LIKE ?", "%"+*info.CourseName+"%")
	}
	if info.Performance != nil && *info.Performance != "" {
		db = db.Where("performance LIKE ?", "%"+*info.Performance+"%")
	}
	if info.Homework != nil && *info.Homework != "" {
		db = db.Where("homework LIKE ?", "%"+*info.Homework+"%")
	}
	if info.Remark != nil && *info.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+*info.Remark+"%")
	}
	if info.StudentId != nil {
		db = db.Where("student_id = ?", *info.StudentId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("ID desc").Preload("Student").Find(&courseRecords).Error
	return courseRecords, total, err
}
func (courseRecordService *CourseRecordService) GetCourseRecordPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
