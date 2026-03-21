package bi

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	ClassRecordRouter
	StudentRouter
	ClassRouter
}

var (
	classRecordApi = api.ApiGroupApp.BiApiGroup.ClassRecordApi
	studentApi     = api.ApiGroupApp.BiApiGroup.StudentApi
	classApi       = api.ApiGroupApp.BiApiGroup.ClassApi
)
