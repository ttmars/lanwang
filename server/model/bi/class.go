// 自动生成模板Class
package bi

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 班级 结构体  Class
type Class struct {
	global.GVA_MODEL
	Name     *string    `json:"name" form:"name" gorm:"comment:班级名称;column:name;size:255;"`     //班级名称
	Remark   *string    `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:512;"` //备注
	Students []*Student `json:"students" form:"students" gorm:"many2many:student_class;"`
}

// TableName 班级 Class自定义表名 class
func (Class) TableName() string {
	return "class"
}
