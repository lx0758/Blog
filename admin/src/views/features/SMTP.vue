<template>
  <el-form label-width="100px">
    <el-form-item label="启用">
      <el-switch v-model="state.enable" />
    </el-form-item>
    <el-form-item label="主机">
      <el-input v-model="state.hostname" />
    </el-form-item>
    <el-form-item label="端口">
      <el-input v-model="state.port" type="number" />
    </el-form-item>
    <el-form-item label="SSL">
      <el-switch v-model="state.ssl" />
    </el-form-item>
    <el-form-item label="用户名">
      <el-input v-model="state.username" />
    </el-form-item>
    <el-form-item label="密码">
      <el-input v-model="state.password" type="password" show-password />
    </el-form-item>
    <el-form-item label="发件人">
      <el-input v-model="state.fromName" />
    </el-form-item>
    <el-form-item label="发件邮箱">
      <el-input v-model="state.fromEmail" />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">保存</el-button>
      <el-button type="primary" @click="onTest">测试</el-button>
    </el-form-item>
  </el-form>

  <el-dialog title="发送测试邮件" v-model="sendEmailDialogState.isShow" :close-on-click-modal="false">
    <el-form ref="sendEmailDialogStateFormRef" :model="sendEmailDialogState.formModel" :rules="sendEmailDialogState.formRules" label-width="120px">
      <el-form-item label="邮箱" prop="email">
        <el-input v-model="sendEmailDialogState.formModel.email" placeholder="请输入名邮箱地址"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onDialogSubmit">发送</el-button>
        <el-button @click="sendEmailDialogState.hide()">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { DialogState } from '@/bean'
import { updateSMTP, querySMTP, testSMTP } from '@/api'

interface State {
  enable: boolean
  hostname: string|null
  port: number|null
  ssl: boolean|null
  username: string|null
  password: string|null
  fromName: string|null
  fromEmail: string|null
}

interface SendEmailState {
  email: string
}

const state = ref<State>({
  enable: false,
  hostname: null,
  port: null,
  ssl: false,
  username: null,
  password: null,
  fromName: null,
  fromEmail: null,
})

const sendEmailDialogStateFormRef = ref()
const sendEmailDialogState = ref(new DialogState<SendEmailState>(
  sendEmailDialogStateFormRef,
  {} as SendEmailState,
  {
    email: [
      { required: true, message: '邮箱不能为空', trigger: 'blur' }
    ]
  }
))

function onRefresh() {
  querySMTP().then((data) => {
    data = data.data
    state.value.enable = data.enable
    state.value.hostname = data.hostname
    state.value.port = data.port
    state.value.ssl = data.ssl
    state.value.username = data.username
    state.value.password = data.password
    state.value.fromName = data.fromName
    state.value.fromEmail = data.fromEmail
  })
}

function onSubmit() {
  updateSMTP(
    state.value.enable,
    state.value.hostname,
    state.value.port,
    state.value.ssl,
    state.value.username,
    state.value.password,
    state.value.fromName,
    state.value.fromEmail,
  ).then(() => {
    ElMessage.success('更新成功')
    onRefresh()
  })
}
function onTest() {
  sendEmailDialogState.value.show()
}
function onDialogSubmit() {
  if (sendEmailDialogState.value.formRef == undefined) return
  sendEmailDialogState.value.formRef.validate((valid: boolean) => {
    if (!valid) return
    let form = sendEmailDialogState.value.formModel
    testSMTP(form.email).then(() => {
      ElMessage.success('发送成功')
      sendEmailDialogState.value.reset()
      sendEmailDialogState.value.hide()
    })
  })
}

onMounted(() => {
  onRefresh()
})
</script>

<style></style>
