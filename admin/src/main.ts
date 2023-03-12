import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import ElementPlus from 'element-plus'
import zh from 'element-plus/es/locale/lang/zh-cn'
import 'element-plus/lib/theme-chalk/index.css'

import VueMarkdownEditor from '@kangc/v-md-editor';
import '@kangc/v-md-editor/lib/style/base-editor.css';
import vuepressTheme from '@kangc/v-md-editor/lib/theme/vuepress.js';
import '@kangc/v-md-editor/lib/theme/style/vuepress.css';
import Prism from 'prismjs';
VueMarkdownEditor.use(vuepressTheme, {
    Prism,
});

createApp(App)
    .use(store)
    .use(router)
    .use(ElementPlus, {
        locale: zh, size: 'default'
    })
    .use(VueMarkdownEditor)
    .mount('#app')
