<template>
  <el-container style="height: 100%">

    <el-aside unique-opened="true" style="width: auto">
      <el-menu :default-active="nowRoute" style="height: 100%" unique-opened router :collapse="isCollapse">
        <el-menu-item index="dashboard">
          <i class="el-icon-s-data"></i>
          <span>仪表盘</span>
        </el-menu-item>
        <el-submenu index="content">
          <template #title>
            <i class="el-icon-menu"></i>
            <span>内容</span>
          </template>
          <el-menu-item index="article">
            <i class="el-icon-s-order"></i>
            <span>文章</span>
          </el-menu-item>
          <el-menu-item index="comment">
            <i class="el-icon-s-comment"></i>
            <span>评论</span>
          </el-menu-item>
          <el-menu-item index="category">
            <i class="el-icon-s-ticket"></i>
            <span>分类</span>
          </el-menu-item>
          <el-menu-item index="tag">
            <i class="el-icon-s-flag"></i>
            <span>标签</span>
          </el-menu-item>
        </el-submenu>
        <el-submenu index="other">
          <template #title>
            <i class="el-icon-menu"></i>
            <span>其他</span>
          </template>
          <el-menu-item index="upload">
            <i class="el-icon-picture"></i>
            <span>文件</span>
          </el-menu-item>
          <el-menu-item index="link">
            <i class="el-icon-s-promotion"></i>
            <span>友链</span>
          </el-menu-item>
          <el-menu-item index="url">
            <i class="el-icon-info"></i>
            <span>短链</span>
          </el-menu-item>
        </el-submenu>
        <el-menu-item index="setting">
          <i class="el-icon-s-tools"></i>
          <span>设置</span>
        </el-menu-item>
      </el-menu>
    </el-aside>

    <el-container>

      <el-header style="box-shadow: 0 -8px 20px 2px #DEDEE3;">
        <el-row type="flex" align="middle" style="height: 100%">
          <el-col :span="16">
            <el-container style="align-items: center">
              <div @click="onToggleMenu()" style="padding: 10px 10px 10px 0; font-size: 20px">
                <i v-if="!isCollapse" class="el-icon-s-unfold"></i>
                <i v-if="isCollapse" class="el-icon-s-fold"></i>
              </div>
              <el-breadcrumb separator="/">
                <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
                <el-breadcrumb-item>{{ this.$route.name }}</el-breadcrumb-item>
              </el-breadcrumb>
            </el-container>
          </el-col>
          <el-col :span="8">
            <el-container style="align-items: center; float: right;">
              <i class="el-icon-full-screen" @click="onFullScreen()" style="padding: 10px; font-size: 20px"></i>
              <i class="el-icon-refresh-right" @click="onRefresh()" style="padding: 10px; font-size: 20px"></i>
              <el-dropdown @command="onOption" style="padding: 10px">
                <span class="el-dropdown-link">
                  {{ user.nickname }}<i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item icon="el-icon-user" command="profile">个人设置</el-dropdown-item>
                    <el-dropdown-item icon="el-icon-lock" command="password">修改密码</el-dropdown-item>
                    <el-dropdown-item icon="el-icon-switch-button" command="logout" divided>注销登录</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
              <el-avatar :size="36" :src="user.avatar" v-if="user.avatar.length > 0"></el-avatar>
              <el-avatar :size="36" :src="require('@/assets/avatar.gif')" v-else></el-avatar>
            </el-container>
          </el-col>
        </el-row>
      </el-header>

      <el-main>

        <router-view v-if="isRouterAlive"/>

      </el-main>

    </el-container>

  </el-container>

  <el-dialog title="个人资料" v-model="changeProfile.show"  width="600px" :close-on-click-modal="false">
    <el-form :model="changeProfile" :rules="changeProfileRules" ref="changeProfile"  label-width="120px">
      <el-form-item prop="avatar" label="头像">
        <el-input v-model="changeProfile.avatar" auto-complete="off" placeholder="请输入头像"></el-input>
      </el-form-item>
      <el-form-item prop="nickname" label="昵称">
        <el-input v-model="changeProfile.nickname" auto-complete="off" placeholder="请输入昵称"></el-input>
      </el-form-item>
      <el-form-item prop="description" label="签名">
        <el-input v-model="changeProfile.description" auto-complete="off" placeholder="请输入签名"></el-input>
      </el-form-item>
      <el-form-item prop="email" label="Email">
        <el-input v-model="changeProfile.email" auto-complete="off" placeholder="请确认邮箱"></el-input>
      </el-form-item>

      <el-form-item prop="github" label="Github">
        <el-input v-model="changeProfile.github" auto-complete="off" placeholder="请输入 Github 账号"></el-input>
      </el-form-item>
      <el-form-item prop="weibo" label="Weibo">
        <el-input v-model="changeProfile.weibo" auto-complete="off" placeholder="请输入 Weibo 账号"></el-input>
      </el-form-item>
      <el-form-item prop="google" label="Google">
        <el-input v-model="changeProfile.google" auto-complete="off" placeholder="请输入 Google 账号"></el-input>
      </el-form-item>
      <el-form-item prop="twitter" label="Twitter">
        <el-input v-model="changeProfile.twitter" auto-complete="off" placeholder="请输入 Twitter 账号"></el-input>
      </el-form-item>
      <el-form-item prop="facebook" label="Facebook">
        <el-input v-model="changeProfile.facebook" auto-complete="off" placeholder="请输入 Facebook 账号"></el-input>
      </el-form-item>
      <el-form-item prop="stackOverflow" label="StackOverflow">
        <el-input v-model="changeProfile.stackOverflow" auto-complete="off" placeholder="请输入 StackOverflow 账号"></el-input>
      </el-form-item>
      <el-form-item prop="youtube" label="Youtube">
        <el-input v-model="changeProfile.youtube" auto-complete="off" placeholder="请输入 Youtube 账号"></el-input>
      </el-form-item>
      <el-form-item prop="instagram" label="Instagram">
        <el-input v-model="changeProfile.instagram" auto-complete="off" placeholder="请输入 Instagram 账号"></el-input>
      </el-form-item>
      <el-form-item prop="skype" label="Skype">
        <el-input v-model="changeProfile.skype" auto-complete="off" placeholder="请输入 Skype 账号"></el-input>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onChangeProfile()">更新</el-button>
        <el-button @click="changeProfile.show = false">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>

  <el-dialog title="修改密码" v-model="changePassword.show"  width="600px" :close-on-click-modal="false">
    <el-form :model="changePassword" :rules="changePasswordRules" ref="changePassword"  label-width="120px">
      <el-form-item prop="oldPassword" label="原始密码">
        <el-input type="password" v-model="changePassword.oldPassword" auto-complete="off" placeholder="请输入原始密码"></el-input>
      </el-form-item>
      <el-form-item prop="newPassword" label="新的密码">
        <el-input type="password" v-model="changePassword.newPassword" auto-complete="off" placeholder="请输入新密码"></el-input>
      </el-form-item>
      <el-form-item prop="confirmPassword" label="确认密码">
        <el-input type="password" v-model="changePassword.confirmPassword" auto-complete="off" placeholder="请确认新密码"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onChangePassword()">修改密码</el-button>
        <el-button @click="changePassword.show = false">取消修改</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>

