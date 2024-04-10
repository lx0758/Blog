<template>
  <el-form label-width="80px" ref="fragment">
    <el-row>
      <el-col :span="8">
        <el-form-item label="片段键名" prop="key">
          <el-input v-model="state.key" placeholder="请输入文章片段" />
        </el-form-item>
      </el-col>
    </el-row>

    <el-container prop="content">
      <limit-height-container bottom="210" min="300">
        <v-md-editor v-model="state.content" :disabled-menus="[]" @save="onEditorSaveFragment"
          @upload-image="onEditorUploadImage" />
      </limit-height-container>
    </el-container>

    <el-row style="margin-top: 20px;">
      <el-col :span="12">
        <el-container>
          <el-form-item label="片段状态" prop="status" style="width: 120px">
            <el-switch v-model="state.status" :active-value="1" :inactive-value="0" />
          </el-form-item>
        </el-container>
      </el-col>
      <el-col :span="12">
        <el-container style="float: right">
          <el-button type="info" plain icon="DocumentDelete" @click="onPageBack">返回</el-button>
          <el-button type="success" plain icon="DocumentChecked" @click="onSaveFragment">保存</el-button>
          <el-button type="primary" plain icon="DocumentAdd" @click="onPublishFragment">发布</el-button>
        </el-container>
      </el-col>
    </el-row>
  </el-form>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useRoute, useRouter, onBeforeRouteLeave } from 'vue-router';
import { ElMessage, ElMessageBox } from 'element-plus'
import { queryFragmentInfo, addFragment, updateFragment, addFile } from '@/api'

interface State {
  id: number
  key: string
  content: string
  status: number
}

const route = useRoute()
const router = useRouter()

const bottom = 250
const height = ref(window.innerHeight - bottom)

const state = ref<State>({
  id: 0,
  status: 0,
} as State)


function onRefreshHeight() {
  height.value = window.innerHeight - bottom
}

function onQueryFragmentInfo() {
  if (state.value.id != undefined && state.value.id != 0) {
    queryFragmentInfo(state.value.id).then((data) => {
      data = data.data
      state.value.key = data.key
      state.value.content = data.content
      state.value.status = data.status
    })
  }
}

function onEditorSaveFragment(text: string, html: string) {
  onSaveFragment()
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

function onSaveFragment() {
  if (!onCheckFragment()) return
  if (state.value.id != undefined && state.value.id != 0) {
    updateFragment(
      state.value.id,
      state.value.key,
      state.value.content,
      state.value.status
    ).then(() => {
      ElMessage.success('更新成功')
    })
  } else {
    addFragment(
      state.value.key,
      state.value.content,
      0
    ).then((data) => {
      data = data.data
      ElMessage.success('保存成功')
      state.value.id = data.id
    })
  }
}

function onPublishFragment() {
  if (!onCheckFragment()) return
  state.value.status = 1
  if (state.value.id != undefined && state.value.id != 0) {
    updateFragment(
      state.value.id,
      state.value.key,
      state.value.content,
      state.value.status,
    ).then(() => {
      ElMessage.success('更新成功')
    })
  } else {
    addFragment(
      state.value.key,
      state.value.content,
      state.value.status,
    ).then((data) => {
      data = data.data
      state.value.id = data.id
      ElMessage.success('发布成功')
    })
  }
}

function onCheckFragment(): boolean {
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
  onQueryFragmentInfo()

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
