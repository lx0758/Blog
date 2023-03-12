<template>

  <el-row :gutter="20" style="margin-top: 20px">

    <el-col :span="4" :offset="0">
      <el-card style="height: 120px; background: rgb(240, 249, 235)">
        <el-row>
          <span style="font-size:14px; color: rgba(0,0,0,.45);">文章</span>
        </el-row>
        <el-row>
          <span style="font-size:32px; color: rgba(0,0,0,1); margin-top: 18px">{{ articleCount }}</span>
        </el-row>
      </el-card>
    </el-col>

    <el-col :span="4">
      <el-card style="height: 120px; background: rgb(240, 249, 235)">
        <el-row>
          <span style="font-size:14px; color: rgba(0,0,0,.45);">分类</span>
        </el-row>
        <el-row>
          <span style="font-size:32px; color: rgba(0,0,0,1); margin-top: 18px">{{ categoryCount }}</span>
        </el-row>
      </el-card>
    </el-col>

    <el-col :span="4">
      <el-card style="height: 120px; background: rgb(240, 249, 235)">
        <el-row>
          <span style="font-size:14px; color: rgba(0,0,0,.45);">标签</span>
        </el-row>
        <el-row>
          <span style="font-size:32px; color: rgba(0,0,0,1); margin-top: 18px">{{ tagCount }}</span>
        </el-row>
      </el-card>
    </el-col>

    <el-col :span="4">
      <el-card style="height: 120px; background: rgb(240, 249, 235)">
        <el-row>
          <span style="font-size:14px; color: rgba(0,0,0,.45);">文件</span>
        </el-row>
        <el-row>
          <span style="font-size:32px; color: rgba(0,0,0,1); margin-top: 18px">{{ uploadCount }}</span>
        </el-row>
      </el-card>
    </el-col>

    <el-col :span="4">
      <el-card style="height: 120px; background: rgb(253, 246, 236)">
        <el-row>
          <span style="font-size:14px; color: rgba(0,0,0,.45);">评论</span>
        </el-row>
        <el-row>
          <span style="font-size:32px; color: rgba(0,0,0,1); margin-top: 18px">{{ commentCount }}</span>
        </el-row>
      </el-card>
    </el-col>

    <el-col :span="4">
      <el-card style="height: 120px; background: rgb(254, 240, 240)">
        <el-row>
          <span style="font-size:14px; color: rgba(0,0,0,.45);">浏览量</span>
        </el-row>
        <el-row>
          <span style="font-size:32px; color: rgba(0,0,0,1); margin-top: 18px">{{ browseCount }}</span>
        </el-row>
      </el-card>
    </el-col>

  </el-row>

  <el-row :gutter="30" style="margin-top: 40px;">

    <el-col :span="8">
      <el-card  shadow="hover">
        <template #header>
          <span>最新文章</span>
        </template>
        <el-table
            :data="newArticles"
            :show-header="false"
            style="width: 100%">
          <el-table-column prop="title" :show-overflow-tooltip="true"/>
          <el-table-column prop="createTime" :formatter="onFormatDate" width="60" align="center"/>
        </el-table>
      </el-card>
    </el-col>

    <el-col :span="8">
      <el-card shadow="hover">
        <template #header>
          <span>最新评论</span>
        </template>
        <el-table
            :data="newComments"
            :show-header="false"
            style="width: 100%">
          <el-table-column prop="content" :show-overflow-tooltip="true"/>
          <el-table-column prop="nickname" :show-overflow-tooltip="true" width="80" align="center"/>
          <el-table-column prop="createTime" :formatter="onFormatDate" width="60" align="center"/>
        </el-table>
      </el-card>
    </el-col>

  </el-row>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { queryDashboard } from '@/api';
import dayjs from "dayjs";

export default defineComponent({
  name: 'Dashboard',
  components: {},
  mounted() {
    queryDashboard()
    .then((data) => {
      data = data.data
      this.articleCount = data.articleCount
      this.categoryCount = data.categoryCount
      this.tagCount = data.tagCount
      this.uploadCount = data.uploadCount
      this.commentCount = data.commentCount
      this.browseCount = data.browseCount
      this.newArticles = data.newArticles
      this.newComments = data.newComments
    })
  },
  data() {
    return {
      articleCount: 0,
      categoryCount: 0,
      tagCount: 0,
      uploadCount: 0,
      commentCount: 0,
      browseCount: 0,

      newArticles: [],
      newComments: [],
    }
  },
  methods: {
    onFormatDate(row: any, column: any) {
      const date = row[column.property];
      if (date == null) return "-"
      return dayjs(date).format("MM-DD");
    },
  }
});
</script>

<style lang="less">
</style>
