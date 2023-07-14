<template>
  <el-container
    class="full-page"
    direction="vertical"
  >
    <el-space wrap>
      <el-input placeholder="请输入文件名" v-model="state.filter.displayName" clearable />
      <el-input placeholder="请输入类型" v-model="state.filter.type" clearable />
      <el-button type="primary" plain icon="Search" @click="onFilterSearch">搜索</el-button>
      <el-button type="info" plain icon="Delete" @click="onFilterClear">清空</el-button>
      <el-button type="primary" @click="onAddUpload">上传文件</el-button>
    </el-space>

    <el-divider />

    <el-table
      class="full-page"
      stripe
      :data="state.data.list"
      @sort-change="onSortChange"
    >
      <el-table-column prop="id" label="ID" width="70" fixed sortable="custom"></el-table-column>
      <el-table-column prop="displayName" label="文件名" width="300" fixed sortable="custom">
        <template #default="scope">
          <el-link :href="scope.row.url" type="primary" target="_blank">{{
            scope.row.displayName
          }}</el-link>
        </template>
      </el-table-column>
      <el-table-column
        :formatter="onFormatFileSize"
        prop="length"
        label="大小"
        width="100"
        sortable="custom"
      ></el-table-column>
      <el-table-column prop="type" label="类型" width="120" sortable="custom"></el-table-column>
      <el-table-column prop="path" label="存储路径" min-width="250" :show-overflow-tooltip="true">
        <template #default="scope">
          <el-link :href="scope.row.url" type="info" target="_blank">{{ scope.row.path }}</el-link>
        </template>
      </el-table-column>
      <el-table-column prop="authorName" label="作者" width="80"></el-table-column>
      <date-time-column prop="createTime" label="创建时间" width="170" format="YYYY-MM-DD HH:mm:ss" ></date-time-column>
      <date-time-column prop="updateTime" label="更新时间" width="170" format="YYYY-MM-DD HH:mm:ss" ></date-time-column>
      <el-table-column label="操作" width="160">
        <template #default="scope">
          <el-button plain size="small" @click="onUpdateUpload(scope.row)">更新</el-button>
          <el-button plain type="danger" size="small" @click="onDeleteUpload(scope.row)"
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
    :title="(editFileDialogState.formModel.id != 0 ? '新增' : '编辑') + '文件'"
    v-model="editFileDialogState.isShow"
    width="500px"
    :close-on-click-modal="false"
  >
    <el-form ref="editFileDialogStateFormRef" :model="editFileDialogState.formModel" label-width="80px">
      <el-form-item label="选择文件">
        <el-upload
          drag
          action=""
          :on-change="onUploadChange"
          :multiple="false"
          :auto-upload="false"
          :show-file-list="false"
        >
          <i class="el-icon-upload"></i>
          <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
        </el-upload>
      </el-form-item>
      <el-form-item label="文件名称">
        <span>{{ onFormatFileName(editFileDialogState.formModel.displayName, '请先选择文件') }}</span>
      </el-form-item>
      <el-form-item label="文件长度">
        <span>{{ onFormatByteSize(editFileDialogState.formModel.length, '请先选择文件') }}</span>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onDialogSubmit">上传</el-button>
        <el-button @click="editFileDialogState.hide()">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ListState, DialogState } from '@/bean'
import { addFile, updateFile, deleteFile, queryFile } from '@/api'

interface FileFilter {
  displayName: string|null
  type: string|null
}

interface EditFileState {
  id: number
  displayName: string
  length: number
  description: string
  file: any
}

const state = ref(new ListState<FileFilter>({} as FileFilter))

const editFileDialogStateFormRef = ref()
const editFileDialogState = ref(new DialogState<EditFileState>(
  editFileDialogStateFormRef,
  {} as EditFileState,
  {},
))

function onRefresh() {
  queryFile(
    state.value.filter.displayName,
    state.value.filter.type,
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
function onFormatFileSize(row: any, column: any) {
  const fileSize = row[column.property]
  return onFormatByteSize(fileSize)
}
function onFormatFileName(name: string, defaultText: string | null = null) {
  if (!name) {
    if (!defaultText) return '-'
    return defaultText
  }
  return name
}
function onFormatByteSize(size: number, defaultText: string | null = null): string {
  if (!size) {
    if (!defaultText) return '-'
    return defaultText
  }
  let temp
  if (size < 1024) {
    return size + ' B'
  } else if (size < 1024 * 1024) {
    temp = size / 1024
    temp = temp.toFixed(2)
    return temp + ' KB'
  } else if (size < 1024 * 1024 * 1024) {
    temp = size / (1024 * 1024)
    temp = temp.toFixed(2)
    return temp + ' MB'
  } else {
    temp = size / (1024 * 1024 * 1024)
    temp = temp.toFixed(2)
    return temp + ' GB'
  }
}
function onAddUpload() {
  editFileDialogState.value.reset()
  editFileDialogState.value.show()
}
function onUpdateUpload(row: any) {
  editFileDialogState.value.formModel = {
    id: row.id,
    displayName: row.displayName,
    length: row.length,
  } as EditFileState
  editFileDialogState.value.show()
}
function onDeleteUpload(row: any) {
  ElMessageBox.confirm('确认删除?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteFile(row.id).then(() => {
      ElMessage.success('删除成功')
      onRefresh()
    })
  })
}
function onUploadChange(file: any, fileList: any) {
  if (fileList.length == 0) return
  let lastFile = fileList[fileList.length - 1]
  editFileDialogState.value.formModel.file = lastFile
  editFileDialogState.value.formModel.displayName = lastFile.name
  editFileDialogState.value.formModel.length = lastFile.size
}
function onDialogSubmit() {
  let dialogData = editFileDialogState.value.formModel
  let file = dialogData.file
  if (!file || !file.raw) {
    ElMessage('文件不能为空')
    return
  }
  if (dialogData.id != undefined && dialogData.id != 0) {
    updateFile(dialogData.id, file.raw).then(() => {
      ElMessage.success('上传成功')
      editFileDialogState.value.hide()
      onRefresh()
    })
  } else {
    addFile(file.raw).then(() => {
      ElMessage.success('上传成功')
      editFileDialogState.value.hide()
      onRefresh()
    })
  }
}

onMounted(() => {
  onRefresh()
})
</script>

<style></style>
