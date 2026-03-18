package bi

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ ClassRecordApi }

var classRecordService = service.ServiceGroupApp.BiServiceGroup.ClassRecordService
