import service from '@/utils/request'
// @Tags ClassRecord
// @Summary 创建课程记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ClassRecord true "创建课程记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /classRecord/createClassRecord [post]
export const createClassRecord = (data) => {
  return service({
    url: '/classRecord/createClassRecord',
    method: 'post',
    data
  })
}

// @Tags ClassRecord
// @Summary 删除课程记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ClassRecord true "删除课程记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /classRecord/deleteClassRecord [delete]
export const deleteClassRecord = (params) => {
  return service({
    url: '/classRecord/deleteClassRecord',
    method: 'delete',
    params
  })
}

// @Tags ClassRecord
// @Summary 批量删除课程记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除课程记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /classRecord/deleteClassRecord [delete]
export const deleteClassRecordByIds = (params) => {
  return service({
    url: '/classRecord/deleteClassRecordByIds',
    method: 'delete',
    params
  })
}

// @Tags ClassRecord
// @Summary 更新课程记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ClassRecord true "更新课程记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /classRecord/updateClassRecord [put]
export const updateClassRecord = (data) => {
  return service({
    url: '/classRecord/updateClassRecord',
    method: 'put',
    data
  })
}

// @Tags ClassRecord
// @Summary 用id查询课程记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ClassRecord true "用id查询课程记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /classRecord/findClassRecord [get]
export const findClassRecord = (params) => {
  return service({
    url: '/classRecord/findClassRecord',
    method: 'get',
    params
  })
}

// @Tags ClassRecord
// @Summary 分页获取课程记录列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取课程记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /classRecord/getClassRecordList [get]
export const getClassRecordList = (params) => {
  return service({
    url: '/classRecord/getClassRecordList',
    method: 'get',
    params
  })
}

// @Tags ClassRecord
// @Summary 不需要鉴权的课程记录接口
// @Accept application/json
// @Produce application/json
// @Param data query biReq.ClassRecordSearch true "分页获取课程记录列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /classRecord/getClassRecordPublic [get]
export const getClassRecordPublic = () => {
  return service({
    url: '/classRecord/getClassRecordPublic',
    method: 'get',
  })
}
