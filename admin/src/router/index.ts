import { createRouter, createMemoryHistory } from 'vue-router'
import { useUserStore } from '@/stores'

const routes = [
  {
    path: '/',
    name: '主页',
    component: () => import('@/views/Main.vue'),
    redirect: 'dashboard',
    children: [
      {
        path: '/dashboard',
        name: '仪表盘',
        component: () => import('@/views/Dashboard.vue')
      },
      {
        path: '/article',
        name: '文章',
        component: () => import('@/views/Article.vue')
      },
      {
        path: '/article-edit/:id*',
        name: '文章编辑',
        component: () => import('@/views/ArticleEditor.vue'),
        meta: {
          notKeepAlive: true
        }
      },
      {
        path: '/fragment',
        name: '片段',
        component: () => import('@/views/Fragment.vue')
      },
      {
        path: '/fragment-edit/:id*',
        name: '片段编辑',
        component: () => import('@/views/FragmentEditor.vue'),
        meta: {
          notKeepAlive: true
        }
      },
      {
        path: '/comment',
        name: '评论',
        component: () => import('@/views/Comment.vue')
      },
      {
        path: '/category',
        name: '分类',
        component: () => import('@/views/Category.vue')
      },
      {
        path: '/tag',
        name: '标签',
        component: () => import('@/views/Tag.vue')
      },
      {
        path: '/file',
        name: '文件',
        component: () => import('@/views/File.vue')
      },
      {
        path: '/link',
        name: '友链',
        component: () => import('@/views/Link.vue')
      },
      {
        path: '/url',
        name: '短链',
        component: () => import('@/views/Url.vue')
      },
      {
        path: '/config',
        name: '网站设置',
        component: () => import('@/views/Config.vue')
      },
      {
        path: '/smtp',
        name: '邮件设置',
        component: () => import('@/views/features/SMTP.vue')
      }
    ]
  },
  {
    path: '/login',
    name: '登录',
    component: () => import('@/views/Login.vue')
  }
]

const router = createRouter({
  history: createMemoryHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  // 不存在的路由会显示空白页, 这里强制回到上一层
  if (to.matched.length === 0) {
    from.name ? next({ name: from.name }) : next('/')
    return
  }
  // 检查本地登录状态
  const loginStatus = useUserStore()
  if (loginStatus.isLoggedIn || to.path === '/login') next()
  else next('/login')
})

export default router
