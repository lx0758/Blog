<template>
  <el-container direction="vertical">

    <el-space wrap>
      <el-input
          placeholder="请输入键"
          v-model="filter.key"
          clearable/>
      <el-input
          placeholder="请输入链接"
          v-model="filter.url"
          clearable/>
      <el-input
          placeholder="请输入描述"
          v-model="filter.description"
          clearable/>
      <blog-select v-model:value="filter.status" v-bind:type="7"></blog-select>
      <el-button type="primary" plain icon="el-icon-search" @click="onFilterSearch">搜索</el-button>
      <el-button type="info" plain icon="el-icon-delete" @click="onFilterClear">清空</el-button>
      <el-button type="primary" @click="onAddUrl">新增短链</el-button>
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
      <el-table-column prop="key" label="主键" width="120" fixed sortable="custom">
        <template #default="scope">
          <el-link :href="'/u/' + scope.row.key" type="primary" target="_blank">{{ scope.row.key }}</el-link>
        </template>
      </el-table-column>
      <el-table-column prop="url" label="链接" min-width="250" :show-overflow-tooltip="true">
        <template  #default="scope">
          <el-link :href="scope.row.url" type="info" target="_blank">{{ scope.row.url }}</el-link>
        </template>
      </el-table-column>
      <el-table-column prop="description" label="描述" min-width="150"></el-table-column>
      <el-table-column prop="views" label="访问量" width="90" sortable="custom"></el-table-column>
      <el-table-column prop="totalViews" label="总访问量" width="110" sortable="custom"></el-table-column>
      <el-table-column :formatter="onFormatDate" prop="createTime" label="创建时间" width ="160" sortable="custom"></el-table-column>
      <el-table-column :formatter="onFormatDate" prop="updateTime" label="更新时间" width="160" sortable="custom"></el-table-column>
      <el-table-column prop="status" label="状态" width="70">
        <template #default="scope">
          <el-tag
              :type="scope.row.status === 1 ? 'success' : 'danger'"
              disable-transitions
              size="medium">{{ scope.row.status === 1 ? '启用' : '禁用'}}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="160">
        <template #default="scope">
          <el-button plain size="mini" @click="onEditUrl(scope.row)">编辑</el-button>
          <el-button plain type="danger" size="mini" @click="onDeleteUrl(scope.row)">删除</el-button>
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

  <el-dialog :title="(!dialogData.id ? '新增' : '编辑') + '短链'" v-model="dialog" :close-on-click-modal="false">
    <el-form ref="dialog" :model="dialogData" :rules="dialogRules" label-width="120px">
      <el-form-item label="主键" prop="key">
        <el-input v-model="dialogData.key" placeholder="请输入键"></el-input>
      </el-form-item>
      <el-form-item label="链接" prop="url">
        <el-input v-model="dialogData.url" placeholder="请输入链接"></el-input>
      </el-form-item>
      <el-form-item label="描述" prop="description">
        <el-input v-model="dialogData.description" placeholder="请输入描述"></el-input>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <blog-select v-model:value="dialogData.status" v-bind:type="7"></blog-select>
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
import {addUrl, deleteUrl, updateUrl, queryUrl} from "@/api";
import BlogSelect from "@/components/BlogSelect.vue";

export default defineComponent({
  name: 'Url',
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
        key: null,
        url: null,
        description: null,
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

      dialog: false,
      dialogData: {},
      dialogRules: {
        key: [
          {required: true, message: '键不能为空', trigger: 'blur'},
        ],
        url: [
          {required: true, message: '链接不能为空', trigger: 'blur'},
        ],
        description: [
          {required: true, message: '描述不能为空', trigger: 'blur'},
        ],
        status: [
          {required: true, message: '请选择状态', trigger: 'blur'},
        ],
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
      this.filter.key = null
      this.filter.url = null
      this.filter.description = null
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
      this.data.pageNum = currentPage;
      this.onRefresh();
    },
    onAddUrl() {
      this.dialogData = {}
      this.dialog = true
    },
    onEditUrl(row: any) {
      this.dialogData = {
        id: row.id,
        key: row.key,
        url: row.url,
        description: row.description,
        status: row.status,
      };
      this.dialog = true
    },
    onDeleteUrl(row: any) {
      this.$confirm('确认删除?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteUrl(row.id)
            .then(() => {
              this.$message.success("删除成功");
              this.onRefresh()
            })
      });
    },
    onRefresh() {
      queryUrl(
          this.filter.key,
          this.filter.url,
          this.filter.description,
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

    onDialogSubmit() {
      let from: any = this.$refs['dialog'];
      from.validate((valid: boolean) => {
        if (!valid) return
        let dialogData = this.dialogData as any
        if (!dialogData.id) {
          addUrl(
              dialogData.key,
              dialogData.url,
              dialogData.description,
              dialogData.status,
          )
              .then(() => {
                this.$message.success("新增成功");
                this.dialog = false
                this.onRefresh()
              })
        } else {
          updateUrl(
              dialogData.id,
              dialogData.key,
              dialogData.url,
              dialogData.description,
              dialogData.status,
          )
              .then(() => {
                this.$message.success("更新成功");
                this.dialog = false
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
