<template>
  <el-container direction="vertical">

    <el-space wrap>
      <el-input
          placeholder="请输入名称"
          v-model="filter.name"
          clearable/>
      <el-button type="primary" plain icon="el-icon-search" @click="onFilterSearch">搜索</el-button>
      <el-button type="info" plain icon="el-icon-delete" @click="onFilterClear">清空</el-button>
      <el-button type="primary" @click="onAddTag">新增标签</el-button>
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
      <el-table-column prop="name" label="名称" min-width="100" fixed sortable="custom"></el-table-column>
      <el-table-column prop="articleCount" label="文章数量" width="110" sortable="custom"></el-table-column>
      <el-table-column :formatter="onFormatDate" prop="createTime" label="创建时间" width="160" sortable="custom"></el-table-column>
      <el-table-column label="操作" width="160">
        <template #default="scope">
          <el-button plain size="mini" @click="onEditTag(scope.row)">编辑</el-button>
          <el-button plain type="danger" size="mini" @click="onDeleteTag(scope.row)">删除</el-button>
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

  <el-dialog :title="(!dialogData.id ? '新增' : '编辑') + '标签'" v-model="dialog" :close-on-click-modal="false">
    <el-form ref="dialog" :model="dialogData" :rules="dialogRules" label-width="120px">
      <el-form-item label="名称" prop="name">
        <el-input v-model="dialogData.name" placeholder="请输入名称"></el-input>
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
import {addTag, deleteTag, updateTag, queryTag} from "@/api";

export default defineComponent({
  name: 'Tag',
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
        name: null,
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
        name: [
          {required: true, message: '名称不能为空', trigger: 'blur'},
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
      this.filter.name = null
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
    onAddTag() {
      this.dialogData = {}
      this.dialog = true
    },
    onEditTag(row: any) {
      this.dialogData = {
        id: row.id,
        name: row.name,
      };
      this.dialog = true
    },
    onDeleteTag(row: any) {
      this.$confirm('确认删除?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteTag(row.id)
            .then(() => {
              this.$message.success("删除成功");
              this.onRefresh()
            })
      });
    },
    onRefresh() {
      queryTag(
          this.filter.name,
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
          addTag(
              dialogData.name,
          )
              .then(() => {
                this.$message.success("新增成功");
                this.dialog = false
                this.onRefresh()
              })
        } else {
          updateTag(
              dialogData.id,
              dialogData.name,
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
