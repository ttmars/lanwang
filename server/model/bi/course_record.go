// 自动生成模板CourseRecord
package bi

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 课程记录 结构体  CourseRecord
type CourseRecord struct {
	global.GVA_MODEL
	SerialNumber *string `json:"serialNumber" form:"serialNumber" gorm:"comment:上课编号;column:serial_number;size:255;"` //上课编号
	CourseName   *string `json:"courseName" form:"courseName" gorm:"comment:课程名称;column:course_name;size:255;"`       //课程名称
	Performance  *string `json:"performance" form:"performance" gorm:"comment:上课表现;column:performance;size:512;"`     //上课表现
	Homework     *string `json:"homework" form:"homework" gorm:"comment:作业;column:homework;size:512;"`                //作业
	Remark       *string `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:512;"`                      //备注
	StudentId    *int64  `json:"studentId" form:"studentId" gorm:"comment:学生ID;column:student_id;"`                   //学生ID
	Student      Student
	StudentIds   []int64 `json:"studentIds" form:"studentIds" gorm:"-"` //多个学生
}

// TableName 课程记录 CourseRecord自定义表名 course_record
func (CourseRecord) TableName() string {
	return "course_record"
}
