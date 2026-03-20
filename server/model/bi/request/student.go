package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type StudentSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Name           *string     `json:"name" form:"name"`
	Age            *int        `json:"age" form:"age"`
	Address        *string     `json:"address" form:"address"`
	Phone          *string     `json:"phone" form:"phone"`
	Grade          *string     `json:"grade" form:"grade"`
	Email          *string     `json:"email" form:"email"`
	ParentName     *string     `json:"parentName" form:"parentName"`
	ParentPhone    *string     `json:"parentPhone" form:"parentPhone"`
	Remark         *string     `json:"remark" form:"remark"`
	request.PageInfo
}
