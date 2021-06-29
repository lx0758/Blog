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
                  {{ user.name }}<i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item icon="el-icon-user" command="profile">个人设置</el-dropdown-item>
                    <el-dropdown-item icon="el-icon-setting" command="website" divided>网站信息</el-dropdown-item>
                    <el-dropdown-item icon="el-icon-delete" command="logout" divided>注销登录</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
              <el-avatar :size="36" :src="user.avatarUrl"></el-avatar>
            </el-container>
          </el-col>
        </el-row>
      </el-header>

      <el-main>

        <router-view v-if="isRouterAlive"/>

      </el-main>

    </el-container>

  </el-container>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import {logout} from "@/api"
import screenfull from 'screenfull'

export default defineComponent({
  name: 'Main',
  mounted() {
    this.nowRoute = this.$route.path.replace('/', '');
  },
  watch: {
    $route(to, from) {
      this.nowRoute = to.path.replace('/', '');
    }
  },
  data() {
    return {
      isCollapse: false,
      nowRoute: '',
      user: {
        name: "Admin",
        avatarUrl: "https://lh3.googleusercontent.com/ogw/ADea4I6BkK_sPAjxNWu4DQ_x6V1vSUlkha_GBtdos4nH=s83-c-mo",
      },
      isRouterAlive: true,
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
          this.onProfile()
          break
        case "website":
          this.onWebsite()
          break
        case "logout":
          this.onLogout()
          break
      }
    },
    onProfile() {
      this.$router.replace('/profile')
    },
    onWebsite() {
      this.$router.replace('/website')
    },
    onLogout: function () {
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
