package bi

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bi"
	biReq "github.com/flipped-aurora/gin-vue-admin/server/model/bi/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ClassApi struct{}

// CreateClass 创建班级
// @Tags Class
// @Summary 创建班级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body bi.Class true "创建班级"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /class/createClass [post]
func (classApi *ClassApi) CreateClass(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var class bi.Class
	err := c.ShouldBindJSON(&class)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = classService.CreateClass(ctx, &class)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteClass 删除班级
// @Tags Class
// @Summary 删除班级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body bi.Class true "删除班级"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /class/deleteClass [delete]
func (classApi *ClassApi) DeleteClass(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := classService.DeleteClass(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteClassByIds 批量删除班级
// @Tags Class
// @Summary 批量删除班级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /class/deleteClassByIds [delete]
func (classApi *ClassApi) DeleteClassByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := classService.DeleteClassByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateClass 更新班级
// @Tags Class
// @Summary 更新班级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body bi.Class true "更新班级"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /class/updateClass [put]
func (classApi *ClassApi) UpdateClass(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var class bi.Class
	err := c.ShouldBindJSON(&class)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = classService.UpdateClass(ctx, class)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindClass 用id查询班级
// @Tags Class
// @Summary 用id查询班级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询班级"
// @Success 200 {object} response.Response{data=bi.Class,msg=string} "查询成功"
// @Router /class/findClass [get]
func (classApi *ClassApi) FindClass(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reclass, err := classService.GetClass(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reclass, c)
}

// GetClassList 分页获取班级列表
// @Tags Class
// @Summary 分页获取班级列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query biReq.ClassSearch true "分页获取班级列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /class/getClassList [get]
func (classApi *ClassApi) GetClassList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo biReq.ClassSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := classService.GetClassInfoList(ctx, pageInfo)
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

// GetClassPublic 不需要鉴权的班级接口
// @Tags Class
// @Summary 不需要鉴权的班级接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /class/getClassPublic [get]
func (classApi *ClassApi) GetClassPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	classService.GetClassPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的班级接口信息",
	}, "获取成功", c)
}
