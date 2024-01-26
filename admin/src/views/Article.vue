<template>
  <el-container
    class="full-page"
    direction="vertical"
  >
    <el-space wrap>
      <el-input placeholder="请输入标题" v-model="state.filter.title" clearable />
      <blog-selector v-model="state.filter.categoryId" v-bind:type="1" />
      <el-input placeholder="请输入URL" v-model="state.filter.url" clearable />
      <blog-selector v-model="state.filter.enableComment" v-bind:type="3" />
      <blog-selector v-model="state.filter.status" v-bind:type="4" />
      <el-button type="primary" plain icon="Search" @click="onFilterSearch">搜索</el-button>
      <el-button type="info" plain icon="Delete" @click="onFilterClear">清空</el-button>
      <el-button type="primary" @click="onAddArticle">新增文章</el-button>
    </el-space>

    <el-divider />

    <el-table
      class="full-page"
      stripe
      :data="state.data.list"
      @sort-change="onSortChange"
    >
      <el-table-column prop="id" label="ID" width="70" fixed sortable="custom"></el-table-column>
      <el-table-column prop="title" label="标题" min-width="200" fixed>
        <template #default="scope">
          <el-link v-if="scope.row.url" :href="'/article/' + scope.row.url" type="primary" target="_blank">
            {{scope.row.title}}
          </el-link>
          <el-link v-else :href="'/article/' + scope.row.id" type="primary" target="_blank">
            {{scope.row.title}}
          </el-link>
        </template>
      </el-table-column>
      <el-table-column prop="categoryName" label="分类" width="100"></el-table-column>
      <el-table-column prop="url" label="url" width="200" sortable="custom">
        <template #default="scope">
          <el-link v-if="scope.row.url" :href="'/article/' + scope.row.url" type="primary" target="_blank">
            {{scope.row.url}}</el-link>
          <div v-else>
            {{ scope.row.url ? scope.row.url : '-' }}
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="weight" label="权重" width="80" sortable="custom"></el-table-column>
      <el-table-column prop="views" label="浏览数" width="90" sortable="custom"></el-table-column>
      <date-time-column prop="createTime" label="创建时间" width="170" format="YYYY-MM-DD HH:mm:ss" ></date-time-column>
      <date-time-column prop="updateTime" label="更新时间" width="170" format="YYYY-MM-DD HH:mm:ss" ></date-time-column>
      <el-table-column prop="enableComment" label="评论" width="70">
        <template #default="scope">
          <el-tag
            :type="scope.row.enableComment ? 'success' : 'danger'"
            disable-transitions
            size="small"
            >{{ scope.row.enableComment ? '允许' : '禁止' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="status" label="状态" width="70">
        <template #default="scope">
          <el-tag
            :type="scope.row.status === 1 ? 'success' : 'danger'"
            disable-transitions
            size="small"
            >{{ scope.row.status === 1 ? '发布' : '草稿' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="160">
        <template #default="scope">
          <el-button plain size="small" @click="onEditArticle(scope.row)">编辑</el-button>
          <el-button plain type="danger" size="small" @click="onDeleteArticle(scope.row)">删除</el-button>
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
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ListState } from '@/bean'
import { queryArticle, deleteArticle } from '@/api'

interface ArticleFilter {
  title: string|null
  categoryId: number|null
  url: string|null
  enableComment: boolean|null
  status: number|null
}

const router = useRouter()

const state = ref(new ListState<ArticleFilter>({} as ArticleFilter))

function onRefresh() {
  queryArticle(
    state.value.filter.title,
    state.value.filter.categoryId,
    state.value.filter.url,
    state.value.filter.enableComment,
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
function onAddArticle() {
  router.push({
    path: '/article-edit',
    name: '文章编辑'
  })
}
function onEditArticle(row: any) {
  router.push({
    path: '/article-edit',
    name: '文章编辑',
    params: {
      id: row.id
    }
  })
}
function onDeleteArticle(row: any) {
  ElMessageBox.confirm('确认删除?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteArticle(row.id).then(() => {
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
