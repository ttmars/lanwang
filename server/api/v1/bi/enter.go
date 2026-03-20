package bi

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ClassRecordApi
	StudentApi
}

var (
	classRecordService = service.ServiceGroupApp.BiServiceGroup.ClassRecordService
	studentService     = service.ServiceGroupApp.BiServiceGroup.StudentService
)
