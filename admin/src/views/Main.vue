<template>
  <el-container style="height: 100%">
    <el-aside unique-opened="true" style="width: auto">
      <el-menu :default-active="state.defaultActive" style="height: 100%" unique-opened router :collapse="state.menuIsCollapse">
        <el-menu-item index="/dashboard">
          <template #title>
            <el-icon>
              <Odometer />
            </el-icon>
            <span>仪表盘</span>
          </template>
        </el-menu-item>
        <el-sub-menu index="/content">
          <template #title>
            <el-icon>
              <MessageBox />
            </el-icon>
            <span>内容</span>
          </template>
          <el-menu-item index="/article">
            <el-icon>
              <Document />
            </el-icon>
            <span>文章</span>
          </el-menu-item>
          <el-menu-item index="/fragment">
            <el-icon>
              <DocumentCopy />
            </el-icon>
            <span>片段</span>
          </el-menu-item>
          <el-menu-item index="/comment">
            <el-icon>
              <ChatDotSquare />
            </el-icon>
            <span>评论</span>
          </el-menu-item>
          <el-menu-item index="/category">
            <el-icon>
              <Filter />
            </el-icon>
            <span>分类</span>
          </el-menu-item>
          <el-menu-item index="/tag">
            <el-icon>
              <PriceTag />
            </el-icon>
            <span>标签</span>
          </el-menu-item>
        </el-sub-menu>
        <el-sub-menu index="/other">
          <template #title>
            <el-icon>
              <Notification />
            </el-icon>
            <span>其他</span>
          </template>
          <el-menu-item index="/file">
            <el-icon>
              <Files />
            </el-icon>
            <span>文件</span>
          </el-menu-item>
          <el-menu-item index="/link">
            <el-icon>
              <Link />
            </el-icon>
            <span>友链</span>
          </el-menu-item>
          <el-menu-item index="/url">
            <el-icon>
              <Connection />
            </el-icon>
            <span>短链</span>
          </el-menu-item>
        </el-sub-menu>
        <el-sub-menu index="/setting">
          <template #title>
            <el-icon>
              <Setting />
            </el-icon>
            <span>设置</span>
          </template>
          <el-menu-item index="/config">
            <el-icon>
              <Operation />
            </el-icon>
            <span>通用</span>
          </el-menu-item>
          <el-menu-item index="/smtp">
            <el-icon>
              <Message />
            </el-icon>
            <span>邮件</span>
          </el-menu-item>
        </el-sub-menu>
      </el-menu>
    </el-aside>

    <el-container>
      <el-header style="box-shadow: 0 -8px 20px 2px #dedee3; height: 50px">
        <el-row type="flex" align="middle" style="height: 50px">
          <el-col :span="16">
            <el-container style="align-items: center">
              <div @click="onToggleMenu()" style="padding: 10px 10px 10px 0; font-size: 22px">
                <el-icon v-if="state.menuIsCollapse">
                  <Expand />
                </el-icon>
                <el-icon v-if="!state.menuIsCollapse">
                  <Fold />
                </el-icon>
              </div>
              <el-breadcrumb separator="/">
                <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
                <el-breadcrumb-item>{{ $route.name }}</el-breadcrumb-item>
              </el-breadcrumb>
            </el-container>
          </el-col>
          <el-col :span="8">
            <el-container style="align-items: center; float: right">
              <el-icon @click="onFullScreen()" style="padding: 10px; font-size: 20px">
                <FullScreen />
              </el-icon>
              <el-icon @click="onRefreshView()" style="padding: 10px; font-size: 20px">
                <Refresh />
              </el-icon>
              <el-dropdown @command="onOption" style="padding: 10px">
                <span class="el-dropdown-link">
                  {{ state.nickname }}
                  <el-icon>
                    <ArrowDownBold />
                  </el-icon>
                </span>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item icon="User" command="profile">个人设置</el-dropdown-item>
                    <el-dropdown-item icon="Lock" command="password">修改密码</el-dropdown-item>
                    <el-dropdown-item icon="SwitchButton" command="logout" divided>注销登录</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
              <el-avatar :size="36" :src="state.avatar"
                v-if="state.avatar != null && state.avatar.length > 0"></el-avatar>
              <el-avatar :size="36" :src="avatar_gif" v-else></el-avatar>
            </el-container>
          </el-col>
        </el-row>
      </el-header>

      <el-main>
        <router-view v-slot="{ Component }" v-if="state.isRouterAvailable">
          <component :is="Component" />
        </router-view>
      </el-main>
    </el-container>
  </el-container>

  <el-dialog title="个人资料" width="600px" v-model="changeUserInfoDialogState.isShow" :close-on-click-modal="false">
    <el-form label-width="120px" ref="changeUserInfoDialogStateFormRef" :model="changeUserInfoDialogState.formModel"
      :rules="changeUserInfoDialogState.formRules">
      <el-form-item prop="avatar" label="头像">
        <el-input v-model="changeUserInfoDialogState.formModel.avatar" auto-complete="off" placeholder="请输入头像" />
      </el-form-item>
      <el-form-item prop="nickname" label="昵称">
        <el-input v-model="changeUserInfoDialogState.formModel.nickname" auto-complete="off" placeholder="请输入昵称" />
      </el-form-item>
      <el-form-item prop="description" label="签名">
        <el-input v-model="changeUserInfoDialogState.formModel.description" auto-complete="off" placeholder="请输入签名" />
      </el-form-item>
      <el-form-item prop="email" label="Email">
        <el-input v-model="changeUserInfoDialogState.formModel.email" auto-complete="off" placeholder="请确认邮箱" />
      </el-form-item>

      <el-form-item prop="github" label="Github">
        <el-input v-model="changeUserInfoDialogState.formModel.github" auto-complete="off" placeholder="请输入 Github 账号" />
      </el-form-item>
      <el-form-item prop="weibo" label="Weibo">
        <el-input v-model="changeUserInfoDialogState.formModel.weibo" auto-complete="off" placeholder="请输入 Weibo 账号" />
      </el-form-item>
      <el-form-item prop="google" label="Google">
        <el-input v-model="changeUserInfoDialogState.formModel.google" auto-complete="off" placeholder="请输入 Google 账号" />
      </el-form-item>
      <el-form-item prop="twitter" label="Twitter">
        <el-input v-model="changeUserInfoDialogState.formModel.twitter" auto-complete="off"
          placeholder="请输入 Twitter 账号" />
      </el-form-item>
      <el-form-item prop="facebook" label="Facebook">
        <el-input v-model="changeUserInfoDialogState.formModel.facebook" auto-complete="off"
          placeholder="请输入 Facebook 账号" />
      </el-form-item>
      <el-form-item prop="stackOverflow" label="StackOverflow">
        <el-input v-model="changeUserInfoDialogState.formModel.stackOverflow" auto-complete="off"
          placeholder="请输入 StackOverflow 账号" />
      </el-form-item>
      <el-form-item prop="youtube" label="Youtube">
        <el-input v-model="changeUserInfoDialogState.formModel.youtube" auto-complete="off"
          placeholder="请输入 Youtube 账号" />
      </el-form-item>
      <el-form-item prop="instagram" label="Instagram">
        <el-input v-model="changeUserInfoDialogState.formModel.instagram" auto-complete="off"
          placeholder="请输入 Instagram 账号" />
      </el-form-item>
      <el-form-item prop="skype" label="Skype">
        <el-input v-model="changeUserInfoDialogState.formModel.skype" auto-complete="off" placeholder="请输入 Skype 账号" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onChangeProfile()">更新</el-button>
        <el-button @click="changeUserInfoDialogState.hide()">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>

  <el-dialog title="修改密码" width="600px" v-model="changePasswordDialogState.isShow" :close-on-click-modal="false">
    <el-form label-width="120px" ref="changePasswordDialogStateFormRef" :model="changePasswordDialogState.formModel"
      :rules="changePasswordDialogState.formRules">
      <el-form-item prop="oldPassword" label="原始密码">
        <el-input type="password" v-model="changePasswordDialogState.formModel.oldPassword" auto-complete="off"
          placeholder="请输入原始密码" />
      </el-form-item>
      <el-form-item prop="newPassword" label="新的密码">
        <el-input type="password" v-model="changePasswordDialogState.formModel.newPassword" auto-complete="off"
          placeholder="请输入新密码" />
      </el-form-item>
      <el-form-item prop="confirmPassword" label="确认密码">
        <el-input type="password" v-model="changePasswordDialogState.formModel.confirmPassword" auto-complete="off"
          placeholder="请确认新密码" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onChangePassword()">修改密码</el-button>
        <el-button @click="changePasswordDialogState.hide()">取消修改</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup lang="ts">
