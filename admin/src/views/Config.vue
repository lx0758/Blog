<template>
  <el-container class="full-page" direction="vertical">
    <el-space wrap>
      <el-input placeholder="请输入键" v-model="state.filter.key" clearable />
      <el-input placeholder="请输入值" v-model="state.filter.value" clearable />
      <el-input placeholder="请输入描述" v-model="state.filter.description" clearable />
      <el-button type="primary" plain icon="Search" @click="onFilterSearch">搜索</el-button>
      <el-button type="info" plain icon="Delete" @click="onFilterClear">清空</el-button>
      <el-button type="primary" @click="onAddConfig">新增配置</el-button>
    </el-space>

    <el-divider />

    <el-table class="full-page" stripe :data="state.data.list" @sort-change="onSortChange">
      <el-table-column prop="key" label="键" width="250" fixed sortable="custom"></el-table-column>
      <el-table-column prop="value" label="值" min-width="300"></el-table-column>
      <el-table-column prop="description" label="描述" min-width="150"></el-table-column>
      <date-time-column prop="createTime" label="创建时间" width="170" format="YYYY-MM-DD HH:mm:ss"></date-time-column>
      <date-time-column prop="updateTime" label="更新时间" width="170" format="YYYY-MM-DD HH:mm:ss"></date-time-column>
      <el-table-column label="操作" width="160">
        <template #default="scope">
          <el-button plain size="small" @click="onEditConfig(scope.row)">编辑</el-button>
          <el-button plain type="danger" size="small" @click="onDeleteConfig(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination :page-count="state.data.size" @current-change="onCurrentPageChange"
      style="justify-content: center; margin-top: 20px" background layout="prev, pager, next" />
  </el-container>

  <el-dialog :title="(editConfigDialogState.formModel.key == undefined ? '新增' : '编辑') + '配置'"
    v-model="editConfigDialogState.isShow" :close-on-click-modal="false">
    <el-form ref="editConfigDialogStateFormRef" :model="editConfigDialogState.formModel"
      :rules="editConfigDialogState.formRules" label-width="120px">
      <el-form-item label="键" prop="key">
        <el-input v-if="editConfigDialogState.formModel.update" v-model="editConfigDialogState.formModel.key"
          placeholder="请输入配置键" disabled></el-input>
        <el-input v-else v-model="editConfigDialogState.formModel.key" placeholder="请输入配置键"></el-input>
      </el-form-item>
      <el-form-item label="值" prop="value">
        <el-input v-model="editConfigDialogState.formModel.value" placeholder="请输入配置值"></el-input>
      </el-form-item>
      <el-form-item label="描述" prop="description">
        <el-input v-model="editConfigDialogState.formModel.description" placeholder="请输入描述"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onDialogSubmit">确定</el-button>
        <el-button @click="editConfigDialogState.hide()">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ListState, DialogState } from '@/bean'
import { addConfig, deleteConfig, updateConfig, queryConfig } from '@/api'

interface ConfigFilter {
  key: string | null
  value: string | null
  description: string | null
}

interface EditConfigState {
  update: boolean
  key: string
  value: string
  description: string
}

const state = ref(new ListState<ConfigFilter>({} as ConfigFilter))

const editConfigDialogStateFormRef = ref()
const editConfigDialogState = ref(new DialogState<EditConfigState>(
  editConfigDialogStateFormRef,
  {} as EditConfigState,
  {
    key: [
      { required: true, message: '键不能为空', trigger: 'blur' }
    ],
    description: [
      { required: true, message: '描述不能为空', trigger: 'blur' }
    ],
  },
))

function onRefresh() {
  queryConfig(
    state.value.filter.key,
    state.value.filter.value,
    state.value.filter.description,
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

function onAddConfig() {
  editConfigDialogState.value.reset()
  editConfigDialogState.value.show()
}

function onEditConfig(row: any) {
  editConfigDialogState.value.formModel = {
    update: true,
    key: row.key,
    value: row.value,
    description: row.description
  }
  editConfigDialogState.value.show()
}

function onDeleteConfig(row: any) {
  ElMessageBox.confirm('确认删除?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteConfig(row.key).then(() => {
      ElMessage.success('删除成功')
      onRefresh()
    })
  })
}

function onDialogSubmit() {
  if (editConfigDialogState.value.formRef == undefined) return
  editConfigDialogState.value.formRef.validate((valid: boolean) => {
    if (!valid) return
    let dialogData = editConfigDialogState.value.formModel
    if (dialogData.update) {
      updateConfig(dialogData.key, dialogData.value, dialogData.description).then(() => {
        ElMessage.success('更新成功')
        editConfigDialogState.value.hide()
        onRefresh()
      })
    } else {
      addConfig(dialogData.key, dialogData.value, dialogData.description).then(() => {
        ElMessage.success('新增成功')
        editConfigDialogState.value.hide()
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
