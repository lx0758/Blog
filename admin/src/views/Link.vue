<template>
  <el-container class="full-page" direction="vertical">
    <el-space wrap>
      <el-input placeholder="请输入标题" v-model="state.filter.title" clearable />
      <el-input placeholder="请输入链接" v-model="state.filter.link" clearable />
      <blog-selector v-model="state.filter.status" v-bind:type="6" />
      <el-button type="primary" plain icon="Search" @click="onFilterSearch">搜索</el-button>
      <el-button type="info" plain icon="Delete" @click="onFilterClear">清空</el-button>
      <el-button type="primary" @click="onAddLink">新增短链</el-button>
    </el-space>

    <el-divider />

    <el-table class="full-page" stripe :data="state.data.list" @sort-change="onSortChange">
      <el-table-column prop="id" label="ID" width="70" fixed sortable="custom"></el-table-column>
      <el-table-column prop="title" label="标题" width="240" fixed sortable="custom">
        <template #default="scope">
          <el-link :href="scope.row.url" type="primary" target="_blank">
            {{ scope.row.title }}
          </el-link>
        </template>
      </el-table-column>
      <el-table-column prop="url" label="链接地址" min-width="250" :show-overflow-tooltip="true">
        <template #default="scope">
          <el-link :href="scope.row.url" type="info" target="_blank">
            {{ scope.row.url }}
          </el-link>
        </template>
      </el-table-column>
      <el-table-column prop="weight" label="权重" width="80" sortable="custom"></el-table-column>
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
          <el-button plain size="small" @click="onEditLink(scope.row)">编辑</el-button>
          <el-button plain type="danger" size="small" @click="onDeleteLink(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination :page-count="state.data.size" @current-change="onCurrentPageChange"
      style="justify-content: center; margin-top: 20px" background layout="prev, pager, next" />
  </el-container>

  <el-dialog :title="(editLinkDialogState.formModel.id != 0 ? '新增' : '编辑') + '友链'" v-model="editLinkDialogState.isShow"
    :close-on-click-modal="false">
    <el-form ref="editLinkDialogStateFormRef" :model="editLinkDialogState.formModel"
      :rules="editLinkDialogState.formRules" label-width="120px">
      <el-form-item label="标题" prop="title">
        <el-input v-model="editLinkDialogState.formModel.title" placeholder="请输入标题"></el-input>
      </el-form-item>
      <el-form-item label="链接" prop="url">
        <el-input v-model="editLinkDialogState.formModel.url" placeholder="请输入链接地址"></el-input>
      </el-form-item>
      <el-form-item label="权重" prop="weight">
        <el-input v-model="editLinkDialogState.formModel.weight" placeholder="请输入链接权重" type="number"></el-input>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <blog-selector v-model="editLinkDialogState.formModel.status" v-bind:type="6" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onDialogSubmit">确定</el-button>
        <el-button @click="editLinkDialogState.hide()">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ListState, DialogState } from '@/bean'
import { addLink, deleteLink, updateLink, queryLink } from '@/api'

interface LinkFilter {
  title: string | null
  link: string | null
  status: number | null
}

interface EditLinkState {
  id: number
  title: string
  url: string
  weight: number
  status: number
}

const state = ref(new ListState<LinkFilter>({} as LinkFilter))

const editLinkDialogStateFormRef = ref()
const editLinkDialogState = ref(new DialogState<EditLinkState>(
  editLinkDialogStateFormRef,
  {} as EditLinkState,
  {
    title: [
      { required: true, message: '标题不能为空', trigger: 'blur' }
    ],
    url: [
      { required: true, message: '链接不能为空', trigger: 'blur' }
    ],
    weight: [
      { required: true, message: '请输入权重', trigger: 'blur' }
    ],
    status: [
      { required: true, message: '请选择状态', trigger: 'blur' }
    ],
  },
))

function onRefresh() {
  queryLink(
    state.value.filter.title,
    state.value.filter.link,
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

function onAddLink() {
  editLinkDialogState.value.reset()
  editLinkDialogState.value.show()
}

function onEditLink(row: any) {
  editLinkDialogState.value.formModel = {
    id: row.id,
    title: row.title,
    url: row.url,
    weight: row.weight,
    status: row.status
  }
  editLinkDialogState.value.show()
}

function onDeleteLink(row: any) {
  ElMessageBox.confirm('确认删除?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteLink(row.id).then(() => {
      ElMessage.success('删除成功')
      onRefresh()
    })
  })
}

function onDialogSubmit() {
  if (editLinkDialogState.value.formRef == undefined) return
  editLinkDialogState.value.formRef.validate((valid: boolean) => {
    if (!valid) return
    let dialogData = editLinkDialogState.value.formModel
    if (dialogData.id != undefined && dialogData.id != 0) {
      updateLink(
        dialogData.id,
        dialogData.title,
        dialogData.url,
        dialogData.weight,
        dialogData.status,
      ).then(() => {
        ElMessage.success('更新成功')
        editLinkDialogState.value.hide()
        onRefresh()
      })
    } else {
      addLink(
        dialogData.title,
        dialogData.url,
        dialogData.weight,
        dialogData.status,
      ).then(() => {
        ElMessage.success('新增成功')
        editLinkDialogState.value.hide()
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
