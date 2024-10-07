<template>
  <el-container class="full-page" direction="vertical">
    <el-space wrap>
      <el-input placeholder="请输入键" v-model="state.filter.key" clearable />
      <blog-selector v-model="state.filter.status" v-bind:type="8" />
      <el-button type="primary" plain icon="Search" @click="onFilterSearch">搜索</el-button>
      <el-button type="info" plain icon="Delete" @click="onFilterClear">清空</el-button>
      <el-button type="primary" @click="onAddFragment">新增片段</el-button>
    </el-space>

    <el-divider />

    <el-table class="full-page" stripe :data="state.data.list" @sort-change="onSortChange">
      <el-table-column prop="id" label="ID" width="70" fixed sortable="custom"></el-table-column>
      <el-table-column prop="key" label="键" min-width="200"></el-table-column>
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
          <el-button plain size="small" @click="onEditFragment(scope.row)">编辑</el-button>
          <el-button plain type="danger" size="small" @click="onDeleteFragment(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination :page-count="state.data.size" @current-change="onCurrentPageChange"
      style="justify-content: center; margin-top: 20px" background layout="prev, pager, next" />
  </el-container>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ListState } from '@/bean'
import { queryFragment, deleteFragment } from '@/api'

interface FragmentFilter {
  key: string | null
  status: number | null
}

const router = useRouter()

const state = ref(new ListState<FragmentFilter>({} as FragmentFilter))

function onRefresh() {
  queryFragment(
    state.value.filter.key,
    null,
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

function onAddFragment() {
  router.push({
    path: '/fragment-edit',
    name: '片段编辑'
  })
}

function onEditFragment(row: any) {
  router.push({
    path: '/fragment-edit',
    name: '片段编辑',
    params: {
      id: row.id
    }
  })
}

function onDeleteFragment(row: any) {
  ElMessageBox.confirm('确认删除?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteFragment(row.id).then(() => {
      ElMessage.success('删除成功')
      onRefresh()
    })
  })
}

onMounted(() => {
  onRefresh()
})
</script>

<style></style>
