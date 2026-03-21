package bi

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bi"
	biReq "github.com/flipped-aurora/gin-vue-admin/server/model/bi/request"
)

type StudentService struct{}

// CreateStudent 创建学生记录
// Author [yourname](https://github.com/yourname)
func (studentService *StudentService) CreateStudent(ctx context.Context, student *bi.Student) (err error) {
	err = global.GVA_DB.Create(student).Error
	return err
}

// DeleteStudent 删除学生记录
// Author [yourname](https://github.com/yourname)
func (studentService *StudentService) DeleteStudent(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&bi.Student{}, "id = ?", ID).Error
	return err
}

// DeleteStudentByIds 批量删除学生记录
// Author [yourname](https://github.com/yourname)
func (studentService *StudentService) DeleteStudentByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]bi.Student{}, "id in ?", IDs).Error
	return err
}

// UpdateStudent 更新学生记录
// Author [yourname](https://github.com/yourname)
func (studentService *StudentService) UpdateStudent(ctx context.Context, student bi.Student) (err error) {
	err = global.GVA_DB.Model(&bi.Student{}).Where("id = ?", student.ID).Updates(&student).Error
	return err
}

// GetStudent 根据ID获取学生记录
// Author [yourname](https://github.com/yourname)
func (studentService *StudentService) GetStudent(ctx context.Context, ID string) (student bi.Student, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&student).Error
	return
}

// GetStudentInfoList 分页获取学生记录
// Author [yourname](https://github.com/yourname)
func (studentService *StudentService) GetStudentInfoList(ctx context.Context, info biReq.StudentSearch) (list []bi.Student, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&bi.Student{})
	var students []bi.Student
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.Name != nil && *info.Name != "" {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
	}
	if info.Age != nil {
		db = db.Where("age = ?", *info.Age)
	}
	if info.Address != nil && *info.Address != "" {
		db = db.Where("address LIKE ?", "%"+*info.Address+"%")
	}
	if info.Phone != nil && *info.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+*info.Phone+"%")
	}
	if info.Grade != nil && *info.Grade != "" {
		db = db.Where("grade LIKE ?", "%"+*info.Grade+"%")
	}
	if info.Email != nil && *info.Email != "" {
		db = db.Where("email LIKE ?", "%"+*info.Email+"%")
	}
	if info.ParentName != nil && *info.ParentName != "" {
		db = db.Where("parent_name LIKE ?", "%"+*info.ParentName+"%")
	}
	if info.ParentPhone != nil && *info.ParentPhone != "" {
		db = db.Where("parent_phone LIKE ?", "%"+*info.ParentPhone+"%")
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

	err = db.Order("ID desc").Find(&students).Error
	return students, total, err
}
func (studentService *StudentService) GetStudentPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
