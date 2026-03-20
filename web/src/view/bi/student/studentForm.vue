
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="姓名:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入姓名" />
</el-form-item>
        <el-form-item label="年龄:" prop="age">
    <el-input v-model.number="formData.age" :clearable="true" placeholder="请输入年龄" />
</el-form-item>
        <el-form-item label="住址:" prop="address">
    <el-input v-model="formData.address" :clearable="true" placeholder="请输入住址" />
</el-form-item>
        <el-form-item label="手机号:" prop="phone">
    <el-input v-model="formData.phone" :clearable="true" placeholder="请输入手机号" />
</el-form-item>
        <el-form-item label="年级:" prop="grade">
    <el-input v-model="formData.grade" :clearable="true" placeholder="请输入年级" />
</el-form-item>
        <el-form-item label="邮箱:" prop="email">
    <el-input v-model="formData.email" :clearable="true" placeholder="请输入邮箱" />
</el-form-item>
        <el-form-item label="家长姓名:" prop="parentName">
    <el-input v-model="formData.parentName" :clearable="true" placeholder="请输入家长姓名" />
</el-form-item>
        <el-form-item label="家长手机号:" prop="parentPhone">
    <el-input v-model="formData.parentPhone" :clearable="true" placeholder="请输入家长手机号" />
</el-form-item>
        <el-form-item label="备注:" prop="remark">
    <el-input v-model="formData.remark" :clearable="true" placeholder="请输入备注" />
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
  createStudent,
  updateStudent,
  findStudent
} from '@/api/bi/student'

defineOptions({
    name: 'StudentForm'
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
const formData = ref({
            name: '',
            age: undefined,
            address: '',
            phone: '',
            grade: '',
            email: '',
            parentName: '',
            parentPhone: '',
            remark: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findStudent({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
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
               res = await createStudent(formData.value)
               break
             case 'update':
               res = await updateStudent(formData.value)
               break
             default:
               res = await createStudent(formData.value)
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
