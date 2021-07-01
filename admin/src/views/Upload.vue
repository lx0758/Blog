<template>
  <el-container direction="vertical">

    <el-space wrap>
      <el-input
          placeholder="请输入文件名"
          v-model="filter.displayName"
          clearable/>
      <el-input
          placeholder="请输入类型"
          v-model="filter.type"
          clearable/>
      <el-select v-model="filter.status" placeholder="状态">
        <el-option
            v-for="item in filter.statusOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value">
        </el-option>
      </el-select>
      <el-button type="primary" plain icon="el-icon-search" @click="onFilterSearch">搜索</el-button>
      <el-button type="info" plain icon="el-icon-delete" @click="onFilterClear">清空</el-button>
      <el-button type="primary" @click="onAddUpload">上传文件</el-button>
    </el-space>

    <el-divider/>

    <el-table
        :data="data.list"
        border
        stripe
        style="width: 100%; height: auto">
      <el-table-column prop="id" label="ID" width="60" fixed></el-table-column>
      <el-table-column prop="displayName" label="文件名" width="300" fixed>
        <template #default="scope">
          <el-link :href="scope.row.path" type="primary" target="_blank">{{ scope.row.displayName }}</el-link>
        </template>
      </el-table-column>
      <el-table-column :formatter="onFormatSize" prop="length" label="大小" width="120"></el-table-column>
      <el-table-column prop="type" label="类型" width="100"></el-table-column>
      <el-table-column prop="path" label="路径" min-width="250" :show-overflow-tooltip="true">
        <template  #default="scope">
          <el-link :href="scope.row.path" type="info" target="_blank">{{ scope.row.path }}</el-link>
        </template>
      </el-table-column>
      <el-table-column prop="authorName" label="作者" width="80"></el-table-column>
      <el-table-column :formatter="onFormatDate" prop="createTime" label="上传时间" width ="160"></el-table-column>
      <el-table-column prop="status" label="状态" width="70">
        <template #default="scope">
          <el-tag
              :type="scope.row.status === 1 ? 'success' : scope.row.status === -1 ? 'info' : 'danger'"
              disable-transitions
              size="medium">{{ scope.row.status === 1 ? '启用' : scope.row.status === -1 ? '删除' : '禁用'}}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="80">
        <template #default="scope">
          <el-button plain type="danger" size="mini" @click="onDeleteUpload(scope.row)">删除</el-button>
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
import {addUpload, deleteUpload, queryUpload,} from "@/api";

export default defineComponent({
  name: 'Upload',
  components: {},
  mounted() {
    this.onRefresh()
  },
  data() {
    return {
      filter: {
        displayName: null,
        type: null,
        status: null,
        statusOptions: [
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
      },

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
      this.filter.displayName = null
      this.filter.type = null
      this.filter.status = null
      this.onRefresh()
    },
    onFormatSize(row: any, column: any) {
      let temp;
      const fileSize = row[column.property];
      if (fileSize < 1024) {
        return fileSize + ' B';
      } else if (fileSize < (1024 * 1024)) {
        temp = fileSize / 1024;
        temp = temp.toFixed(2);
        return temp + ' KB';
      } else if (fileSize < (1024 * 1024 * 1024)) {
        temp = fileSize / (1024 * 1024);
        temp = temp.toFixed(2);
        return temp + ' MB';
      } else {
        temp = fileSize / (1024 * 1024 * 1024);
        temp = temp.toFixed(2);
        return temp + ' GB';
      }
    },
    onFormatDate(row: any, column: any) {
      const date = row[column.property];
      return dayjs(date).format("YYYY-MM-DD HH:mm:ss");
    },
    onCurrentPageChange(currentPage: number) {
      this.data.pageNum = currentPage;
      this.onRefresh();
    },
    onEditUpload(row: any) {
      // TODO: 2021-6-28
      console.log("onEditUpload:" + row.key)
    },
    onDeleteUpload(row: any) {
      this.$confirm('确认删除?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteUpload(row.id)
            .then(() => {
              this.$message.success("删除成功");
              this.onRefresh()
            })
      });
    },
    onRefresh() {
      queryUpload(
          this.filter.displayName,
          this.filter.type,
          this.filter.status,
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
