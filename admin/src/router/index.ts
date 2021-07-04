import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: '主页',
    component: () => import("@/views/Main.vue"),
    redirect: "dashboard",
    children: [
      {
        path: '/dashboard',
        name: '仪表盘',
        component: () => import("@/views/Dashboard.vue"),
      },
      {
        path: '/article',
        name: '文章',
        component: () => import("@/views/Article.vue"),
      },
      {
        path: '/comment',
        name: '评论',
        component: () => import("@/views/Comment.vue"),
      },
      {
        path: '/category',
        name: '分类',
        component: () => import("@/views/Category.vue"),
      },
      {
        path: '/tag',
        name: '标签',
        component: () => import("@/views/Tag.vue"),
      },
      {
        path: '/upload',
        name: '文件',
        component: () => import("@/views/Upload.vue"),
      },
      {
        path: '/link',
        name: '友链',
        component: () => import("@/views/Link.vue"),
      },
      {
        path: '/url',
        name: '短链',
        component: () => import("@/views/Url.vue"),
      },
      {
        path: '/setting',
        name: '网站设置',
        component: () => import("@/views/Setting.vue"),
      },
    ]
  },
  {
    path: '/login',
    name: '登录',
    component: () => import("@/views/Login.vue"),
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  // 不存在的路由会显示空白页, 这里强制回到上一层
  if (to.matched.length === 0) {
    from.name ? next({ name:from.name }) : next("/")
    return
  }
  // 检查本地登录状态
  const isLogin = localStorage.getItem("isLogin") === "true";
  if (isLogin || to.path === "/login")
    next()
  else
    next("/login")
})

export default router
