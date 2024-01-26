<template>
  <el-form label-width="80px" ref="article">
    <el-row>
      <el-col :span="8">
        <el-form-item label="文章标题" prop="title">
          <el-input v-model="state.title" placeholder="请输入文章标题" />
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="文章分类" prop="categoryId">
          <blog-selector v-model="state.categoryId" v-bind:type="1" style="width: 100%" />
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="文章权重" prop="weight">
          <el-input v-model="state.weight" placeholder="请输入文章权重" type="number" />
        </el-form-item>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="8">
        <el-form-item label="文章链接" prop="url">
          <el-input v-model="state.url" placeholder="请输入文章链接" />
        </el-form-item>
      </el-col>
      <el-col :span="16">
        <el-form-item label="文章标签" prop="tags">
          <el-select v-model="state.tags" clearable multiple filterable allow-create default-first-option
            placeholder="请输入标签(回车添加)" style="width: 100%">
            <el-option v-for="item in state.tags" :key="item" :label="item" :value="item"></el-option>
          </el-select>
        </el-form-item>
      </el-col>
    </el-row>

    <el-container prop="content">
      <limit-height-container bottom="260" min="300">
        <v-md-editor v-model="state.content" :disabled-menus="[]" @save="onEditorSaveArticle"
          @upload-image="onEditorUploadImage" />
      </limit-height-container>
    </el-container>

    <el-row style="margin-top: 20px;">
      <el-col :span="12">
        <el-container>
          <el-form-item label="评论状态" prop="enableComment" style="width: 120px">
            <el-switch v-model="state.enableComment" :active-value="true" :inactive-value="false" />
          </el-form-item>
          <el-form-item label="文章状态" prop="status" style="width: 120px">
            <el-switch v-model="state.status" :active-value="1" :inactive-value="0" />
          </el-form-item>
        </el-container>
      </el-col>
      <el-col :span="12">
        <el-container style="float: right">
          <el-button type="info" plain icon="DocumentDelete" @click="onPageBack">返回</el-button>
          <el-button type="success" plain icon="DocumentChecked" @click="onSaveArticle">保存</el-button>
          <el-button type="primary" plain icon="DocumentAdd" @click="onPublishArticle">发布</el-button>
        </el-container>
      </el-col>
    </el-row>
  </el-form>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useRoute, useRouter, onBeforeRouteLeave } from 'vue-router';
import { ElMessage, ElMessageBox } from 'element-plus'
import { queryArticleInfo, addArticle, updateArticle, addFile } from '@/api'

interface State {
  id: number
  title: string
  categoryId: number
  url: string
  tags: Array<string>
  content: string
  weight: number
  enableComment: boolean
  status: number
}

const route = useRoute()
const router = useRouter()

const bottom = 300
const height = ref(window.innerHeight - bottom)

const state = ref<State>({
  id: 0,
  enableComment: true,
  status: 0,
} as State)


function onRefreshHeight() {
  height.value = window.innerHeight - bottom
}

function onQueryArticleInfo() {
  if (state.value.id != undefined && state.value.id != 0) {
    queryArticleInfo(state.value.id).then((data) => {
      data = data.data
      state.value.title = data.title
      state.value.categoryId = data.categoryId
      state.value.url = data.url
      state.value.tags = data.tags
      state.value.content = data.content
      state.value.weight = data.weight
      state.value.enableComment = data.enableComment
      state.value.status = data.status
    })
  }
}

function onEditorSaveArticle(text: string, html: string) {
  onSaveArticle()
}

function onEditorUploadImage(event: any, insertImage: any, files: any[]) {
  addFile(files[0]).then((data) => {
    data = data.data
    insertImage({
      url: data.url,
      desc: data.displayName
    })
  })
}

function onPageBack() {
  router.back()
}

function onSaveArticle() {
  if (!onCheckArticle()) return
  if (state.value.id != undefined && state.value.id != 0) {
    updateArticle(
      state.value.id,
      state.value.title,
      state.value.categoryId,
      state.value.url,
      state.value.tags,
      state.value.content,
      state.value.weight,
      state.value.enableComment,
      state.value.status
    ).then(() => {
      ElMessage.success('更新成功')
    })
  } else {
    addArticle(
      state.value.title,
      state.value.categoryId,
      state.value.url,
      state.value.tags,
      state.value.content,
      state.value.weight,
      state.value.enableComment,
      0
    ).then((data) => {
      data = data.data
      ElMessage.success('保存成功')
      state.value.id = data.id
    })
  }
}

function onPublishArticle() {
  if (!onCheckArticle()) return
  state.value.status = 1
  if (state.value.id != undefined && state.value.id != 0) {
    updateArticle(
      state.value.id,
      state.value.title,
      state.value.categoryId,
      state.value.url,
      state.value.tags,
      state.value.content,
      state.value.weight,
      state.value.enableComment,
      state.value.status,
    ).then(() => {
      ElMessage.success('更新成功')
    })
  } else {
    addArticle(
      state.value.title,
      state.value.categoryId,
      state.value.url,
      state.value.tags,
      state.value.content,
      state.value.weight,
      state.value.enableComment,
      state.value.status,
    ).then((data) => {
      data = data.data
      state.value.id = data.id
      ElMessage.success('发布成功')
    })
  }
}

function onCheckArticle(): boolean {
  if (!state.value.title || state.value.title === '') {
    ElMessage.error('标题不能为空')
    return false
  }
  if (!state.value.categoryId || state.value.categoryId === 0) {
    ElMessage.error('分类不能为空')
    return false
  }
  if (!state.value.content || state.value.content === '') {
    ElMessage.error('内容不能为空')
    return false
  }
  return true
}

onMounted(() => {
  let params = route.params
  if (params && params.id) {
    state.value.id = Number.parseInt(params.id as string)
  }
  onQueryArticleInfo()

  window.onresize = () => {
    onRefreshHeight()
  }
  onRefreshHeight()
})
onUnmounted(() => {
  window.onresize = null
})
onBeforeRouteLeave((to, from, next) => {
  ElMessageBox.confirm('系统可能不会保存您所做的更改。', '离开此页面？', {
    confirmButtonText: '离开',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(() => {
      next()
    })
    .catch(() => {
      next(false)
    })
});
</script>

<style scoped></style>
