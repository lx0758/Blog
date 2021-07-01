<template>
  <el-container direction="vertical">

    <el-space wrap>
      <el-input
          placeholder="请输入标题"
          v-model="filter.title"
          clearable/>
      <el-input
          placeholder="请输入链接"
          v-model="filter.link"
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
      <el-button type="primary" @click="onEditLink">新增短链</el-button>
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
      <el-table-column prop="url" label="链接地址" min-width="250" :show-overflow-tooltip="true">
        <template  #default="scope">
          <el-link :href="scope.row.url" type="info" target="_blank">{{ scope.row.url }}</el-link>
        </template>
      </el-table-column>
      <el-table-column :formatter="onFormatDate" prop="createTime" label="创建时间" width ="160"></el-table-column>
      <el-table-column :formatter="onFormatDate" prop="updateTime" label="更新时间" width="160"></el-table-column>
      <el-table-column prop="status" label="状态" width="70">
        <template #default="scope">
          <el-tag
              :type="scope.row.status === 1 ? 'success' : scope.row.status === -1 ? 'info' : 'danger'"
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

  <el-dialog :title="(!dialogData.id ? '新增' : '编辑') + '友链'" v-model="dialog">
    <el-form ref="dialog" :model="dialogData" :rules="dialogRules" label-width="120px">
      <el-form-item label="标题" prop="title">
        <el-input v-model="dialogData.title" placeholder="请输入标题"></el-input>
      </el-form-item>
      <el-form-item label="链接" prop="url">
        <el-input v-model="dialogData.url" placeholder="请输入链接地址"></el-input>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select v-model="dialogData.status" placeholder="状态">
          <el-option label="启用" :value="1"></el-option>
          <el-option label="禁用" :value="0"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onDialogSubmit">确定</el-button>
        <el-button @click="dialog = false">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>

</template>

<script lang="ts">
import { defineComponent } from 'vue';
import dayjs from "dayjs";
import {addLink, deleteLink, updateLink, queryLink} from "@/api";

export default defineComponent({
  name: 'Link',
  components: {},
  mounted() {
    this.onRefresh()
  },
  data() {
    return {
      filter: {
        title: null,
        link: null,
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

      dialog: false,
      dialogData: {},
      dialogRules: {
        title: [
          {required: true, message: '标题不能为空', trigger: 'blur'},
        ],
        url: [
          {required: true, message: '链接不能为空', trigger: 'blur'},
        ],
        status: [
          {required: true, message: '请选择状态', trigger: 'blur'},
        ],
      },
    }
  },
  methods: {
    onFilterSearch() {
      this.onRefresh()
    },
    onFilterClear() {
      this.filter.title = null
      this.filter.link = null
      this.filter.status = null
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
      if (row == null) {
        this.dialogData = {}
      } else {
        this.dialogData = {
          id: row.id,
          title: row.title,
          url: row.url,
          status: row.status,
        };
      }
      this.dialog = true
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
          this.filter.title,
          this.filter.link,
          this.filter.status,
          this.data.pageNum,
          this.data.pageSize,
      )
          .then(data => {
            this.data = data.data;
          })
    },

    onDialogSubmit() {
      let from: any = this.$refs['dialog'];
      from.validate((valid: boolean) => {
        if (!valid) return
        let dialogData = this.dialogData as any
        if (!dialogData.id) {
          addLink(
              dialogData.title,
              dialogData.url,
              dialogData.status,
          )
              .then(() => {
                this.$message.success("新增成功");
                this.onRefresh()
              })
        } else {
          updateLink(
              dialogData.id,
              dialogData.title,
              dialogData.url,
              dialogData.status,
          )
              .then(() => {
                this.$message.success("更新成功");
                this.onRefresh()
              })
        }
      })
    },
  }
});
</script>

<style lang="less">
</style>
