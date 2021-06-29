<template>
  <el-container direction="vertical">

    <el-space wrap>
      <el-input
          placeholder="请输入名称"
          v-model="filterName"
          clearable/>
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
      <el-table-column prop="name" label="名称" min-width="100" fixed></el-table-column>
      <el-table-column :formatter="onFormatDate" prop="createTime" label="创建时间" width="160"></el-table-column>
      <el-table-column label="操作" width="160">
        <template #default="scope">
          <el-button plain size="mini" @click="onEditCategory(scope.row)">编辑</el-button>
          <el-button plain type="danger" size="mini" @click="onDeleteCategory(scope.row)">删除</el-button>
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
import dayjs from "dayjs";
import {deleteCategory, queryCategory} from "@/api";

export default defineComponent({
  name: 'Category',
  components: {},
  mounted() {
    this.onRefresh()
  },
  data() {
    return {
      filterName: null,

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
      this.filterName = null
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
    onEditCategory(row: any) {
      // TODO: 2021-6-28
      console.log("onEditCategory:" + row.key)
    },
    onDeleteCategory(row: any) {
      this.$confirm('确认删除?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteCategory(row.id)
            .then(() => {
              this.$message.success("删除成功");
              this.onRefresh()
            })
      });
    },
    onRefresh() {
      queryCategory(
          this.filterName,
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
