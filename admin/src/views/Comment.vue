<template>
  <el-container direction="vertical">

    <el-space wrap>
      <el-input
          placeholder="请输入文章ID"
          v-model="filterArticle"
          clearable/>
      <el-input
          placeholder="请输入作者"
          v-model="filterAuthor"
          clearable/>
      <el-input
          placeholder="请输入内容"
          v-model="filterContent"
          clearable/>
      <el-input
          placeholder="请输入邮箱"
          v-model="filterEmail"
          clearable/>
      <el-input
          placeholder="请输入IP地址"
          v-model="filterIP"
          clearable/>
      <el-select v-model="filterStatus" placeholder="状态">
        <el-option
            v-for="item in filterStatusOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value">
        </el-option>
      </el-select>
      <el-button type="primary" plain @click="onFilterSearch">搜索</el-button>
      <el-button type="info" plain @click="onFilterClear">清空</el-button>
    </el-space>

    <el-divider/>

    <el-table
        :data="data.list"
        border
        stripe
        style="width: 100%; height: auto">
      <el-table-column prop="id" label="ID" width="60" fixed></el-table-column>
      <el-table-column prop="articleName" label="文章标题" min-width="200" :show-overflow-tooltip="true" fixed>
        <template #default="scope">
          <el-link :href="'/article/' + scope.row.articleId" type="primary" target="_blank">{{ scope.row.articleName }}</el-link>
        </template>
      </el-table-column>
      <el-table-column prop="nickname" label="昵称" width="120"></el-table-column>
      <el-table-column prop="email" label="邮箱" width="180"></el-table-column>
      <el-table-column prop="url" label="链接" width="200">
        <template #default="scope">
          {{ scope.row.url ? scope.row.url : '-'}}
        </template>
      </el-table-column>
      <el-table-column prop="ip" label="IP" width="140"></el-table-column>
      <el-table-column prop="content" label="内容" min-width="300" :show-overflow-tooltip="true"></el-table-column>
      <el-table-column :formatter="onFormatDate" prop="createTime" label="创建时间" width="160"></el-table-column>
      <el-table-column :formatter="onFormatDate" prop="updateTime" label="更新时间" width="160"></el-table-column>
      <el-table-column prop="status" label="状态" width="70">
        <template #default="scope">
          <el-tag
              :type="scope.row.status === 1 ? 'success' : 'danger'"
              disable-transitions
              size="medium">{{ scope.row.status === 1 ? '发布' : scope.row.status === -1 ? '删除' : '草稿'}}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="160">
        <template #default="scope">
          <el-button plain size="mini" @click="onEditComment(scope.row)">编辑</el-button>
          <el-button plain type="danger" size="mini" @click="onDeleteComment(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
        :page-count="data.pages"
        @current-change="onCurrentPageChange"
        style="text-align: center; margin-top: 20px"
        background
        layout="prev, pager, next"/>

  </el-container>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import dayjs from "dayjs"
import {queryComment, deleteComment} from "@/api"

export default defineComponent({
  name: 'Comment',
  components: {},
  mounted() {
    this.onRefresh()
  },
  data() {
    return {
      filterArticle: null,
      filterAuthor: null,
      filterContent: null,
      filterEmail: null,
      filterIP: null,
      filterStatus: null,
      filterStatusOptions: [
        {
          value: -1,
          label: '删除',
        },
        {
          value: 0,
          label: '草稿',
        },
        {
          value: 1,
          label: '发布',
        },
      ],

      data: {
        pageNum: 1,
        pageSize: 20,
        pages: 0,
        total: 0,
        list: [],
      },
    }
  },
  methods: {
    onFilterSearch() {
      this.onRefresh()
    },
    onFilterClear() {
      this.filterArticle = null
      this.filterAuthor = null
      this.filterContent = null
      this.filterEmail = null
      this.filterIP = null
      this.filterStatus = null
      this.onRefresh()
    },
    onFormatDate(row: any, column: any) {
      const date = row[column.property];
      return dayjs(date).format("YYYY-MM-DD HH:mm:ss");
    },
    onCurrentPageChange(currentPage: number) {
      this.data.pageNum = currentPage;
      this.onRefresh();
    },
    onEditComment(row: any) {
      // TODO: 2021-6-28
      console.log("onEditComment:" + row.title)
    },
    onDeleteComment(row: any) {
      this.$confirm('确认删除?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteComment(row.id)
            .then(() => {
              this.$message.success("删除成功");
              this.onRefresh()
            })
      });
    },
    onRefresh() {
      queryComment(
          this.filterArticle,
          this.filterAuthor,
          this.filterContent,
          this.filterEmail,
          this.filterIP,
          this.filterStatus,
          this.data.pageNum,
          this.data.pageSize,
      )
          .then(data => {
            this.data = data.data;
          })
    },
  }
});
</script>

<style lang="less">
</style>