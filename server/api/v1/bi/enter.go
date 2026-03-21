package bi

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ClassRecordApi
	StudentApi
	ClassApi
}

var (
	classRecordService = service.ServiceGroupApp.BiServiceGroup.ClassRecordService
	studentService     = service.ServiceGroupApp.BiServiceGroup.StudentService
	classService       = service.ServiceGroupApp.BiServiceGroup.ClassService
)
