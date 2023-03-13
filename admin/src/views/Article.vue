<template>
  <el-container direction="vertical">

    <el-space wrap>
      <el-input
            placeholder="请输入标题"
            v-model="filter.title"
            size="small"
            clearable/>
      <blog-select v-model:value="filter.categoryId" v-bind:type="1" size="small"></blog-select>
      <el-input
          placeholder="请输入URL"
          v-model="filter.url"
          size="small"
          clearable/>
      <blog-select v-model:value="filter.enableComment" v-bind:type="3" size="small"></blog-select>
      <blog-select v-model:value="filter.status" v-bind:type="4" size="small"></blog-select>
      <el-button type="primary" plain icon="el-icon-search" @click="onFilterSearch" size="small">搜索</el-button>
      <el-button type="info" plain icon="el-icon-delete" @click="onFilterClear" size="small">清空</el-button>
      <el-button type="primary" @click="onAddArticle" size="small">新增文章</el-button>
    </el-space>

    <el-divider/>

    <el-table
        :data="data.list"
        border
        stripe
        style="width: 100%;"
        @sort-change="onSortChange"
        :height="mainContentHeight">
      <el-table-column prop="id" label="ID" width="60" fixed sortable="custom"></el-table-column>
      <el-table-column prop="title" label="标题" min-width="250" fixed>
        <template #default="scope">
          <el-link :href="'/article/' + scope.row.id" type="primary" target="_blank">{{ scope.row.title }}</el-link>
        </template>
      </el-table-column>
      <el-table-column prop="categoryName" label="分类" width="100"></el-table-column>
      <el-table-column prop="url" label="url" width="100" sortable="custom">
        <template #default="scope">
          {{ scope.row.url ? scope.row.url : '-'}}
        </template>
      </el-table-column>
      <el-table-column prop="weight" label="权重" width="80" sortable="custom"></el-table-column>
      <el-table-column prop="views" label="浏览数" width="90" sortable="custom"></el-table-column>
      <el-table-column :formatter="onFormatDate" prop="createTime" label="创建时间" width="160" sortable="custom"></el-table-column>
      <el-table-column :formatter="onFormatDate" prop="updateTime" label="更新时间" width="160" sortable="custom"></el-table-column>
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
              size="medium">{{ scope.row.status === 1 ? '发布' : '草稿'}}
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
import {queryArticle, deleteArticle} from "@/api"
import BlogSelect from '@/components/BlogSelect.vue'

export default defineComponent({
  name: 'Article',
  components: {
    BlogSelect
  },
  mounted() {
    this.onRefresh()
    window.onresize = () => {
      this.onRefreshMainContentHeight()
    }
    this.onRefreshMainContentHeight()
  },
  unmounted() {
    window.onresize = null;
  },
  data() {
    const mainContentBottom = 250
    return {
      mainContentBottom,
      mainContentHeight: window.innerHeight - mainContentBottom,

      filter: {
        title: null,
        categoryId: null,
        url: null,
        enableComment: null,
        status: null,
      },

      data: {
        pageNum: 1,
        pageSize: 20,
        pages: 0,
        total: 0,
        list: [],
      },

      order: {
        name: null,
        method: null,
      },
    }
  },
  methods: {
    onRefreshMainContentHeight() {
      this.mainContentHeight = window.innerHeight - this.mainContentBottom
    },

    onFilterSearch() {
      this.onRefresh()
    },
    onFilterClear() {
      this.filter.title = null
      this.filter.categoryId = null
      this.filter.url = null
      this.filter.enableComment = null
      this.filter.status = null
      this.onRefresh()
    },
    onSortChange(column: any) {
      this.order.name = column.prop
      this.order.method = column.order
      this.onRefresh()
    },
    onFormatDate(row: any, column: any) {
      const date = row[column.property];
      if (date == null) return "-"
      return dayjs(date).format("YYYY-MM-DD HH:mm:ss");
    },
    onCurrentPageChange(currentPage: number) {
      const changed = this.data.pageNum != currentPage
      this.data.pageNum = currentPage;
      if (changed) {
        this.onRefresh();
      }
    },
    onAddArticle() {
      this.$router.push({
        path: 'editor',
        name: '文章编辑',
      })
    },
    onEditArticle(row: any) {
      this.$router.push({
        path: 'editor',
        name: '文章编辑',
        params: {
          articleId: row.id,
        }
      })
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
          this.filter.title,
          this.filter.categoryId,
          this.filter.url,
          this.filter.enableComment,
          this.filter.status,
          this.data.pageNum,
          this.data.pageSize,
          this.order.name,
          this.order.method,
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