<template>
  <el-container class="login" direction="vertical">
    <p id="title">登录</p>
    <el-form :model="account" :rules="rules" ref="account" status-icon label-width="0">
      <el-form-item prop="username">
        <el-input v-model="account.username" auto-complete="off" placeholder="请输入用户名"></el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input type="password" v-model="account.password" auto-complete="off" placeholder="请输入密码"></el-input>
      </el-form-item>
      <el-row type="flex" justify="space-between">
        <el-col :span="15">
          <el-form-item prop="captcha">
            <el-input v-model="account.captcha" auto-complete="off" placeholder="请输入验证码"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <img @click="onRefreshCaptcha()" style="width: 100%; height: 40px" :src="captchaUrl"/>
        </el-col>
      </el-row>
      <el-form-item>
        <el-button type="primary" @click="onLogin()" style="width:100%;">登录</el-button>
      </el-form-item>
    </el-form>
  </el-container>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import {login} from "@/api"

export default defineComponent({
  name: 'Login',
  components: {},
  data() {
    return {
      account: {
        username: '',
        password: '',
        captcha: '',
      },
      rules: {
        username: [
          {required: true, message: '用户名不能为空', trigger: 'blur'},
          {min: 1, max: 16, message: '用户名长度不合法', trigger: 'blur'}
        ],
        password: [
          {required: true, message: '密码不能为空', trigger: 'blur'},
          {min: 1, max: 16, message: '密码长度不合法', trigger: 'blur'}
        ],
        captcha: [
          {required: true, message: '验证码不能为空', trigger: 'blur'},
          {min: 1, max: 8, message: '验证码长度不合法', trigger: 'blur'}
        ],
      },
      captchaUrl: '',
    }
  },
  mounted() {
    document.addEventListener('keydown', this.onKeyDown);
    this.onRefreshCaptcha()
  },
  unmounted() {
    document.removeEventListener('keydown', this.onKeyDown);
  },
  methods: {
    onRefreshCaptcha() {
      this.captchaUrl = '/api/captcha?t=' + Math.random()
    },
    onLogin() {
      let from: any = this.$refs['account'];
      from.validate((valid: boolean) => {
        if (!valid) return
        login(this.account.username, this.account.password, this.account.captcha)
            .then(data => {
              localStorage.setItem("isLogin", String(true));
              this.$message.success("登录成功")
              this.$router.replace('/')
            })
            .catch(() => {
              this.account.captcha = ''
              this.onRefreshCaptcha()
            })
      })
    },
    onKeyDown(e:KeyboardEvent) {
      if (e.keyCode === 13) this.onLogin();
    },
  },
});
</script>

<style lang="less">
.login {
  width: 400px;
  height: 340px;
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
</style>