<template>
  <el-container
    class="full-page"
    direction="vertical"
  >
    <el-space wrap>
      <el-input placeholder="请输入名称" v-model="state.filter.name" clearable />
      <el-button type="primary" plain icon="Search" @click="onFilterSearch">搜索</el-button>
      <el-button type="info" plain icon="Delete" @click="onFilterClear">清空</el-button>
      <el-button type="primary" @click="onAddTag">新增标签</el-button>
    </el-space>

    <el-divider />

    <el-table
      class="full-page"
      stripe
      :data="state.data.list"
      @sort-change="onSortChange"
    >
      <el-table-column prop="id" label="ID" width="70" fixed sortable="custom"></el-table-column>
      <el-table-column
        prop="name"
        label="名称"
        min-width="100"
        fixed
        sortable="custom"
      ></el-table-column>
      <el-table-column
        prop="articleCount"
        label="文章数量"
        width="110"
        sortable="custom"
      ></el-table-column>
      <date-time-column prop="createTime" label="创建时间" width="170" format="YYYY-MM-DD HH:mm:ss" ></date-time-column>
      <el-table-column label="操作" width="160">
        <template #default="scope">
          <el-button plain size="small" @click="onEditTag(scope.row)">编辑</el-button>
          <el-button plain type="danger" size="small" @click="onDeleteTag(scope.row)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :page-count="state.data.pages"
      @current-change="onCurrentPageChange"
      style="text-align: center; margin-top: 20px"
      background
      layout="prev, pager, next"
    />
  </el-container>

  <el-dialog
    :title="(editTagDialogState.formModel.id != 0 ? '新增' : '编辑') + '标签'"
    v-model="editTagDialogState.isShow"
    :close-on-click-modal="false"
  >
    <el-form ref="editTagDialogStateFormRef" :model="editTagDialogState.formModel" :rules="editTagDialogState.formRules" label-width="120px">
      <el-form-item label="名称" prop="name">
        <el-input v-model="editTagDialogState.formModel.name" placeholder="请输入名称"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onDialogSubmit">确定</el-button>
        <el-button @click="editTagDialogState.hide()">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ListState, DialogState } from '@/bean'
import { addTag, deleteTag, updateTag, queryTag } from '@/api'

interface TagFilter {
  name: string|null
}

interface EditTagState {
  id: number
  name: string
}

const state = ref(new ListState<TagFilter>({} as TagFilter))

const editTagDialogStateFormRef = ref()
const editTagDialogState = ref(new DialogState<EditTagState>(
  editTagDialogStateFormRef,
  {} as EditTagState,
  {
    name: [
      { required: true, message: '名称不能为空', trigger: 'blur' }
    ],
  },
))

function onRefresh() {
  queryTag(
    state.value.filter.name,
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
function onAddTag() {
  editTagDialogState.value.reset()
  editTagDialogState.value.show()
}
function onEditTag(row: any) {
  editTagDialogState.value.formModel = {
    id: row.id,
    name: row.name
  }
  editTagDialogState.value.show()
}
function onDeleteTag(row: any) {
  ElMessageBox.confirm('确认删除?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteTag(row.id).then(() => {
      ElMessage.success('删除成功')
      onRefresh()
    })
  })
}
function onDialogSubmit() {
  if (editTagDialogState.value.formRef == undefined) return
  editTagDialogState.value.formRef.validate((valid: boolean) => {
    if (!valid) return
    let dialogData = editTagDialogState.value.formModel
    if (dialogData.id != undefined && dialogData.id != 0) {
      updateTag(dialogData.id, dialogData.name).then(() => {
        ElMessage.success('更新成功')
        editTagDialogState.value.hide()
        onRefresh()
      })
    } else {
      addTag(dialogData.name).then(() => {
        ElMessage.success('新增成功')
        editTagDialogState.value.hide()
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
