package bi

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ClassRouter struct{}

// InitClassRouter 初始化 班级 路由信息
func (s *ClassRouter) InitClassRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	classRouter := Router.Group("class").Use(middleware.OperationRecord())
	classRouterWithoutRecord := Router.Group("class")
	classRouterWithoutAuth := PublicRouter.Group("class")
	{
		classRouter.POST("createClass", classApi.CreateClass)             // 新建班级
		classRouter.DELETE("deleteClass", classApi.DeleteClass)           // 删除班级
		classRouter.DELETE("deleteClassByIds", classApi.DeleteClassByIds) // 批量删除班级
		classRouter.PUT("updateClass", classApi.UpdateClass)              // 更新班级
	}
	{
		classRouterWithoutRecord.GET("findClass", classApi.FindClass)       // 根据ID获取班级
		classRouterWithoutRecord.GET("getClassList", classApi.GetClassList) // 获取班级列表
	}
	{
		classRouterWithoutAuth.GET("getClassPublic", classApi.GetClassPublic) // 班级开放接口
	}
}
