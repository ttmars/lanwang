package bi

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bi"
	biReq "github.com/flipped-aurora/gin-vue-admin/server/model/bi/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ClassRecordApi struct{}

// CreateClassRecord 创建课程记录
// @Tags ClassRecord
// @Summary 创建课程记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body bi.ClassRecord true "创建课程记录"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /classRecord/createClassRecord [post]
func (classRecordApi *ClassRecordApi) CreateClassRecord(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var classRecord bi.ClassRecord
	err := c.ShouldBindJSON(&classRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = classRecordService.CreateClassRecord(ctx, &classRecord)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteClassRecord 删除课程记录
// @Tags ClassRecord
// @Summary 删除课程记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body bi.ClassRecord true "删除课程记录"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /classRecord/deleteClassRecord [delete]
func (classRecordApi *ClassRecordApi) DeleteClassRecord(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := classRecordService.DeleteClassRecord(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteClassRecordByIds 批量删除课程记录
// @Tags ClassRecord
// @Summary 批量删除课程记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /classRecord/deleteClassRecordByIds [delete]
func (classRecordApi *ClassRecordApi) DeleteClassRecordByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := classRecordService.DeleteClassRecordByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateClassRecord 更新课程记录
// @Tags ClassRecord
// @Summary 更新课程记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body bi.ClassRecord true "更新课程记录"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /classRecord/updateClassRecord [put]
func (classRecordApi *ClassRecordApi) UpdateClassRecord(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var classRecord bi.ClassRecord
	err := c.ShouldBindJSON(&classRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = classRecordService.UpdateClassRecord(ctx, classRecord)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindClassRecord 用id查询课程记录
// @Tags ClassRecord
// @Summary 用id查询课程记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询课程记录"
// @Success 200 {object} response.Response{data=bi.ClassRecord,msg=string} "查询成功"
// @Router /classRecord/findClassRecord [get]
func (classRecordApi *ClassRecordApi) FindClassRecord(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reclassRecord, err := classRecordService.GetClassRecord(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reclassRecord, c)
}

// GetClassRecordList 分页获取课程记录列表
// @Tags ClassRecord
// @Summary 分页获取课程记录列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query biReq.ClassRecordSearch true "分页获取课程记录列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /classRecord/getClassRecordList [get]
func (classRecordApi *ClassRecordApi) GetClassRecordList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo biReq.ClassRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := classRecordService.GetClassRecordInfoList(ctx, pageInfo)
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

// GetClassRecordPublic 不需要鉴权的课程记录接口
// @Tags ClassRecord
// @Summary 不需要鉴权的课程记录接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /classRecord/getClassRecordPublic [get]
func (classRecordApi *ClassRecordApi) GetClassRecordPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	classRecordService.GetClassRecordPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的课程记录接口信息",
	}, "获取成功", c)
}
