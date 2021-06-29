<template>
  <el-container direction="vertical">

    <el-space wrap>
      <el-input
          placeholder="请输入标题"
          v-model="filterTitle"
          clearable/>
      <el-input
          placeholder="请输入链接"
          v-model="filterLink"
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
      <el-table-column prop="title" label="标题" width="240" fixed>
        <template #default="scope">
          <el-link :href="scope.row.url" type="primary" target="_blank">{{ scope.row.title }}</el-link>
        </template>
      </el-table-column>
      <el-table-column prop="url" label="Link" min-width="250" :show-overflow-tooltip="true">
        <template  #default="scope">
          <el-link :href="scope.row.url" type="info" target="_blank">{{ scope.row.url }}</el-link>
        </template>
      </el-table-column>
      <el-table-column :formatter="onFormatDate" prop="createTime" label="创建时间" width ="160"></el-table-column>
      <el-table-column :formatter="onFormatDate" prop="updateTime" label="更新时间" width="160"></el-table-column>
      <el-table-column prop="status" label="状态" width="70">
        <template #default="scope">
          <el-tag
              :type="scope.row.status === 1 ? 'success' : 'danger'"
              disable-transitions
              size="medium">{{ scope.row.status === 1 ? '启用' : scope.row.status === -1 ? '删除' : '禁用'}}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="160">
        <template #default="scope">
          <el-button plain size="mini" @click="onEditLink(scope.row)">编辑</el-button>
          <el-button plain type="danger" size="mini" @click="onDeleteLink(scope.row)">删除</el-button>
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
import {deleteLink, queryLink} from "@/api";

export default defineComponent({
  name: 'Link',
  components: {},
  mounted() {
    this.onRefresh()
  },
  data() {
    return {
      filterTitle: null,
      filterLink: null,
      filterStatus: null,
      filterStatusOptions: [
        {
          value: -1,
          label: '已删除',
        },
        {
          value: 0,
          label: '禁用',
        },
        {
          value: 1,
          label: '启用',
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
      this.filterTitle = null
      this.filterLink = null
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
    onEditLink(row: any) {
      // TODO: 2021-6-28
      console.log("onEditLink:" + row.key)
    },
    onDeleteLink(row: any) {
      this.$confirm('确认删除?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteLink(row.id)
            .then(() => {
              this.$message.success("删除成功");
              this.onRefresh()
            })
      });
    },
    onRefresh() {
      queryLink(
          this.filterTitle,
          this.filterLink,
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
