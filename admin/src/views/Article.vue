<template>
  <el-container direction="vertical">

    <el-space wrap>
      <el-input
            placeholder="请输入标题"
            v-model="filterTitle"
            clearable/>
      <el-select v-model="filterCategory" placeholder="分类">
        <el-option
            v-for="item in filterCategoryOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value">
        </el-option>
      </el-select>
      <el-select v-model="filterFormat" placeholder="格式">
        <el-option
            v-for="item in filterFormatOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value">
        </el-option>
      </el-select>
      <el-select v-model="filterComment" placeholder="评论">
        <el-option
            v-for="item in filterCommentOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value">
        </el-option>
      </el-select>
      <el-select v-model="filterStatus" placeholder="状态">
        <el-option
            v-for="item in filterStatusOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value">
        </el-option>
      </el-select>
      <el-button type="primary" plain icon="el-icon-search" @click="onFilterSearch">搜索</el-button>
      <el-button type="info" plain icon="el-icon-delete" @click="onFilterClear">清空</el-button>
    </el-space>

    <el-divider/>

    <el-table
        :data="data.list"
        border
        stripe
        style="width: 100%; height: auto">
      <el-table-column prop="id" label="ID" width="60" fixed></el-table-column>
      <el-table-column prop="title" label="标题" min-width="250" fixed>
        <template #default="scope">
          <el-link :href="'/article/' + scope.row.id" type="primary" target="_blank">{{ scope.row.title }}</el-link>
        </template>
      </el-table-column>
      <el-table-column prop="categoryName" label="分类" width="100"></el-table-column>
      <el-table-column prop="format" label="格式" width="100">
        <template #default="scope">
          <el-tag
              :type="scope.row.format === 1 ? 'success' : 'danger'"
              disable-transitions
              size="medium">{{ scope.row.format === 1 ? 'Html' : 'Markdown'}}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="url" label="url" width="100">
        <template #default="scope">
          {{ scope.row.url ? scope.row.url : '-'}}
        </template>
      </el-table-column>
      <el-table-column prop="weight" label="权重" width="60"></el-table-column>
      <el-table-column prop="views" label="浏览数" width="70"></el-table-column>
      <el-table-column :formatter="onFormatDate" prop="createTime" label="创建时间" width="160"></el-table-column>
      <el-table-column :formatter="onFormatDate" prop="updateTime" label="更新时间" width="160"></el-table-column>
      <el-table-column prop="enableComment" label="评论" width="70">
        <template #default="scope">
          <el-tag
              :type="scope.row.enableComment ? 'success' : 'danger'"
              disable-transitions
              size="medium">{{ scope.row.enableComment ? '允许' : '禁止'}}
          </el-tag>
        </template>
      </el-table-column>
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
          <el-button plain size="mini" @click="onEditArticle(scope.row)">编辑</el-button>
          <el-button plain type="danger" size="mini" @click="onDeleteArticle(scope.row)">删除</el-button>
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
import {queryArticle, deleteArticle, queryCategoryOptions} from "@/api"

export default defineComponent({
  name: 'Article',
  components: {},
  mounted() {
    this.onRefresh()
    this.onQueryCategoryOptions()
  },
  data() {
    return {
      filterTitle: null,
      filterCategory: null,
      filterCategoryOptions: [],
      filterFormat: null,
      filterFormatOptions: [
        {
          value: 0,
          label: 'Markdown',
        },
        {
          value: 1,
          label: 'Html',
        },
      ],
      filterComment: null,
      filterCommentOptions: [
        {
          value: 0,
          label: '禁止',
        },
        {
          value: 1,
          label: '允许',
        },
      ],
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
    onQueryCategoryOptions() {
      queryCategoryOptions()
          .then(data => {
            this.filterCategoryOptions = data
          })
    },
    onFilterSearch() {
      this.onRefresh()
    },
    onFilterClear() {
      this.filterTitle = null
      this.filterCategory = null
      this.filterFormat = null
      this.filterComment = null
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
    onEditArticle(row: any) {
      // TODO: 2021-6-28
      console.log("onEditArticle:" + row.title)
    },
    onDeleteArticle(row: any) {
      this.$confirm('确认删除?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteArticle(row.id)
            .then(() => {
              this.$message.success("删除成功");
              this.onRefresh()
            })
      });
    },
    onRefresh() {
      queryArticle(
          this.filterTitle,
          this.filterCategory,
          this.filterFormat,
          this.filterComment,
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