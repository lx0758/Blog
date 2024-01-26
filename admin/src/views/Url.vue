<template>
  <el-container class="full-page" direction="vertical">
    <el-space wrap>
      <el-input placeholder="请输入键" v-model="state.filter.key" clearable />
      <el-input placeholder="请输入链接" v-model="state.filter.url" clearable />
      <el-input placeholder="请输入描述" v-model="state.filter.description" clearable />
      <blog-selector v-model="state.filter.status" v-bind:type="7" />
      <el-button type="primary" plain icon="Search" @click="onFilterSearch">搜索</el-button>
      <el-button type="info" plain icon="Delete" @click="onFilterClear">清空</el-button>
      <el-button type="primary" @click="onAddUrl">新增短链</el-button>
    </el-space>

    <el-divider />

    <el-table class="full-page" stripe :data="state.data.list" @sort-change="onSortChange">
      <el-table-column prop="id" label="ID" width="70" fixed sortable="custom"></el-table-column>
      <el-table-column prop="key" label="主键" width="120" fixed sortable="custom">
        <template #default="scope">
          <el-link :href="'/u/' + scope.row.key" type="primary" target="_blank">
            {{ scope.row.key }}
          </el-link>
        </template>
      </el-table-column>
      <el-table-column prop="url" label="链接" min-width="250" :show-overflow-tooltip="true">
        <template #default="scope">
          <el-link :href="scope.row.url" type="info" target="_blank">
            {{ scope.row.url }}
          </el-link>
        </template>
      </el-table-column>
      <el-table-column prop="description" label="描述" min-width="150"></el-table-column>
      <el-table-column prop="views" label="访问量" width="90" sortable="custom"></el-table-column>
      <el-table-column prop="totalViews" label="总访问量" width="110" sortable="custom"></el-table-column>
      <date-time-column prop="createTime" label="创建时间" width="170" format="YYYY-MM-DD HH:mm:ss"></date-time-column>
      <date-time-column prop="updateTime" label="更新时间" width="170" format="YYYY-MM-DD HH:mm:ss"></date-time-column>
      <el-table-column prop="status" label="状态" width="70">
        <template #default="scope">
          <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'" disable-transitions size="small">
            {{ scope.row.status === 1 ? '启用' : '禁用' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="160">
        <template #default="scope">
          <el-button plain size="small" @click="onEditUrl(scope.row)">编辑</el-button>
          <el-button plain type="danger" size="small" @click="onDeleteUrl(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination :page-count="state.data.pages" @current-change="onCurrentPageChange"
      style="text-align: center; margin-top: 20px" background layout="prev, pager, next" />
  </el-container>

  <el-dialog :title="(editUrlDialogState.formModel.id != 0 ? '新增' : '编辑') + '短链'" v-model="editUrlDialogState.isShow"
    :close-on-click-modal="false">
    <el-form ref="editUrlDialogStateFormRef" :model="editUrlDialogState.formModel" :rules="editUrlDialogState.formRules"
      label-width="120px">
      <el-form-item label="主键" prop="key">
        <el-input v-model="editUrlDialogState.formModel.key" placeholder="请输入键"></el-input>
      </el-form-item>
      <el-form-item label="链接" prop="url">
        <el-input v-model="editUrlDialogState.formModel.url" placeholder="请输入链接"></el-input>
      </el-form-item>
      <el-form-item label="描述" prop="description">
        <el-input v-model="editUrlDialogState.formModel.description" placeholder="请输入描述"></el-input>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <blog-selector v-model="editUrlDialogState.formModel.status" v-bind:type="7" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onDialogSubmit">确定</el-button>
        <el-button @click="editUrlDialogState.hide()">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ListState, DialogState } from '@/bean'
import { addUrl, deleteUrl, updateUrl, queryUrl } from '@/api'

interface UrlFilter {
  key: string | null
  url: string | null
  description: string | null
  status: number | null
}

interface EditUrlState {
  id: number
  key: string
  url: string
  description: string
  status: number
}

const state = ref(new ListState<UrlFilter>({} as UrlFilter))

const editUrlDialogStateFormRef = ref()
const editUrlDialogState = ref(new DialogState<EditUrlState>(
  editUrlDialogStateFormRef,
  {} as EditUrlState,
  {
    key: [
      { required: true, message: '键不能为空', trigger: 'blur' }
    ],
    url: [
      { required: true, message: '链接不能为空', trigger: 'blur' }
    ],
    description: [
      { required: true, message: '描述不能为空', trigger: 'blur' }
    ],
    status: [
      { required: true, message: '请选择状态', trigger: 'blur' }
    ],
  },
))

function onRefresh() {
  queryUrl(
    state.value.filter.key,
    state.value.filter.url,
    state.value.filter.description,
    state.value.filter.status,
    state.value.data.pageNum,
    state.value.data.pageSize,
    state.value.order.name,
    state.value.order.method
  ).then((data) => {
    state.value.data = data.data
  })
}

function onFilterSearch() {
  onRefresh()
}

function onFilterClear() {
  state.value.clearFilter()
  onRefresh()
}

function onSortChange(column: any) {
  state.value.saveOrder(column.prop, column.order)
  onRefresh()
}

function onCurrentPageChange(currentPage: number) {
  const changed = state.value.data.pageNum != currentPage
  state.value.data.pageNum = currentPage
  if (changed) {
    onRefresh()
  }
}

function onAddUrl() {
  editUrlDialogState.value.reset()
  editUrlDialogState.value.show()
}

function onEditUrl(row: any) {
  editUrlDialogState.value.formModel = {
    id: row.id,
    key: row.key,
    url: row.url,
    description: row.description,
    status: row.status
  }
  editUrlDialogState.value.show()
}

function onDeleteUrl(row: any) {
  ElMessageBox.confirm('确认删除?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteUrl(row.id).then(() => {
      ElMessage.success('删除成功')
      onRefresh()
    })
  })
}

function onDialogSubmit() {
  if (editUrlDialogState.value.formRef == undefined) return
  editUrlDialogState.value.formRef.validate((valid: boolean) => {
    if (!valid) return
    let dialogData = editUrlDialogState.value.formModel
    if (dialogData.id != undefined && dialogData.id != 0) {
      updateUrl(
        dialogData.id,
        dialogData.key,
        dialogData.url,
        dialogData.description,
        dialogData.status
      ).then(() => {
        ElMessage.success('更新成功')
        editUrlDialogState.value.hide()
        onRefresh()
      })
    } else {
      addUrl(
        dialogData.key,
        dialogData.url,
        dialogData.description,
        dialogData.status
      ).then(() => {
        ElMessage.success('新增成功')
        editUrlDialogState.value.hide()
        onRefresh()
      })
    }
  })
}

onMounted(() => {
  onRefresh()
})
</script>

<style></style>