</template>

<script lang="ts">
import { defineComponent } from 'vue';
import {queryProfile, updateProfile, updatePassword, logout} from "@/api"
import screenfull from 'screenfull'

export default defineComponent({
  name: 'Main',
  mounted() {
    this.queryProfile()
  },
  data() {
    const changePassword = {
      show: false,
      oldPassword: '',
      newPassword: '',
      confirmPassword: '',
    }
    const validateConfirmPassword = (rule: any, value: string, callback: any) => {
      if (!value || value === '') {
        callback(new Error('请再次输入密码'));
      } else if (value !== changePassword.newPassword) {
        callback(new Error('两次输入密码不一致!'));
      } else {
        callback();
      }
    };
    return {
      isCollapse: false,
      nowRoute: '',
      user: {
        avatar: '',
        nickname: '',
      },
      isRouterAlive: true,

      changeProfile: {
        show: false,

        avatar: '',
        nickname: '',
        description: '',
        email: '',

        github: '',
        weibo: '',
        google: '',
        twitter: '',
        facebook: '',
        stackOverflow: '',
        youtube: '',
        instagram: '',
        skype: '',
      },
      changeProfileRules: {
        nickname: [
          {required: true, message: '昵称不能为空', trigger: 'blur'},
        ],
        email: [
          {required: true, message: '邮箱不能为空', trigger: 'blur'},
        ],
      },

      changePassword,
      changePasswordRules: {
        oldPassword: [
          {required: true, message: '原始密码不能为空', trigger: 'blur'},
        ],
        newPassword: [
          {required: true, message: '新密码不能为空', trigger: 'blur'},
          {min: 6, max: 16, message: '新密码长度不合法', trigger: 'blur'},
        ],
        confirmPassword: [
          {required: true, message: '请再次输入密码', trigger: 'blur'},
          { validator: validateConfirmPassword, trigger: 'blur' }
        ],
      },
    };
  },
  methods: {
    onToggleMenu() {
      this.isCollapse = !this.isCollapse;
    },
    onFullScreen() {
      if (!screenfull.isEnabled) {
        this.$message("您的浏览器不能全屏")
        return
      }
      screenfull.toggle()
    },
    onRefresh() {
      this.isRouterAlive = false
      this.$nextTick(() => {
        this.isRouterAlive = true
      });
    },
    onOption(command: string) {
      switch (command) {
        case "profile":
          this.onShowChangeProfile()
          break
        case "password":
          this.onShowChangePassword()
          break
        case "logout":
          this.onLogout()
          break
      }
    },
    queryProfile() {
      queryProfile()
          .then(data => {
            data = data.data
            this.user = data
          })
    },
    onShowChangeProfile() {
      this.changeProfile.show = true

      this.changeProfile.avatar = ''
      this.changeProfile.nickname = ''
      this.changeProfile.description = ''
      this.changeProfile.email = ''

      this.changeProfile.github = ''
      this.changeProfile.weibo = ''
      this.changeProfile.google = ''
      this.changeProfile.twitter = ''
      this.changeProfile.facebook = ''
      this.changeProfile.stackOverflow = ''
      this.changeProfile.youtube = ''
      this.changeProfile.instagram = ''
      this.changeProfile.skype = ''
      queryProfile()
          .then(data => {
            data = data.data

            this.changeProfile.avatar = data.avatar
            this.changeProfile.nickname = data.nickname
            this.changeProfile.description = data.description
            this.changeProfile.email = data.email

            this.changeProfile.github = data.accounts.github
            this.changeProfile.weibo = data.accounts.weibo
            this.changeProfile.google = data.accounts.google
            this.changeProfile.twitter = data.accounts.twitter
            this.changeProfile.facebook = data.accounts.facebook
            this.changeProfile.stackOverflow = data.accounts.stackOverflow
            this.changeProfile.youtube = data.accounts.youtube
            this.changeProfile.instagram = data.accounts.instagram
            this.changeProfile.skype = data.accounts.skype
          })
    },
    onChangeProfile() {
      let from: any = this.$refs['changeProfile'];
      from.validate((valid: boolean) => {
        if (!valid) return
        updateProfile(
            this.changeProfile.avatar,
            this.changeProfile.nickname,
            this.changeProfile.description,
            this.changeProfile.email,

            this.changeProfile.github,
            this.changeProfile.weibo,
            this.changeProfile.google,
            this.changeProfile.twitter,
            this.changeProfile.facebook,
            this.changeProfile.stackOverflow,
            this.changeProfile.youtube,
            this.changeProfile.instagram,
            this.changeProfile.skype,
        )
            .then(_ => {
              this.$message.success("个人资料更新成功")
              this.changeProfile.show = false
              this.queryProfile()
            })
      })
    },
    onShowChangePassword() {
      this.changePassword.show = true
      this.changePassword.oldPassword = ''
      this.changePassword.newPassword = ''
      this.changePassword.confirmPassword = ''
    },
    onChangePassword() {
      let from: any = this.$refs['changePassword'];
      from.validate((valid: boolean) => {
        if (!valid) return
        updatePassword(this.changePassword.oldPassword, this.changePassword.newPassword)
            .then(_ => {
              this.$message.success("密码修改成功")
              this.changePassword.show = false
            })
      })
    },
    onLogout () {
      this.$confirm('确认注销?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        logout()
            .then(() => {
              localStorage.setItem("isLogin", String(false));
              this.$message.success("注销成功")
              this.$router.replace('/login')
            })
      });
    },
  }
});
</script>

<style lang="less">
.el-menu:not(.el-menu--collapse) {
  width: 200px;
}
</style>
