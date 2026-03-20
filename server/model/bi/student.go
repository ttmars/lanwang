// 自动生成模板Student
package bi

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 学生 结构体  Student
type Student struct {
	global.GVA_MODEL
	Name        *string `json:"name" form:"name" gorm:"comment:姓名;column:name;size:255;"`                          //姓名
	Age         *int32  `json:"age" form:"age" gorm:"comment:年龄;column:age;"`                                      //年龄
	Address     *string `json:"address" form:"address" gorm:"comment:住址;column:address;size:512;"`                 //住址
	Phone       *string `json:"phone" form:"phone" gorm:"comment:手机号;column:phone;size:255;"`                      //手机号
	Grade       *string `json:"grade" form:"grade" gorm:"comment:年级;column:grade;size:255;"`                       //年级
	Email       *string `json:"email" form:"email" gorm:"comment:邮箱;column:email;size:255;"`                       //邮箱
	ParentName  *string `json:"parentName" form:"parentName" gorm:"comment:家长姓名;column:parent_name;size:255;"`     //家长姓名
	ParentPhone *string `json:"parentPhone" form:"parentPhone" gorm:"comment:家长手机号;column:parent_phone;size:255;"` //家长手机号
	Remark      *string `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:512;"`                    //备注
}

// TableName 学生 Student自定义表名 student
func (Student) TableName() string {
	return "student"
}
