package bi

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bi"
	biReq "github.com/flipped-aurora/gin-vue-admin/server/model/bi/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CourseRecordApi struct{}

// CreateCourseRecord 创建课程记录
// @Tags CourseRecord
// @Summary 创建课程记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body bi.CourseRecord true "创建课程记录"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /courseRecord/createCourseRecord [post]
func (courseRecordApi *CourseRecordApi) CreateCourseRecord(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var courseRecord bi.CourseRecord
	err := c.ShouldBindJSON(&courseRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = courseRecordService.CreateCourseRecord(ctx, &courseRecord)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteCourseRecord 删除课程记录
// @Tags CourseRecord
// @Summary 删除课程记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body bi.CourseRecord true "删除课程记录"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /courseRecord/deleteCourseRecord [delete]
func (courseRecordApi *CourseRecordApi) DeleteCourseRecord(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := courseRecordService.DeleteCourseRecord(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCourseRecordByIds 批量删除课程记录
// @Tags CourseRecord
// @Summary 批量删除课程记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /courseRecord/deleteCourseRecordByIds [delete]
func (courseRecordApi *CourseRecordApi) DeleteCourseRecordByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := courseRecordService.DeleteCourseRecordByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCourseRecord 更新课程记录
// @Tags CourseRecord
// @Summary 更新课程记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body bi.CourseRecord true "更新课程记录"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /courseRecord/updateCourseRecord [put]
func (courseRecordApi *CourseRecordApi) UpdateCourseRecord(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var courseRecord bi.CourseRecord
	err := c.ShouldBindJSON(&courseRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = courseRecordService.UpdateCourseRecord(ctx, courseRecord)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCourseRecord 用id查询课程记录
// @Tags CourseRecord
// @Summary 用id查询课程记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询课程记录"
// @Success 200 {object} response.Response{data=bi.CourseRecord,msg=string} "查询成功"
// @Router /courseRecord/findCourseRecord [get]
func (courseRecordApi *CourseRecordApi) FindCourseRecord(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	recourseRecord, err := courseRecordService.GetCourseRecord(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(recourseRecord, c)
}

// GetCourseRecordList 分页获取课程记录列表
// @Tags CourseRecord
// @Summary 分页获取课程记录列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query biReq.CourseRecordSearch true "分页获取课程记录列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /courseRecord/getCourseRecordList [get]
func (courseRecordApi *CourseRecordApi) GetCourseRecordList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo biReq.CourseRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := courseRecordService.GetCourseRecordInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetCourseRecordPublic 不需要鉴权的课程记录接口
// @Tags CourseRecord
// @Summary 不需要鉴权的课程记录接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /courseRecord/getCourseRecordPublic [get]
func (courseRecordApi *CourseRecordApi) GetCourseRecordPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	courseRecordService.GetCourseRecordPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的课程记录接口信息",
	}, "获取成功", c)
}
