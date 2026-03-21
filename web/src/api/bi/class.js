import service from '@/utils/request'
// @Tags Class
// @Summary 创建班级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Class true "创建班级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /class/createClass [post]
export const createClass = (data) => {
  return service({
    url: '/class/createClass',
    method: 'post',
    data
  })
}

// @Tags Class
// @Summary 删除班级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Class true "删除班级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /class/deleteClass [delete]
export const deleteClass = (params) => {
  return service({
    url: '/class/deleteClass',
    method: 'delete',
    params
  })
}

// @Tags Class
// @Summary 批量删除班级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除班级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /class/deleteClass [delete]
export const deleteClassByIds = (params) => {
  return service({
    url: '/class/deleteClassByIds',
    method: 'delete',
    params
  })
}

// @Tags Class
// @Summary 更新班级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Class true "更新班级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /class/updateClass [put]
export const updateClass = (data) => {
  return service({
    url: '/class/updateClass',
    method: 'put',
    data
  })
}

// @Tags Class
// @Summary 用id查询班级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Class true "用id查询班级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /class/findClass [get]
export const findClass = (params) => {
  return service({
    url: '/class/findClass',
    method: 'get',
    params
  })
}

// @Tags Class
// @Summary 分页获取班级列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取班级列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /class/getClassList [get]
export const getClassList = (params) => {
  return service({
    url: '/class/getClassList',
    method: 'get',
    params
  })
}

// @Tags Class
// @Summary 不需要鉴权的班级接口
// @Accept application/json
// @Produce application/json
// @Param data query biReq.ClassSearch true "分页获取班级列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /class/getClassPublic [get]
export const getClassPublic = () => {
  return service({
    url: '/class/getClassPublic',
    method: 'get',
  })
}
