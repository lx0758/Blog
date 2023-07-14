<template>
  <el-container class="login" direction="vertical">
    <p id="title">登录</p>
    <el-form
      status-icon
      label-width="0"
      ref="formRef"
      :model="formModel"
      :rules="formRules"
    >
      <el-form-item prop="username">
        <el-input
          class="item"
          v-model="formModel.username"
          auto-complete="off"
          placeholder="请输入用户名" />
      </el-form-item>
      <el-form-item prop="password">
        <el-input
          class="item"
          type="password"
          v-model="formModel.password"
          auto-complete="off"
          placeholder="请输入密码" />
      </el-form-item>
      <el-row type="flex" justify="space-between">
        <el-col :span="15">
          <el-form-item prop="captcha">
            <el-input
              class="item"
              v-model="formModel.captcha"
              auto-complete="off"
              placeholder="请输入验证码" />
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <img @click="onRefreshCaptcha()" style="width: 100%; height: 40px" :src="captchaUrl" />
        </el-col>
      </el-row>
      <el-form-item>
        <el-button class="item" type="primary" @click="onLogin()" style="width: 100%">登录</el-button>
      </el-form-item>
    </el-form>
  </el-container>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'
import { login } from '@/api'
import { useUserStore } from '@/stores'

interface LoginState {
  username: string
  password: string
  captcha: string
}

const router = useRouter()

const formRef = ref<FormInstance>()
const formModel = ref<LoginState>({} as LoginState)
const formRules = ref<FormRules<LoginState>>({
  username: [
    { required: true, message: '用户名不能为空', trigger: 'blur' },
    { min: 1, max: 16, message: '用户名长度不合法', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '密码不能为空', trigger: 'blur' },
    { min: 1, max: 16, message: '密码长度不合法', trigger: 'blur' }
  ],
  captcha: [
    { required: true, message: '验证码不能为空', trigger: 'blur' },
    { min: 1, max: 8, message: '验证码长度不合法', trigger: 'blur' }
  ],
})
const captchaUrl = ref('/api/captcha')

function onRefreshCaptcha() {
  captchaUrl.value = '/api/captcha?t=' + Math.random()
}
function onLogin() {
  if (formRef.value == undefined) return
  formRef.value.validate((valid: boolean) => {
    if (!valid) return
    login(
      formModel.value.username,
      formModel.value.password,
      formModel.value.captcha,
    ).then(() => {
        let userStore = useUserStore()
        userStore.login()
        ElMessage.success('登录成功')
        router.replace('/')
      })
      .catch(() => {
        formModel.value.captcha = ''
        onRefreshCaptcha()
      })
  })
}
function onKeyDown(e: KeyboardEvent) {
  if (e.keyCode === 13) onLogin()
}

onMounted(() => {
  document.addEventListener('keydown', onKeyDown)
  onRefreshCaptcha()
})
onUnmounted(() => {
    document.removeEventListener('keydown', onKeyDown)
})
</script>

<style>
.login {
  width: 400px;
  height: 320px;
  padding: 20px;

  margin: -170px 0 0 -200px;
  position: absolute;
  top: 50%;
  left: 50%;

  background: rgba(255, 255, 255, 1);
  box-shadow: 0 0 30px 0 rgba(7, 106, 228, 0.26);
  border-radius: 10px;
}

#title {
  font-size: 24px;
  font-weight: bold;
  margin: 10px;
  text-align: center;
}

.login .item {
  height: 40px !important;
}
</style>
