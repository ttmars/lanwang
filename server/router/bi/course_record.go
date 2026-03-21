package bi

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CourseRecordRouter struct{}

// InitCourseRecordRouter 初始化 课程记录 路由信息
func (s *CourseRecordRouter) InitCourseRecordRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	courseRecordRouter := Router.Group("courseRecord").Use(middleware.OperationRecord())
	courseRecordRouterWithoutRecord := Router.Group("courseRecord")
	courseRecordRouterWithoutAuth := PublicRouter.Group("courseRecord")
	{
		courseRecordRouter.POST("createCourseRecord", courseRecordApi.CreateCourseRecord)             // 新建课程记录
		courseRecordRouter.DELETE("deleteCourseRecord", courseRecordApi.DeleteCourseRecord)           // 删除课程记录
		courseRecordRouter.DELETE("deleteCourseRecordByIds", courseRecordApi.DeleteCourseRecordByIds) // 批量删除课程记录
		courseRecordRouter.PUT("updateCourseRecord", courseRecordApi.UpdateCourseRecord)              // 更新课程记录
	}
	{
		courseRecordRouterWithoutRecord.GET("findCourseRecord", courseRecordApi.FindCourseRecord)       // 根据ID获取课程记录
		courseRecordRouterWithoutRecord.GET("getCourseRecordList", courseRecordApi.GetCourseRecordList) // 获取课程记录列表
	}
	{
		courseRecordRouterWithoutAuth.GET("getCourseRecordPublic", courseRecordApi.GetCourseRecordPublic) // 课程记录开放接口
	}
}
