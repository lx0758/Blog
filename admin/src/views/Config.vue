<template>
  <el-container direction="vertical">

    <el-space wrap>
      <el-input
          placeholder="请输入键"
          v-model="filter.key"
          clearable/>
      <el-input
          placeholder="请输入值"
          v-model="filter.value"
          clearable/>
      <el-input
          placeholder="请输入描述"
          v-model="filter.description"
          clearable/>
      <el-button type="primary" plain icon="el-icon-search" @click="onFilterSearch">搜索</el-button>
      <el-button type="info" plain icon="el-icon-delete" @click="onFilterClear">清空</el-button>
      <el-button type="primary" @click="onAddConfig">新增配置</el-button>
    </el-space>

    <el-divider/>

    <el-table
        :data="data.list"
        border
        stripe
        style="width: 100%;"
        @sort-change="onSortChange"
        :height="mainContentHeight">
      <el-table-column prop="key" label="键" width="250" fixed sortable="custom"></el-table-column>
      <el-table-column prop="value" label="值" min-width="300"></el-table-column>
      <el-table-column prop="description" label="描述" min-width="150"></el-table-column>
      <el-table-column :formatter="onFormatDate" prop="createTime" label="创建时间" width ="160" sortable="custom"></el-table-column>
      <el-table-column :formatter="onFormatDate" prop="updateTime" label="更新时间" width="160" sortable="custom"></el-table-column>
      <el-table-column label="操作" width="160">
        <template #default="scope">
          <el-button plain size="mini" @click="onEditConfig(scope.row)">编辑</el-button>
          <el-button plain type="danger" size="mini" @click="onDeleteConfig(scope.row)">删除</el-button>
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

  <el-dialog :title="(!dialogData.id ? '新增' : '编辑') + '配置'" v-model="dialog" :close-on-click-modal="false">
    <el-form ref="dialog" :model="dialogData" :rules="dialogRules" label-width="120px">
      <el-form-item label="键" prop="key">
        <el-input v-model="dialogData.key" placeholder="请输入配置键"></el-input>
      </el-form-item>
      <el-form-item label="值" prop="value">
        <el-input v-model="dialogData.value" placeholder="请输入配置值"></el-input>
      </el-form-item>
      <el-form-item label="描述" prop="description">
        <el-input v-model="dialogData.description" placeholder="请输入描述"></el-input>
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
import {addConfig, deleteConfig, updateConfig, queryConfig} from "@/api";

export default defineComponent({
  name: 'Setting',
  components: {},
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
        value: null,
        description: null,
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
        value: [
          {required: true, message: '值不能为空', trigger: 'blur'},
        ],
        description: [
          {required: true, message: '描述不能为空', trigger: 'blur'},
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
      this.filter.value = null
      this.filter.description = null
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
    onAddConfig() {
      this.dialogData = {
        id: null,
      }
      this.dialog = true
    },
    onEditConfig(row: any) {
      this.dialogData = {
        id: 1,
        key: row.key,
        value: row.value,
        description: row.description,
      };
      this.dialog = true
    },
    onDeleteConfig(row: any) {
      this.$confirm('确认删除?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteConfig(row.id)
            .then(() => {
              this.$message.success("删除成功");
              this.onRefresh()
            })
      });
    },
    onRefresh() {
      queryConfig(
          this.filter.key,
          this.filter.value,
          this.filter.description,
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
          addConfig(
              dialogData.key,
              dialogData.value,
              dialogData.description,
          )
              .then(() => {
                this.$message.success("新增成功");
                this.dialog = false
                this.onRefresh()
              })
        } else {
          updateConfig(
              dialogData.key,
              dialogData.value,
              dialogData.description,
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
