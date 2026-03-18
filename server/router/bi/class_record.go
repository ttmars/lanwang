package bi

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ClassRecordRouter struct{}

// InitClassRecordRouter 初始化 课程记录 路由信息
func (s *ClassRecordRouter) InitClassRecordRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	classRecordRouter := Router.Group("classRecord").Use(middleware.OperationRecord())
	classRecordRouterWithoutRecord := Router.Group("classRecord")
	classRecordRouterWithoutAuth := PublicRouter.Group("classRecord")
	{
		classRecordRouter.POST("createClassRecord", classRecordApi.CreateClassRecord)             // 新建课程记录
		classRecordRouter.DELETE("deleteClassRecord", classRecordApi.DeleteClassRecord)           // 删除课程记录
		classRecordRouter.DELETE("deleteClassRecordByIds", classRecordApi.DeleteClassRecordByIds) // 批量删除课程记录
		classRecordRouter.PUT("updateClassRecord", classRecordApi.UpdateClassRecord)              // 更新课程记录
	}
	{
		classRecordRouterWithoutRecord.GET("findClassRecord", classRecordApi.FindClassRecord)       // 根据ID获取课程记录
		classRecordRouterWithoutRecord.GET("getClassRecordList", classRecordApi.GetClassRecordList) // 获取课程记录列表
	}
	{
		classRecordRouterWithoutAuth.GET("getClassRecordPublic", classRecordApi.GetClassRecordPublic) // 课程记录开放接口
	}
}
