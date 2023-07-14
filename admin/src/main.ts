import { createApp } from 'vue'

import App from './App.vue'

import router from './router'
import store from './stores'

import ElementPlus from 'element-plus'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import 'element-plus/dist/index.css'

import VueMarkdownEditor from '@kangc/v-md-editor'
import '@kangc/v-md-editor/lib/style/base-editor.css'
import vuepressTheme from '@kangc/v-md-editor/lib/theme/vuepress.js'
import '@kangc/v-md-editor/lib/theme/style/vuepress.css'
import Prism from 'prismjs'

import BlogSelector from '@/components/BlogSelector.vue'
import DateTimeColumn from '@/components/DateTimeColumn.vue'
import LimitHightContainer from '@/components/LimitHightContainer.vue'

const app = createApp(App)

app.use(router)

app.use(store)

app.use(ElementPlus)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

VueMarkdownEditor.use(vuepressTheme, {
  Prism
})
app.use(VueMarkdownEditor)

app.component('blog-selector', BlogSelector)
app.component('date-time-column', DateTimeColumn)
app.component('limit-height-container', LimitHightContainer)

app.mount('#app')
