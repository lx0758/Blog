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
import createMermaidPlugin from '@kangc/v-md-editor/lib/plugins/mermaid/cdn'
import '@kangc/v-md-editor/lib/plugins/mermaid/mermaid.css'
import createTodoListPlugin from '@kangc/v-md-editor/lib/plugins/todo-list/index'
import '@kangc/v-md-editor/lib/plugins/todo-list/todo-list.css'
import createLineNumbertPlugin from '@kangc/v-md-editor/lib/plugins/line-number/index'
import createHighlightLinesPlugin from '@kangc/v-md-editor/lib/plugins/highlight-lines/index'
import '@kangc/v-md-editor/lib/plugins/highlight-lines/highlight-lines.css'
import createCopyCodePlugin from '@kangc/v-md-editor/lib/plugins/copy-code/index'
import '@kangc/v-md-editor/lib/plugins/copy-code/copy-code.css'

import BlogSelector from '@/components/BlogSelector.vue'
import DateTimeColumn from '@/components/DateTimeColumn.vue'
import LimitHeightContainer from '@/components/LimitHeightContainer.vue'

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
VueMarkdownEditor.use(createMermaidPlugin())
VueMarkdownEditor.use(createTodoListPlugin())
VueMarkdownEditor.use(createLineNumbertPlugin())
VueMarkdownEditor.use(createHighlightLinesPlugin())
VueMarkdownEditor.use(createCopyCodePlugin())
app.use(VueMarkdownEditor)

app.component('blog-selector', BlogSelector)
app.component('date-time-column', DateTimeColumn)
app.component('limit-height-container', LimitHeightContainer)

app.mount('#app')