import screenfull from 'screenfull'
import {ref, nextTick, onMounted, watch} from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import avatar_gif from '@/assets/avatar.gif'
import { useUserStore } from '@/stores'
import { DialogState } from '@/bean'
import { queryProfile, updateProfile, updatePassword, logout } from '@/api'

interface State {
  defaultActive: string
  menuIsCollapse: boolean
  isRouterAvailable: boolean

  avatar: string | null
  nickname: string | null
}

interface ChangeUserInfoState {
  avatar: string
  nickname: string
  description: string
  email: string

  github: string
  weibo: string
  google: string
  twitter: string
  facebook: string
  stackOverflow: string
  youtube: string
  instagram: string
  skype: string
}

interface ChangePasswordState {
  oldPassword: string
  newPassword: string
  confirmPassword: string
}

const router = useRouter()

const state = ref<State>({
  defaultActive: '',
  menuIsCollapse: false,
  isRouterAvailable: true,

  avatar: '',
  nickname: '',
})

const validateConfirmPassword = (rule: any, value: string, callback: any) => {
  if (!value || value === '') {
    callback(new Error('请再次输入密码'))
  } else if (value !== changePasswordDialogState.value.formModel.newPassword) {
    callback(new Error('两次输入密码不一致!'))
  } else {
    callback()
  }
}

