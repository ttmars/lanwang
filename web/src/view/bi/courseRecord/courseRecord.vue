
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAtRange">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>

      <el-date-picker
            v-model="searchInfo.createdAtRange"
            class="!w-380px"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
       </el-form-item>
      
            <el-form-item label="上课编号" prop="serialNumber">
  <el-input v-model="searchInfo.serialNumber" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="课程名称" prop="courseName">
    <el-tree-select v-model="searchInfo.courseName" placeholder="请选择课程名称" :data="courseOptions" style="width:100%" filterable :clearable="true" check-strictly ></el-tree-select>
</el-form-item>
            
            <el-form-item label="上课表现" prop="performance">
  <el-input v-model="searchInfo.performance" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="作业" prop="homework">
  <el-input v-model="searchInfo.homework" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="备注" prop="remark">
  <el-input v-model="searchInfo.remark" placeholder="搜索条件" />
</el-form-item>
            

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column align="left" label="上课编号" prop="serialNumber" width="140" />

            <el-table-column align="left" label="学生" prop="studentIds" width="120">
              <template #default="scope">
                {{ scope.row.Student ? scope.row.Student.name : '' }}
              </template>
            </el-table-column>

            <el-table-column align="left" label="课程名称" prop="courseName" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.courseName,courseOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="上课表现" prop="performance" width="300" />

            <el-table-column align="left" label="作业" prop="homework" width="300" />

            <el-table-column align="left" label="备注" prop="remark" width="120" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateCourseRecordFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            <el-button   type="primary" link icon="copy" @click="copyRow(scope.row)">复制</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item v-if="type === 'create'" label="课程名称:" prop="courseName">
    <el-tree-select v-model="formData.courseName" placeholder="请选择课程名称" :data="courseOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
            <el-form-item v-if="type === 'update'" label="上课表现:" prop="performance">
    <el-input v-model="formData.performance" :clearable="true" placeholder="请输入上课表现" />
</el-form-item>

            <el-form-item v-if="type === 'create'" label="学生:" prop="studentIds">
              <el-select v-model="formData.studentIds" multiple placeholder="请选择学生">
                <el-option v-for="item in studentOptions" :key="item.value" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item v-if="type === 'create'" label="班级:" prop="classIds">
              <el-select v-model="formData.classIds" multiple placeholder="请选择班级" @change="onClassChange">
                <el-option v-for="item in classOptions" :key="item.value" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>

            <el-form-item label="作业:" prop="homework">
              <el-input v-model="formData.homework" :clearable="true" placeholder="请输入作业" />
            </el-form-item>
            <el-form-item v-if="type === 'update'" label="备注:" prop="remark">
    <el-input v-model="formData.remark" :clearable="true" placeholder="请输入备注" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="上课编号">
    {{ detailForm.serialNumber }}
</el-descriptions-item>
                    <el-descriptions-item label="学生">
    {{ detailForm.Student ? detailForm.Student.name : '' }}
</el-descriptions-item>
                    <el-descriptions-item label="课程名称">
    {{ detailForm.courseName }}
</el-descriptions-item>
                    <el-descriptions-item label="上课表现">
    {{ detailForm.performance }}
</el-descriptions-item>
                    <el-descriptions-item label="作业">
    {{ detailForm.homework }}
</el-descriptions-item>
                    <el-descriptions-item label="备注">
    {{ detailForm.remark }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createCourseRecord,
  deleteCourseRecord,
  deleteCourseRecordByIds,
  updateCourseRecord,
  findCourseRecord,
  getCourseRecordList
} from '@/api/bi/courseRecord'
import { getStudentList } from '@/api/bi/student'
import { getClassList } from '@/api/bi/class'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'CourseRecord'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const courseOptions = ref([])
const studentOptions = ref([])
const classOptions = ref([])
const studentMap = ref({})
const formData = ref({
            serialNumber: '',
            courseName: '',
            performance: '',
            homework: '',
            remark: '',
            studentIds: [],
            classIds: [],
        })



// 验证规则
const rule = reactive({
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getCourseRecordList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    courseOptions.value = await getDictFunc('course')
    const studentRes = await getStudentList({ page: 1, pageSize: 1000 })
    if (studentRes.code === 0) {
        studentOptions.value = studentRes.data.list.map(s => ({ label: s.name, value: s.ID }))
        studentMap.value = studentRes.data.list.reduce((map, s) => {
            map[s.ID] = s.name
            return map
        }, {})
    }
    const classRes = await getClassList({ page: 1, pageSize: 1000 })
    if (classRes.code === 0) {
        classOptions.value = classRes.data.list.map(c => ({ label: c.name, value: c.ID, students: c.students }))
    }
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 当选择班级时，添加该班级的学生到 studentIds
const onClassChange = (selectedClassIds) => {
    const selectedStudents = new Set(formData.value.studentIds)
    selectedClassIds.forEach(classId => {
        const classItem = classOptions.value.find(c => c.value === classId)
        if (classItem && classItem.students) {
            classItem.students.forEach(student => {
                if (student && student.ID) {
                    selectedStudents.add(student.ID)
                }
            })
        }
    })
    formData.value.studentIds = Array.from(selectedStudents)
}


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteCourseRecordFunc(row)
        })
    }

// 复制行
const copyRow = async (row) => {
    const studentName = row.Student ? row.Student.name : ''
    const courseName = filterDict(row.courseName, courseOptions) || row.courseName || ''
    const performance = row.performance || ''
    const homework = row.homework || ''
    const remark = row.remark || ''
    const text = `学生：${studentName}\n课程：${courseName}\n上课表现：${performance}\n作业：${homework}\n备注：${remark}`
    try {
        await navigator.clipboard.writeText(text)
        ElMessage({
            type: 'success',
            message: '复制成功'
        })
    } catch (err) {
        ElMessage({
            type: 'error',
            message: '复制失败'
        })
    }
}

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteCourseRecordByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateCourseRecordFunc = async(row) => {
    const res = await findCourseRecord({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        formData.value.studentIds = res.data.studentIds || (res.data.studentId ? [res.data.studentId] : [])
        formData.value.classIds = []
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteCourseRecordFunc = async (row) => {
    const res = await deleteCourseRecord({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        serialNumber: '',
        courseName: '',
        performance: '',
        homework: '',
        remark: '',
        studentIds: [],
        classIds: [],
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createCourseRecord(formData.value)
                  break
                case 'update':
                  res = await updateCourseRecord(formData.value)
                  break
                default:
                  res = await createCourseRecord(formData.value)
                  break
              }
              btnLoading.value = false
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findCourseRecord({ ID: row.ID })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}


</script>

<style>

</style>
