
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="上课编号:" prop="serialNumber">
    <el-input v-model="formData.serialNumber" :clearable="true" placeholder="请输入上课编号" />
</el-form-item>
        <el-form-item label="课程名称:" prop="courseName">
    <el-tree-select v-model="formData.courseName" placeholder="请选择课程名称" :data="courseOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
        <el-form-item label="上课表现:" prop="performance">
    <el-input v-model="formData.performance" :clearable="true" placeholder="请输入上课表现" />
</el-form-item>
        <el-form-item label="作业:" prop="homework">
    <el-input v-model="formData.homework" :clearable="true" placeholder="请输入作业" />
</el-form-item>
        <el-form-item label="备注:" prop="remark">
    <el-input v-model="formData.remark" :clearable="true" placeholder="请输入备注" />
</el-form-item>
        <el-form-item label="学生ID:" prop="studentId">
    <el-input v-model.number="formData.studentId" :clearable="true" placeholder="请输入学生ID" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createCourseRecord,
  updateCourseRecord,
  findCourseRecord
} from '@/api/bi/courseRecord'

defineOptions({
    name: 'CourseRecordForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const courseOptions = ref([])
const formData = ref({
            serialNumber: '',
            courseName: '',
            performance: '',
            homework: '',
            remark: '',
            studentId: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCourseRecord({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    courseOptions.value = await getDictFunc('course')
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
