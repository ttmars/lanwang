package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ClassRecordSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	StudentName    *string     `json:"studentName" form:"studentName"`
	ClassName      *string     `json:"className" form:"className"`
	request.PageInfo
}
