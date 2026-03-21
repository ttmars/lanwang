package bi

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	StudentApi
	ClassApi
}

var (
	studentService = service.ServiceGroupApp.BiServiceGroup.StudentService
	classService   = service.ServiceGroupApp.BiServiceGroup.ClassService
)
