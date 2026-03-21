package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CourseRecordSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	SerialNumber   *string     `json:"serialNumber" form:"serialNumber"`
	CourseName     *string     `json:"courseName" form:"courseName"`
	Performance    *string     `json:"performance" form:"performance"`
	Homework       *string     `json:"homework" form:"homework"`
	Remark         *string     `json:"remark" form:"remark"`
	StudentId      *int        `json:"studentId" form:"studentId"`
	request.PageInfo
}
