import service from '@/utils/request'
// @Tags CourseRecord
// @Summary 创建课程记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CourseRecord true "创建课程记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /courseRecord/createCourseRecord [post]
export const createCourseRecord = (data) => {
  return service({
    url: '/courseRecord/createCourseRecord',
    method: 'post',
    data
  })
}

// @Tags CourseRecord
// @Summary 删除课程记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CourseRecord true "删除课程记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /courseRecord/deleteCourseRecord [delete]
export const deleteCourseRecord = (params) => {
  return service({
    url: '/courseRecord/deleteCourseRecord',
    method: 'delete',
    params
  })
}

// @Tags CourseRecord
// @Summary 批量删除课程记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除课程记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /courseRecord/deleteCourseRecord [delete]
export const deleteCourseRecordByIds = (params) => {
  return service({
    url: '/courseRecord/deleteCourseRecordByIds',
    method: 'delete',
    params
  })
}

// @Tags CourseRecord
// @Summary 更新课程记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CourseRecord true "更新课程记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /courseRecord/updateCourseRecord [put]
export const updateCourseRecord = (data) => {
  return service({
    url: '/courseRecord/updateCourseRecord',
    method: 'put',
    data
  })
}

// @Tags CourseRecord
// @Summary 用id查询课程记录
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.CourseRecord true "用id查询课程记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /courseRecord/findCourseRecord [get]
export const findCourseRecord = (params) => {
  return service({
    url: '/courseRecord/findCourseRecord',
    method: 'get',
    params
  })
}

// @Tags CourseRecord
// @Summary 分页获取课程记录列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取课程记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /courseRecord/getCourseRecordList [get]
export const getCourseRecordList = (params) => {
  return service({
    url: '/courseRecord/getCourseRecordList',
    method: 'get',
    params
  })
}

// @Tags CourseRecord
// @Summary 不需要鉴权的课程记录接口
// @Accept application/json
// @Produce application/json
// @Param data query biReq.CourseRecordSearch true "分页获取课程记录列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /courseRecord/getCourseRecordPublic [get]
export const getCourseRecordPublic = () => {
  return service({
    url: '/courseRecord/getCourseRecordPublic',
    method: 'get',
  })
}
