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
          <el-link :href="scope.row.url" type="primary" target="_blank">{{ scope.row.displayName }}</el-link>
        </template>
      </el-table-column>
      <el-table-column :formatter="onFormatFileSize" prop="length" label="大小" width="100"></el-table-column>
      <el-table-column prop="type" label="类型" width="120"></el-table-column>
      <el-table-column prop="path" label="存储路径" min-width="250" :show-overflow-tooltip="true">
        <template  #default="scope">
          <el-link :href="scope.row.url" type="info" target="_blank">{{ scope.row.path }}</el-link>
        </template>
      </el-table-column>
      <el-table-column prop="authorName" label="作者" width="80"></el-table-column>
      <el-table-column :formatter="onFormatDate" prop="createTime" label="上传时间" width ="160"></el-table-column>
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

  <el-dialog title="上传文件" v-model="dialog" width="500px" :close-on-click-modal="false">
    <el-form ref="verify" :model="dialogData" label-width="80px">
      <el-form-item label="选择文件">
        <el-upload
            drag
            action=""
            :on-change="onUploadChange"
            :multiple="false"
            :auto-upload="false"
            :show-file-list="false">
          <i class="el-icon-upload"></i>
          <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
        </el-upload>
      </el-form-item>
      <el-form-item label="文件名称">
        <span>{{onFormatFileName(dialogData.name, '请先选择文件')}}</span>
      </el-form-item>
      <el-form-item label="文件长度">
        <span>{{onFormatByteSize(dialogData.size, '请先选择文件')}}</span>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onDialogSubmit">上传</el-button>
        <el-button @click="dialog = false">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>

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
      },

      data: {
        pageNum: 1,
        pageSize: 20,
        pages: 0,
        total: 0,
        list: [],
      },

      dialog: false,
      dialogData: {},
    }
  },
  methods: {
    onFilterSearch() {
      this.onRefresh()
    },
    onFilterClear() {
      this.filter.displayName = null
      this.filter.type = null
      this.onRefresh()
    },
    onFormatFileSize(row: any, column: any) {
      let temp;
      const fileSize = row[column.property];
      return this.onFormatByteSize(fileSize)
    },
    onFormatFileName(name: string, defaultText: string|null = null) {
      if (!name) {
        if (!defaultText) return '-'
        return defaultText
      }
      return name
    },
    onFormatByteSize(size: number, defaultText: string|null = null): string {
      if (!size) {
        if (!defaultText) return '-'
        return defaultText
      }
      let temp;
      if (size < 1024) {
        return size + ' B';
      } else if (size < (1024 * 1024)) {
        temp = size / 1024;
        temp = temp.toFixed(2);
        return temp + ' KB';
      } else if (size < (1024 * 1024 * 1024)) {
        temp = size / (1024 * 1024);
        temp = temp.toFixed(2);
        return temp + ' MB';
      } else {
        temp = size / (1024 * 1024 * 1024);
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
    onAddUpload() {
      this.dialogData = {}
      this.dialog = true
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
          this.data.pageNum,
          this.data.pageSize,
      )
          .then(data => {
            this.data = data.data;
          })
    },

    onUploadChange(file: any, fileList: any) {
      if (fileList.length == 0) return
      this.dialogData = fileList[fileList.length - 1]
    },
    onDialogSubmit() {
      const dialogData = this.dialogData as any
      if (!dialogData || !dialogData.raw) {
        this.$message('文件不能为空');
        return;
      }
      addUpload(dialogData.raw)
          .then(() => {
            this.$message.success("上传成功");
            this.dialog = false
            this.onRefresh()
          })
    },
  }
});
</script>

<style lang="less">
</style>
