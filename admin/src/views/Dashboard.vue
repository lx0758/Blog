<template>
  <el-row :gutter="20" style="margin-top: 20px">
    <el-col :span="4" :offset="0">
      <el-card style="height: 120px; background: rgb(240, 249, 235)">
        <el-row>
          <span style="font-size: 14px; color: rgba(0, 0, 0, 0.45)">文章</span>
        </el-row>
        <el-row>
          <span style="font-size: 32px; color: rgba(0, 0, 0, 1); margin-top: 18px">{{
            state.articleCount
          }}</span>
        </el-row>
      </el-card>
    </el-col>

    <el-col :span="4">
      <el-card style="height: 120px; background: rgb(240, 249, 235)">
        <el-row>
          <span style="font-size: 14px; color: rgba(0, 0, 0, 0.45)">分类</span>
        </el-row>
        <el-row>
          <span style="font-size: 32px; color: rgba(0, 0, 0, 1); margin-top: 18px">{{
            state.categoryCount
          }}</span>
        </el-row>
      </el-card>
    </el-col>

    <el-col :span="4">
      <el-card style="height: 120px; background: rgb(240, 249, 235)">
        <el-row>
          <span style="font-size: 14px; color: rgba(0, 0, 0, 0.45)">标签</span>
        </el-row>
        <el-row>
          <span style="font-size: 32px; color: rgba(0, 0, 0, 1); margin-top: 18px">{{
            state.tagCount
          }}</span>
        </el-row>
      </el-card>
    </el-col>

    <el-col :span="4">
      <el-card style="height: 120px; background: rgb(240, 249, 235)">
        <el-row>
          <span style="font-size: 14px; color: rgba(0, 0, 0, 0.45)">文件</span>
        </el-row>
        <el-row>
          <span style="font-size: 32px; color: rgba(0, 0, 0, 1); margin-top: 18px">{{
            state.uploadCount
          }}</span>
        </el-row>
      </el-card>
    </el-col>

    <el-col :span="4">
      <el-card style="height: 120px; background: rgb(253, 246, 236)">
        <el-row>
          <span style="font-size: 14px; color: rgba(0, 0, 0, 0.45)">评论</span>
        </el-row>
        <el-row>
          <span style="font-size: 32px; color: rgba(0, 0, 0, 1); margin-top: 18px">{{
            state.commentCount
          }}</span>
        </el-row>
      </el-card>
    </el-col>

    <el-col :span="4">
      <el-card style="height: 120px; background: rgb(254, 240, 240)">
        <el-row>
          <span style="font-size: 14px; color: rgba(0, 0, 0, 0.45)">浏览量</span>
        </el-row>
        <el-row>
          <span style="font-size: 32px; color: rgba(0, 0, 0, 1); margin-top: 18px">{{
            state.browseCount
          }}</span>
        </el-row>
      </el-card>
    </el-col>
  </el-row>

  <el-row :gutter="30" style="margin-top: 40px">
    <el-col :span="12">
      <el-card shadow="hover">
        <template #header>
          <span>最新文章</span>
        </template>
        <el-table :data="state.newArticles" :show-header="false" style="width: 100%">
          <el-table-column prop="title" :show-overflow-tooltip="true" />
          <date-time-column prop="createTime" width="70" format="MM-DD" />
        </el-table>
      </el-card>
    </el-col>

    <el-col :span="12">
      <el-card shadow="hover">
        <template #header>
          <span>最新评论</span>
        </template>
        <el-table :data="state.newComments" :show-header="false" style="width: 100%">
          <el-table-column prop="content">
            <template #default="scope">
              <el-tooltip
                  class="box-item"
                  effect="dark"
                  placement="top">
                <template #content>
                  <pre>{{ scope.row.content }}</pre>
                </template>
                <p class="single-line" style="margin-block-start:0;margin-block-end:0;">
                  {{ scope.row.content }}
                </p>
              </el-tooltip>
            </template>
          </el-table-column>
          <el-table-column prop="nickname" :show-overflow-tooltip="true" width="80" />
          <date-time-column prop="createTime" width="70" format="MM-DD" />
        </el-table>
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { queryDashboard } from '@/api'

interface State {
  articleCount: number
  categoryCount: number
  tagCount: number
  uploadCount: number
  commentCount: number
  browseCount: number

  newArticles: Array<any>
  newComments: Array<any>
}

const state = ref<State>({
  articleCount: 0,
  categoryCount: 0,
  tagCount: 0,
  uploadCount: 0,
  commentCount: 0,
  browseCount: 0,

  newArticles: [],
  newComments: [],
})

onMounted(() => {
  queryDashboard().then((data) => {
    data = data.data
    state.value.articleCount = data.articleCount
    state.value.categoryCount = data.categoryCount
    state.value.tagCount = data.tagCount
    state.value.uploadCount = data.uploadCount
    state.value.commentCount = data.commentCount
    state.value.browseCount = data.browseCount
    state.value.newArticles = data.newArticles
    state.value.newComments = data.newComments
  })
})
</script>

<style></style>