const changeUserInfoDialogStateFormRef = ref()
const changeUserInfoDialogState = ref(new DialogState<ChangeUserInfoState>(
  changeUserInfoDialogStateFormRef,
  {
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
  {
    nickname: [
      { required: true, message: '昵称不能为空', trigger: 'blur' }
    ],
    email: [
      { required: true, message: '邮箱不能为空', trigger: 'blur' }
    ],
  },
))

const changePasswordDialogStateFormRef = ref()
const changePasswordDialogState = ref(new DialogState<ChangePasswordState>(
  changePasswordDialogStateFormRef,
  {
    oldPassword: '',
    newPassword: '',
    confirmPassword: '',
  }, {
  oldPassword: [
    { required: true, message: '原始密码不能为空', trigger: 'blur' }
  ],
  newPassword: [
    { required: true, message: '新密码不能为空', trigger: 'blur' },
    { min: 6, max: 16, message: '新密码长度不合法', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: '请再次输入密码', trigger: 'blur' },
    { validator: validateConfirmPassword, trigger: 'blur' }
  ],
},
))

function onToggleMenu() {
  state.value.menuIsCollapse = !state.value.menuIsCollapse
}

function onFullScreen() {
  if (!screenfull.isEnabled) {
    ElMessage.error('您的浏览器不能全屏')
    return
  }
  screenfull.toggle()
}

function onRefreshView() {
  state.value.isRouterAvailable = false
  nextTick(() => {
    state.value.isRouterAvailable = true
  })
}

function onOption(command: string) {
  switch (command) {
    case 'profile':
      onShowChangeProfile()
      break
    case 'password':
      onShowChangePassword()
      break
    case 'logout':
      onLogout()
      break
  }
}

function onQueryProfile() {
  queryProfile().then((data) => {
    data = data.data
    state.value.avatar = data.avatar
    state.value.nickname = data.nickname
  })
}

function onShowChangeProfile() {
  changeUserInfoDialogState.value.reset()
  changeUserInfoDialogState.value.show()
  queryProfile().then((data) => {
    data = data.data

    let formModel = changeUserInfoDialogState.value.formModel
    formModel.avatar = data.avatar
    formModel.nickname = data.nickname
    formModel.description = data.description
    formModel.email = data.email

    formModel.github = data.accounts.github
    formModel.weibo = data.accounts.weibo
    formModel.google = data.accounts.google
    formModel.twitter = data.accounts.twitter
    formModel.facebook = data.accounts.facebook
    formModel.stackOverflow = data.accounts.stackOverflow
    formModel.youtube = data.accounts.youtube
    formModel.instagram = data.accounts.instagram
    formModel.skype = data.accounts.skype
  })
}

function onChangeProfile() {
  if (changeUserInfoDialogState.value.formRef == undefined) return
  changeUserInfoDialogState.value.formRef.validate((valid: boolean) => {
    if (!valid) return
    updateProfile(
      changeUserInfoDialogState.value.formModel.avatar,
      changeUserInfoDialogState.value.formModel.nickname,
      changeUserInfoDialogState.value.formModel.description,
      changeUserInfoDialogState.value.formModel.email,

      changeUserInfoDialogState.value.formModel.github,
      changeUserInfoDialogState.value.formModel.weibo,
      changeUserInfoDialogState.value.formModel.google,
      changeUserInfoDialogState.value.formModel.twitter,
      changeUserInfoDialogState.value.formModel.facebook,
      changeUserInfoDialogState.value.formModel.stackOverflow,
      changeUserInfoDialogState.value.formModel.youtube,
      changeUserInfoDialogState.value.formModel.instagram,
      changeUserInfoDialogState.value.formModel.skype
    ).then(() => {
      ElMessage.success('个人资料更新成功')
      changeUserInfoDialogState.value.hide()
      onQueryProfile()
    })
  })
}

function onShowChangePassword() {
  changePasswordDialogState.value.reset()
  changePasswordDialogState.value.show()
}

function onChangePassword() {
  if (changePasswordDialogState.value.formRef == undefined) return
  changePasswordDialogState.value.formRef.validate((valid: boolean) => {
    if (!valid) return
    updatePassword(
      changePasswordDialogState.value.formModel.oldPassword,
      changePasswordDialogState.value.formModel.newPassword
    ).then(() => {
      ElMessage.success('密码修改成功')
      changePasswordDialogState.value.hide()
    })
  })
}

function onLogout() {
  ElMessageBox.confirm('确认注销?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    logout().then(() => {
      let userStore = useUserStore()
      userStore.logout()
      ElMessage.success('注销成功')
      router.replace('/login')
    })
  })
}

onMounted(() => {
  onQueryProfile()
})

watch(
    () => router.currentRoute.value.path,
    (toPath) => {
      state.value.defaultActive = toPath
    },
    {immediate: true, deep: true}
)
</script>

<style scoped>
.el-menu:not(.el-menu--collapse) {
  width: 200px;
}
</style>
