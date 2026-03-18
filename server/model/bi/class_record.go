// 自动生成模板ClassRecord
package bi

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 课程记录 结构体  ClassRecord
type ClassRecord struct {
	global.GVA_MODEL
	StudentName      *string `json:"studentName" form:"studentName" gorm:"comment:学生姓名;column:student_name;size:255;"`                 //学生姓名
	ClassName        *string `json:"className" form:"className" gorm:"comment:课程名称;column:class_name;size:255;"`                       //课程名称
	ClassPerformance *string `json:"classPerformance" form:"classPerformance" gorm:"comment:上课表现;column:class_performance;size:1024;"` //上课表现
	Homework         *string `json:"homework" form:"homework" gorm:"comment:作业;column:homework;size:1024;"`                            //作业
	Remark           *string `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:255;"`                                   //备注
}

// TableName 课程记录 ClassRecord自定义表名 class_record
func (ClassRecord) TableName() string {
	return "class_record"
}
